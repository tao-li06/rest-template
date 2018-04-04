package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	Config.load()

	controller := Controller{}
	controller.Init()

	router := mux.NewRouter()
	router.HandleFunc("/api/model/{id}", controller.GetModel).Methods("GET")

	http.ListenAndServe("127.0.0.1:8888", router)
}
