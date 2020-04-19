package main

import (
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"os"
)

func main() {
	// resp, _ := http.Get("https://api.covid19india.org/data.json")

	// Initiate covid array
	var names [][]string

	body, _ := ioutil.ReadFile("data.json")

	// // Declared an empty interface
	var result map[string]interface{}

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(body), &result)

	casesList := result["cases_time_series"].([]interface{})

	for currentDate := 0; currentDate < len(casesList); currentDate++ {
		matchObj := casesList[currentDate]
		matchMap := matchObj.(map[string]interface{})
		stats := []string{matchMap["date"].(string), matchMap["dailyconfirmed"].(string)}
		names = append(names, stats)
	}

	// Create file
	file, _ := os.Create("covidIndia.csv")

	defer file.Close()

	writer := csv.NewWriter(file)

	defer writer.Flush()

	for _, value := range names {
		writer.Write(value)
	}
}
