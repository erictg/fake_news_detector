package main

import (
	"log"
	"github.com/gorilla/mux"
	"net/http"
	"fake_news/sentiment/analysis"
)

func main(){

	router := mux.NewRouter()

	analysis.Init()

	router.HandleFunc("/rest/sentiment", analysis.RestAnalyze).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))

}


