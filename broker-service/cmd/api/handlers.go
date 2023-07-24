package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)


type RequestPayload struct{
	Action string `json:"action"`
	Auth AuthPayload `json:auth,ommitempty`
	Log LogPayload `json:log,ommitempty`
}

type AuthPayload struct{
	Email string `json:email`
	Password string `json:password`
}

type LogPayload struct{
	Name string `json:"name"`
	Data string `json:"data"`
}


func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Message: "Hit the broker",
		Error: false,
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}
func (app *Config) HandleSubmission(w http.ResponseWriter, r *http.Request) {
	var requestPayload RequestPayload

	err :=app.readJSON(w, r, &requestPayload)
	if err !=nil{
		app.errorJSON(w, err)
		return
	}

	switch requestPayload.Action{
	case "auth":
		app.authenticate(w, requestPayload.Auth)
	case "log":
		app.log(w, requestPayload.Log)
	default:
		app.errorJSON(w, errors.New("Unknown action"))
	}

}

func (app *Config) log(w http.ResponseWriter, l LogPayload){
	jsonData, _ := json.MarshalIndent(l, "", "\t")

	request, err := http.NewRequest("POST", "http://logger-service/log", bytes.NewBuffer(jsonData))
	
	if err !=nil{
		app.errorJSON(w, err)
		return
	}

	client := &http.Client{};
	response, err := client.Do(request)

	if err !=nil{
		app.errorJSON(w, err)
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized{
		app.errorJSON(w, errors.New("Something went wrong "))
		return
	}else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("Error calling auth service"))
		return
	}

	var jsonFromService jsonResponse

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err !=nil{
		app.errorJSON(w, err)
		return
	}
	
	if jsonFromService.Error {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "Logged"
	payload.Data = jsonFromService.Data

	app.writeJSON(w, http.StatusAccepted, payload)
}  

func (app *Config) authenticate(w http.ResponseWriter, a AuthPayload){
	jsonData, _ := json.MarshalIndent(a, "", "\t")

	request, err := http.NewRequest("POST", "http://authentication-service/authenticate", bytes.NewBuffer(jsonData))
	
	if err !=nil{
		app.errorJSON(w, err)
		return
	}

	client := &http.Client{};
	response, err := client.Do(request)

	if err !=nil{
		app.errorJSON(w, err)
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized{
		app.errorJSON(w, errors.New("aa"))
		return
	}else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("Error calling auth service"))
		return
	}

	var jsonFromService jsonResponse

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err !=nil{
		app.errorJSON(w, err)
		return
	}
	
	if jsonFromService.Error {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "Authenticated"
	payload.Data = jsonFromService.Data

	app.writeJSON(w, http.StatusAccepted, payload)
}  

