package processing

import (
	"net/http"
	"encoding/json"
	"log"
)

type RestDTO struct{
	Content 	string	`json:"content"`
}

type RestResponse struct{
	Result	bool	`json:"result"`
}
func OverallAnalyze(w http.ResponseWriter, r *http.Request){
	var dto RestDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil{
		log.Println(err)
		http.Error(w, err.Error(), 400)
		return
	}

	result, err := Analysis(dto.Content)
	if err != nil{
		log.Println(err)
		http.Error(w, err.Error(), 400)
		return
	}
	toReturn := new(RestResponse)
	toReturn.Result = result

	json, _ := json.Marshal(toReturn)
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
