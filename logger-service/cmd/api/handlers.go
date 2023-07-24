package main

import (
	"logger-service/data"
	"net/http"
)


type JSONPaylaod struct{
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *Config) WriteLogs(w http.ResponseWriter, r *http.Request) {
	var requestPaylaod JSONPaylaod

	_ = app.readJSON(w, r, &requestPaylaod)

	//insert

	event := data.LogEntry{
		Name: requestPaylaod.Name,
		Data: requestPaylaod.Data,
	}

	err :=app.Models.LogEntry.Insert(event)

	if err!=nil{
		app.errorJSON(w, err)
	}

	resp := jsonResponse{
		Error: false,
		Message: "logged",
	}

	app.writeJSON(w, http.StatusAccepted, resp)

}
