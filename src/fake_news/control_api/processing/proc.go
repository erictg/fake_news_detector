package processing

import (
	"encoding/json"
	"net/http"
	"bytes"
	"log"
)

type SentimentDTO struct{
	Sentences 	[]string	`json:"sentences"`
}

type SentimentResponse struct{
	Responses	[]int	`json:"responses"`
}

func seniment(val []string) ([]int, error){
	baseUrl := "http://sentiment:8000/rest/sentiment"

	sentiment := SentimentDTO{Sentences:val}

	jsonVal, _ := json.Marshal(sentiment)

	resp, err := http.Post(baseUrl, "application/json", bytes.NewBuffer(jsonVal))
	if err != nil{
		log.Println(err)
		return nil, err
	}

	var response SentimentResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil{
		log.Println(err)
		return nil, err
	}


	return response.Responses, nil
}

type VerificationDTO struct{
	Content 	string	`json:"content"`
}

type VerificationResponse struct{
	Total 	float32 	`json:"total"`
	Score	float32		`json:"score"`
}

func verification(val string) (*VerificationResponse, error){
	baseUrl := "http://verification:8002/rest/matching"

	dto := VerificationDTO{Content:val}

	jsonVal, _ := json.Marshal(dto)

	resp, err := http.Post(baseUrl, "application/json", bytes.NewBuffer(jsonVal))
	if err != nil{
		log.Println(err)
		return nil, err
	}

	var response VerificationResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil{
		log.Println(err)
		return nil, err
	}


	return &response, nil


}

type OverallDTO struct{
	Ratio 		float32 	`json:"ratio"`
	Sentiment	float32		`json:"sentiment"`
}

type OverallResponse struct{
	Truthiness 	bool	`json:"truthiness"`
}


func overall(ratio, verification float32) (bool, error){
	baseUrl := "http://main_learn:8003/rest/overall"

	dto := OverallDTO{Ratio:verification, Sentiment:ratio}

	jsonVal, _ := json.Marshal(dto)

	resp, err := http.Post(baseUrl, "application/json", bytes.NewBuffer(jsonVal))
	if err != nil{
		log.Println(err)
		return false, err
	}

	var response OverallResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil{
		log.Println(err)
		return false, err
	}

	return response.Truthiness, nil
}
