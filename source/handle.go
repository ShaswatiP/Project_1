package source
import (
	"encoding/json"
	"net/http"
	"strings"
)

func CountFrequency(resp http.ResponseWriter, req *http.Request) {

	var textInputRequest model.TextInput

	err := json.NewDecoder(req.Body).Decode(&textInputRequest)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
	}

	textArray := strings.Split(textInputRequest.Text, " ")

	wordOccurrenceMap := make(map[string]int)

	for _, text := range textArray {
		wordOccurrenceMap[text] += 1
	}

	mostOccurrenceWordsMap := SortMapByValue(wordOccurrenceMap)

	EncodeJSONBody(resp, http.StatusOK, mostOccurrenceWordsMap)
}

func CallService(resp http.ResponseWriter, req *http.Request) {
	//make a HTTP Post Request
	res, err := http.Post("http://127.0.0.1:8080/api/text", "application/json; content=UTF-8", req.Body)
	if err != nil {
		logrus.Errorf("CallService: error in calling service: %v", err)
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	outputData, _ := ioutil.ReadAll(res.Body)

	err = res.Body.Close()
	if err != nil {
		logrus.Errorf("CallService: error in closing response body: %v", err)
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	//converting byte data into Struct
	var outputDataINJSON []models.WordFrequencyStruct
	err = json.Unmarshal(outputData, &outputDataINJSON)
	if err != nil {
		logrus.Errorf("CallService: error in converting to JSON: %v", err)
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	utils.EncodeJSONBody(resp, http.StatusOK, outputDataINJSON)
}