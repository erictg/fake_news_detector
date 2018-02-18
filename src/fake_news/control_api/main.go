package main

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"fake_news/control_api/processing"
)

func main(){
	router := mux.NewRouter()

	router.HandleFunc("/rest/analysis", processing.OverallAnalyze).Methods("POST")

	c := cors.AllowAll()
	log.Fatal(http.ListenAndServe(":8004", c.Handler(router)))
}
