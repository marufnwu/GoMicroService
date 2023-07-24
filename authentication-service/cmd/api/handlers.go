package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (app *Config) Authenticate(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct{
		Email string `json"email"`
		Password string `json"Password"`
	}

	err := app.readJSON(w, r, &requestPayload)

	if err!=nil{
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	user, err := app.Models.User.GetByEmail(requestPayload.Email);
	if err!=nil{
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	valid, err := user.PasswordMatches(requestPayload.Password)

	if err!=nil || !valid{
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	//write loging log

	app.logRequest("Authentication", fmt.Sprintf("%s Logged in", user.Email))



	payload := jsonResponse{
		Error: false,
		Message: fmt.Sprintf("Login in user %s", user.Email),
		Data: user,
	}

	app.writeJSON(w, http.StatusAccepted, payload)


}


func (app Config) logRequest(name, data string) error{
	var logEntry struct{
		Name string `json:"name"`
		Data string `json:data`
	}

	logEntry.Data = data
	logEntry.Name = name


	jsonData, _ := json.MarshalIndent(logEntry, "", "\t")

	request, err := http.NewRequest("POST", "http://logger-service/log", bytes.NewBuffer(jsonData))
	
	if err !=nil{
		return err
	}

	client := &http.Client{};
	response, err := client.Do(request)

	if err !=nil{
		return err
	}

	defer response.Body.Close()

	return nil
}