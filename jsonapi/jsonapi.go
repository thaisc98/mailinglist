package jsonapi

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"structs"
)


func setJsonHeader(w http.ResponseWriter){
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
}

func fromJson[T any](body io.Reader, target T){
	buf := new(bytes.Buffer)
	buf.ReadFrom(body)
	json.Unmarshal(buf.Bytes(), &target)
}

func returnJson[T any](w http.ResponseWriter, withData func() (T, error)){
	setJsonHeader(w)

	data, serverError := withData()
	if serverError != nil {
		w.WriteHeader(500)
		serverErrorJson, err := json.Marshal(&serverError)
		if err != nil {
			log.Print(err)
			return
		}
		w.Write(serverErrorJson)
		return
	}

	dataJson, err := json.Marshal(&data)
	if err != nil {
		log.Print(err)
		w.WriteHeader(500)
		return
	}

	w.Write(dataJson)
}

func returnErr(w http.ResponseWriter, err error, code int) {
	returnJson(w, func() (interface{}, error) {
		errorMessage := struct {
			Err string
		}{
			Err: err.Error(),
		}
		w.WriteHeader(code)
		return errorMessage, nil
	})
}