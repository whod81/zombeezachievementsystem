package scraper

// scraper.go file

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
	"zombeez/constant"
)

type ScoreBoards map[string]ScoreBoard

type Score struct {
	PlayerID   string `json:"playerID"`
	PlayerName string `json:"playerName"`
	Score      int    `json:"score"`
	Seed       string `json:"seed"`
}

type ScoreBoard struct {
	ScoreType   string
	Seed        string
	ScoreDate   time.Time
	CreatedDate time.Time
	Scores      []Score
}

var DailyBoards ScoreBoards
var WeeklyBoards ScoreBoards

// This is where I create a global data structure where i can grab stuff rather than

func init() {
	MakeBoards()
	// This has to go second
	MakeTallyBoard(DailyBoards, WeeklyBoards)
	MakeCoolBoard(DailyBoards, WeeklyBoards)

}

func MakeBoards() {
	loc, err := time.LoadLocation(constant.TimeZone)
	if err != nil {
		log.Println(err)
	}

	DailyBoards = make(ScoreBoards)
	WeeklyBoards = make(ScoreBoards)

	var dayBoard ScoreBoard
	var weekBoard ScoreBoard

	year, month, day := time.Now().In(loc).Date()
	midnightToday := time.Date(year, month, day, 11, 59, 59, 0, loc)

	//pastDate := midnightToday.AddDate(0, 0, -89)
	// This is when the newest release came out
	pastDate := time.Date(2022, 11, 22, 23, 59, 59, 0, loc)
	year, month, day = pastDate.In(loc).Date()
	pastDate = time.Date(year, month, day, 0, 0, 0, 0, loc)

	for pastDate.Before(midnightToday) {
		dayBoard = GetDaily(pastDate, true)
		DailyBoards[dayBoard.Seed] = dayBoard
		if pastDate.Weekday() == 0 {
			weekBoard = GetWeekly(pastDate, true)
			WeeklyBoards[weekBoard.Seed] = weekBoard
			fmt.Println(weekBoard)
		}
		fmt.Println(dayBoard)
		pastDate = pastDate.AddDate(0, 0, 1)
	}

}

func getFilename(scoreType string, seed string, scoreDate time.Time) string {
	directory := scoreType + "/" + scoreDate.Format("20060102")
	if scoreType == "custom" {
		directory = scoreType
	}
	fileName := seed + ".json"

	return directory + "/" + fileName

}

func getSeedFromDate(ScoreType string, ScoreDate time.Time) string {
	var dateStr string

	if ScoreType == "weekly" {
		// Get day of the week and subtract it to get Sunday's Date
		sundayDate := findSunday(ScoreDate)

		// BB API adds 13 to the month for weekly but we can't use date math for that

		dateStr = sundayDate.Format("0206")
		monthPlus := int(sundayDate.Month()) + 13
		dateStr = strconv.Itoa(monthPlus) + dateStr
	} else {
		dateStr = ScoreDate.Format("10206")
	}

	return dateStr
}

func getDateFromSeed(ScoreType string, seed string) time.Time {
	month := seed[:len(seed)-4]
	monthInt, _ := strconv.Atoi(month)
	if monthInt > 12 {
		// Must be a weekly
		monthInt = monthInt - 13
		month = strconv.Itoa(monthInt)
	}
	t, err := time.Parse("1/02/06 03:04PM", month+"/"+seed[2:4]+"/"+seed[4:6]+" 12:00PM")
	if err != nil {
		log.Println(err)
	}

	loc, _ := time.LoadLocation(constant.TimeZone)

	return t.In(loc)
}

// Return # of array entries
func writeFile(scoreBoard ScoreBoard) (int, error) {
	e, err := json.Marshal(scoreBoard)
	if err != nil {
		log.Println(err)
		return 1, errors.New("could not json encode")
	}

	fileName := getFilename(scoreBoard.ScoreType, scoreBoard.Seed, scoreBoard.ScoreDate)

	// Create directories
	err = os.MkdirAll(filepath.Dir(fileName), os.ModePerm)
	if err != nil {
		log.Println(err)
	}

	// Write files
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("JSON WRITE: " + fileName)
	defer f.Close()
	_, err = f.WriteString(string(e))

	if err != nil {
		log.Fatal(err)
	}

	return len(scoreBoard.Scores), err

}

func CheckCache(scoreType string, seed string, timeArg time.Time) ScoreBoard {
	loc, err := time.LoadLocation(constant.TimeZone)
	if err != nil {
		log.Println(err)
	}

	var cacheBoard ScoreBoard

	if scoreType == "weekly" {
		timeArg = findSunday(timeArg)
	}

	fileName := getFilename(scoreType, seed, timeArg.In(loc))

	// Read in JSON File with scores -- if it doesn't exist, is corrupt or is older than 2 minutes go hit the API
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Println("CACHE CHECK: " + err.Error())
		return cacheBoard
	} else {
		err = json.Unmarshal(content, &cacheBoard)

		if err != nil {
			log.Println(err)
		}

		return cacheBoard

	}

}

func FindHighScore(dayBoards *ScoreBoards) ScoreBoard {
	loc, err := time.LoadLocation(constant.TimeZone)
	if err != nil {
		log.Println(err)
	}

	highScore := ScoreBoard{
		ScoreType:   "none",
		Seed:        "",
		ScoreDate:   time.Date(1970, 1, 0, 0, 0, 0, 0, loc),
		CreatedDate: time.Date(1970, 1, 0, 0, 0, 0, 0, loc),
		Scores: []Score{
			{
				PlayerID:   "",
				PlayerName: "",
				Score:      0,
				Seed:       "",
			},
		},
	}

	for _, value := range *dayBoards {
		if len(value.Scores) > 0 {
			if value.Scores[0].Score > highScore.Scores[0].Score {
				highScore = value
			}
		}
	}

	// I don't need any of the scores from that board except the top one

	if len(highScore.Scores) > 0 {
		highScore.Scores = highScore.Scores[:1]
	}

	return highScore
}
