package service

import (
	"net/http"
	"fake_news/stream_api/model"
	"encoding/json"
	"log"
)

func Post(w http.ResponseWriter, r *http.Request){
	var dto model.TrainingDTO

	_ = json.NewDecoder(r.Body).Decode(&dto)

	err := PostToDB(dto)

	if err != nil{
		log.Println(err)
		http.Error(w, err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusOK)
}