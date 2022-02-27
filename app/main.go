package main

import (
	"encoding/json"
	"net/http"
	"os"
	"fmt"
)

type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func home(w http.ResponseWriter, req *http.Request) {
	status := Status{Code: 200, Message: "Home"}
	response, err := json.Marshal(status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	}
}

func status(w http.ResponseWriter, req *http.Request) {
	status_from_env := os.Getenv("APP_STATUS")
	status := Status{Code: 500, Message: status_from_env}
	if status_from_env == "OK" {
		status.Code = 200
	}
	response, err := json.Marshal(status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
		fmt.Println(string(response))
	}
}

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/status", status)
	http.ListenAndServe(":8080", nil)
}
