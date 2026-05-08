package jsonapi

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"log"
	"mailinglist/mdb"
	"net/http"
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

func CreateEmail(db *sql.DB) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, req *http.Request){
		if req.Method != "POST" {
			return
		} 

		entry := mdb.EmailEntry{}
		fromJson(req.Body, &entry)

		if err := mdb.CreateEmail(db, entry.Email) ; err != nil {
			returnErr(w, err, 400)
			return
		}

		returnJson(w, func() (interface{}, error){
			log.Printf("json createEmail: %v\n", entry.Email)
			return mdb.GetEmail(db, entry.Email)
		})

	})
}

func GetEmail(db *sql.DB) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, req *http.Request){
		if req.Method != "GET" {
			return
		} 

		entry := mdb.EmailEntry{}
		fromJson(req.Body, &entry)

		returnJson(w, func() (interface{}, error){
			log.Printf("json GetEmail: %v\n", entry.Email)
			return mdb.GetEmail(db, entry.Email)
		})

	})
}

func GetEmailBatch(db *sql.DB) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, req *http.Request){
		if req.Method != "GET" {
			return
		} 

		queryOptions := mdb.GetEmailBatchQueryParams{}
		fromJson(req.Body, &queryOptions)

		if queryOptions.Count <= 0 || queryOptions.Page <= 0 {
			returnErr(w, errors.New("PAge and Coutn fields are required and must be > 0"), 400)
			return
		}

		returnJson(w, func() (interface{}, error){
			log.Printf("JSON GetEMailBatch: %v\n", queryOptions)
			return mdb.GetEmailBatch(db, queryOptions)
		})
	})
}

func UpdatEmail(db *sql.DB) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, req *http.Request){
		if req.Method != "PUT" {
			return
		} 

		entry := mdb.EmailEntry{}
		fromJson(req.Body, &entry)

		if err := mdb.UpdateEmail(db, entry) ; err != nil {
			returnErr(w, err, 400)
			return
		}

		returnJson(w, func() (interface{}, error){
			log.Printf("json UpdateEmail: %v\n", entry.Email)
			return mdb.GetEmail(db, entry.Email)
		})

	})
}

func DeleteEmail(db *sql.DB) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, req *http.Request){
		if req.Method != "POST" {
			return
		} 

		entry := mdb.EmailEntry{}
		fromJson(req.Body, &entry)

		if err := mdb.DeleteEmail(db, entry.Email) ; err != nil {
			returnErr(w, err, 400)
			return
		}

		returnJson(w, func() (interface{}, error){
			log.Printf("json DeleteEmail: %v\n", entry.Email)
			return mdb.GetEmail(db, entry.Email)
		})

	})
}

func Serve(db *sql.DB, bind string) {
	http.Handle("/email/create", CreateEmail(db))
	http.Handle("/email/get", GetEmail(db))
	http.Handle("/email/get_batch", GetEmailBatch(db))
	http.Handle("/email/update", UpdatEmail(db))
	http.Handle("/email/delete", DeleteEmail(db))

	err := http.ListenAndServe(bind, nil)
	if err != nil {
		log.Fatalf("JSON server error: %v", err)
	}
}