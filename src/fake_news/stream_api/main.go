package main

import (
	"github.com/gorilla/mux"
	"fake_news/stream_api/service"
	"log"
	"net/http"
	"github.com/rs/cors"
)

func main(){
	router := mux.NewRouter()

	router.HandleFunc("/rest/stream", service.Post).Methods("POST")
	c := cors.AllowAll()
	log.Fatal(http.ListenAndServe(":8001", c.Handler(router)))

}
