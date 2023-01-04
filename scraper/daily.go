package scraper

// daily.go file

import (
	"log"
	"time"
	"zombeez/constant"
)

func WriteDaily(timeArg time.Time) (int, error) {
	scoreBoard := GetDaily(timeArg, false)
	i, err := writeFile(scoreBoard)
	if err != nil {
		log.Println(err)
	}

	return i, err
}

func GetDaily(timeArg time.Time, cacheBool bool) ScoreBoard {
	loc, err := time.LoadLocation(constant.TimeZone)
	if err != nil {
		log.Println(err)
	}

	seed := getSeedFromDate("daily", timeArg)
	var returnBoard ScoreBoard

	if cacheBool {
		cacheBoard := CheckCache("daily", seed, timeArg)
		// log.Println("CACHE DATE: " + cacheBoard.CreatedDate.String())

		year, month, day := timeArg.In(loc).Date()
		midnightDate := time.Date(year, month, day, 23, 59, 59, 0, loc)

		// If the score hasn't been checked since the score date has ended then go grab it via API

		if cacheBoard.CreatedDate.Before(midnightDate) {
			log.Println("API GET DAILY: " + timeArg.String())

			result := getScores("daily", seed)

			returnBoard = ScoreBoard{
				ScoreType:   "daily",
				Seed:        seed,
				ScoreDate:   timeArg.In(loc),
				CreatedDate: time.Now().In(loc),
				Scores:      result,
			}

			DailyBoards[seed] = returnBoard

			writeFile(returnBoard)

		} else {
			returnBoard = cacheBoard
		}
	} else {
		log.Println("API GET DAILY: " + timeArg.String())

		result := getScores("daily", seed)

		returnBoard = ScoreBoard{
			ScoreType:   "daily",
			Seed:        seed,
			ScoreDate:   timeArg.In(loc),
			CreatedDate: time.Now().In(loc),
			Scores:      result,
		}

		DailyBoards[seed] = returnBoard

		writeFile(returnBoard)

	}
	// fmt.Printf("%+v\n", returnBoard)

	return returnBoard
}

func GetYesterday() ScoreBoard {
	loc, err := time.LoadLocation(constant.TimeZone)
	if err != nil {
		log.Println(err)
	}
	scoreDate := time.Now().AddDate(0, 0, -1)

	scoreBoard := GetDaily(scoreDate.In(loc), true)

	log.Println(scoreBoard)

	return scoreBoard

}

func GetToday() ScoreBoard {
	loc, err := time.LoadLocation(constant.TimeZone)
	if err != nil {
		log.Println(err)
	}
	scoreDate := time.Now()
	var scoreBoard ScoreBoard

	// This one is special so we're goign to check if there's a cache before grabbing it from the API
	seed := getSeedFromDate("daily", scoreDate.In(loc))
	scoreBoard = CheckCache("daily", seed, time.Now())

	log.Println("CACHE DATE: " + scoreBoard.CreatedDate.String())

	// Don't do an API call if the cache is within 2 minutes

	if scoreBoard.CreatedDate.Add(2 * time.Minute).Before(scoreDate) {
		scoreBoard = GetDaily(scoreDate.In(loc), false)
		MakeTallyBoard(DailyBoards, WeeklyBoards)
		MakeCoolBoard(DailyBoards, WeeklyBoards)
	}

	DailyBoards[seed] = scoreBoard
	log.Println(scoreBoard)

	// End Cache Check
	return scoreBoard

}
