package analysis

import (
	"github.com/cdipaolo/sentiment"
	"fmt"
)

var model sentiment.Models

func Init(){
	m, err := sentiment.Restore()
	if err != nil{
		panic(fmt.Sprintf("Could not restore model!\n\t%v\n", err))
	}
	model = m
}

func analyze(sent string) (int, error){
	result := model.SentimentAnalysis(sent, sentiment.English)
	return int(result.Score), nil
}

func AnalyzeMany(sentences []string) ([]int){
	toReturn := make([]int, len(sentences))
	for r, sent := range sentences{
		toReturn[r], _ = analyze(sent)
	}

	return toReturn
}