package main

import (
	"github.com/gorilla/mux"
	"fake_news/stream_api/service"
	"log"
	"net/http"
)

func main(){
	router := mux.NewRouter()

	router.HandleFunc("/rest/stream", service.Post).Methods("POST")

	log.Fatal(http.ListenAndServe(":8001", router))

}
