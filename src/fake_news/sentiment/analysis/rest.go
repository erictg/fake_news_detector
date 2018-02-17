package analysis

import (
	"net/http"
	"encoding/json"
)

type SentDTO struct{
	Sentences []string `json:"sentences"`
}

type Response struct{
	Responses []int `json:"responses"`
}

func RestAnalyze(w http.ResponseWriter, r *http.Request){
	var dto SentDTO

	_ = json.NewDecoder(r.Body).Decode(&dto)

	result := AnalyzeMany(dto.Sentences)

	toReturn := new(Response)
	toReturn.Responses = result

	json, _ := json.Marshal(toReturn)
	w.WriteHeader(http.StatusOK)
	w.Write(json)

}