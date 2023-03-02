package scraper

// custom.go file

import (
	"log"
	"time"
	"zombeez/constant"
)

func GetCustom(seed string) ScoreBoard {
	loc, err := time.LoadLocation(constant.TimeZone)
	if err != nil {
		log.Println(err)
	}

	timeArg := time.Now().In(loc)

	var returnBoard ScoreBoard

	// NOT DOING ANY CACHE CHECKING -- Letting calling routine do it

	log.Println("API GET CUSTOM: " + timeArg.String())

	result := getScores("custom", seed)

	returnBoard = ScoreBoard{
		ScoreType:   "custom",
		Seed:        seed,
		ScoreDate:   timeArg.In(loc),
		CreatedDate: time.Now().In(loc),
		Scores:      result,
	}

	writeFile(returnBoard)

	// fmt.Printf("%+v\n", returnBoard)

	return returnBoard
}

func Get420() ScoreBoard {
	scoreDate := time.Now()
	var scoreBoard ScoreBoard

	// This one is special so we're goign to check if there's a cache before grabbing it from the API
	seed := "420420"
	scoreBoard = CheckCache("custom", seed, time.Now())

	log.Println("CACHE DATE: " + scoreBoard.CreatedDate.String())

	// Don't do an API call if the cache is within 2 Hours

	if scoreBoard.CreatedDate.Add(2 * time.Hour).Before(scoreDate) {
		scoreBoard = GetCustom(seed)

	}

	log.Println(scoreBoard)

	return scoreBoard

}

func Get69() ScoreBoard {
	scoreDate := time.Now()
	var scoreBoard ScoreBoard

	// This one is special so we're goign to check if there's a cache before grabbing it from the API
	seed := "696969"
	scoreBoard = CheckCache("custom", seed, time.Now())

	log.Println("CACHE DATE: " + scoreBoard.CreatedDate.String())

	// Don't do an API call if the cache is within 2 Hours

	if scoreBoard.CreatedDate.Add(2 * time.Hour).Before(scoreDate) {
		scoreBoard = GetCustom(seed)

	}

	log.Println(scoreBoard)

	return scoreBoard

}

func Gett007() ScoreBoard {
	scoreDate := time.Now()
	var scoreBoard ScoreBoard

	// This one is special so we're goign to check if there's a cache before grabbing it from the API
	seed := "007007"
	scoreBoard = CheckCache("custom", seed, time.Now())

	log.Println("CACHE DATE: " + scoreBoard.CreatedDate.String())

	// Don't do an API call if the cache is within 2 Hours

	if scoreBoard.CreatedDate.Add(2 * time.Hour).Before(scoreDate) {
		scoreBoard = GetCustom(seed)

	}

	log.Println(scoreBoard)

	return scoreBoard

}
