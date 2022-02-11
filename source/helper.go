package source

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"model"
)

func EncodeJSONBody(resp http.ResponseWriter, statusCode int, data interface{}) {
	resp.WriteHeader(statusCode)
	err := json.NewEncoder(resp).Encode(data)
	if err != nil {
		fmt.Println("Error encoding response %v", err)
	}
}

func SortMapByValue(wordFrequencyMap map[string]int) []model.WordFrequency {

	var wordsFrequencies []model.WordFrequency
	for key, value := range wordFrequencyMap {
		wordsFrequencies = append(wordsFrequencies, model.WordFrequency{Key: key, Value: value})
	}

	sort.Slice(wordsFrequencies, func(i, j int) bool {
		return wordsFrequencies[i].Value > wordsFrequencies[j].Value
	})

	if len(wordsFrequencies) > 10 {
		return wordsFrequencies[:10]
	}

	return wordsFrequencies
}
