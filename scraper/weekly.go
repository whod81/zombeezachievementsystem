package scraper

// weekly.go file

import (
	"log"
	"time"
	"zombeez/constant"
)

func findSunday(timeArg time.Time) time.Time {
	// Get day of the week and subtract it to get Sunday's Date

	dayNumber := int(timeArg.Weekday())
	sundayDate := timeArg.AddDate(0, 0, -dayNumber)

	return sundayDate

}

func WriteWeekly(timeArg time.Time) (int, error) {
	scoreBoard := GetWeekly(timeArg, false)
	i, err := writeFile(scoreBoard)
	if err != nil {
		log.Println(err)
	}

	return i, err
}

func GetWeekly(timeArg time.Time, cacheBool bool) ScoreBoard {
	loc, err := time.LoadLocation(constant.TimeZone)
	if err != nil {
		log.Println(err)
	}

	seed := getSeedFromDate("weekly", timeArg)
	var returnBoard ScoreBoard

	sundayDate := getDateFromSeed("weekly", seed)

	if cacheBool {

		cacheBoard := CheckCache("weekly", seed, timeArg.In(loc))
		//		log.Println("CACHE DATE: " + cacheBoard.CreatedDate.String())

		midnightDate := sundayDate.AddDate(0, 0, 7)

		if cacheBoard.CreatedDate.Before(midnightDate.In(loc)) {
			log.Println("API GET WEEKLY: " + timeArg.In(loc).String())
			result := getScores("weekly", seed)

			returnBoard = ScoreBoard{
				ScoreType:   "weekly",
				Seed:        seed,
				ScoreDate:   sundayDate.In(loc),
				CreatedDate: time.Now().In(loc),
				Scores:      result,
			}

			WeeklyBoards[seed] = returnBoard
			writeFile(returnBoard)

		} else {
			returnBoard = cacheBoard
		}
		// fmt.Printf("%+v\n", returnBoard)
	} else {
		log.Println("API GET WEEKLY: " + timeArg.In(loc).String())
		result := getScores("weekly", seed)

		returnBoard = ScoreBoard{
			ScoreType:   "weekly",
			Seed:        seed,
			ScoreDate:   sundayDate.In(loc),
			CreatedDate: time.Now().In(loc),
			Scores:      result,
		}

		WeeklyBoards[seed] = returnBoard

		writeFile(returnBoard)

	}
	return returnBoard
}

func GetTodayWeekly() ScoreBoard {
	loc, err := time.LoadLocation(constant.TimeZone)
	if err != nil {
		log.Println(err)
	}

	todayDate := time.Now()
	//var todayBoard ScoreBoard

	// This one is special so we're goign to check if there's a cache before grabbing it from the API
	seed := getSeedFromDate("weekly", todayDate.In(loc))
	todayBoard := CheckCache("weekly", seed, time.Now())

	log.Println("CACHE DATE: " + todayBoard.CreatedDate.String())

	// Don't do an API call if the cache is within 2 minutes

	if todayBoard.CreatedDate.Add(2 * time.Minute).Before(todayDate) {
		todayBoard = GetWeekly(time.Now().In(loc), false)

		MakeTallyBoard(DailyBoards, WeeklyBoards)
		MakeCoolBoard(DailyBoards, WeeklyBoards)

	}

	WeeklyBoards[seed] = todayBoard
	log.Println(todayBoard)

	return todayBoard
}
