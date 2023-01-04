package scraper

// getscore.go file

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
	"zombeez/constant"
)

var apiLastRequest time.Time

func getScores(scoreType string, seed string) []Score {
	scoreTypes := map[string]string{
		"daily":  "1",
		"weekly": "3",
		"random": "4",
		"custom": "5",
	}

	d := time.Since(apiLastRequest)
	th, _ := time.ParseDuration(constant.Throttle)
	if d < th {
		sleepAmount := th - d
		log.Printf("Throttling API Call for %d milliseconds", sleepAmount.Milliseconds())
		time.Sleep(sleepAmount)
	}

	response, err := http.Get("http://api.bumblebase.com/v1/scores/" + seed + "/" + scoreTypes[scoreType] + "/10/694206942069/6")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var result []Score
	if err := json.Unmarshal(responseData, &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Print(err.Error())
		os.Exit(1)
	}

	apiLastRequest = time.Now()

	return result
}
