package processing

import (
	"regexp"
	"log"
)

func Analysis(content string) (bool, error){

	sent, err := seniment(regSplit(content, "!?."))
	if err != nil{
		log.Println(err)
		return false, err
	}

	verification, err := verification(content)
	if err != nil{
		log.Println(err)
		return false, err
	}

	log.Println("analysis")

	//some math
	avgSent := averageSentiment(sent)
	if verification.Total == 0{
		avgVer := float32(0)

		return overall(avgSent, avgVer)
	}else{
		avgVer := verification.Score / verification.Total

		return overall(avgSent, avgVer)
	}

}


func averageSentiment(response []int) float32{
	scoreTotal := 0

	for v := range response{
		scoreTotal += v
	}

	return float32(float32(scoreTotal)/float32(len(response)))
}

func regSplit(text string, delimeter string) []string {
	reg := regexp.MustCompile(delimeter)
	indexes := reg.FindAllStringIndex(text, -1)
	laststart := 0
	result := make([]string, len(indexes) + 1)
	for i, element := range indexes {
		result[i] = text[laststart:element[0]]
		laststart = element[1]
	}
	result[len(indexes)] = text[laststart:len(text)]
	return result
}
