package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// writeJSON writes json formatted data to our server
func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, wrap string) error {
	wrapper := make(map[string]interface{})

	wrapper[wrap] = data

	js, err := json.Marshal(wrapper)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(js)
	if err != nil {
		return err
	}

	return nil
}

// errorJSON handles how errors are displayed in json format
func (app *application) errorJSON(w http.ResponseWriter, err error) {
	type jsonError struct {
		Message string `json:"message"`
	}

	theError := jsonError{
		Message: err.Error(),
	}

	err = app.writeJSON(w, http.StatusBadRequest, theError, "error")
	if err != nil {
		log.Println(err)
	}

}
