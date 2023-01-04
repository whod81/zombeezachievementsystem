package scraper

// tally.go file

import (
	"log"
	"sort"
	"time"
	"zombeez/constant"
)

type TallyBoard struct {
	PlayerID   string `json:"playerID"`
	PlayerName string `json:"playerName"`
	Score      int    `json:"score"`
	Tally      Tally  `json:"tally"`
}

type Tally struct {
	Weekly1st  uint16 `json:"weekly1st"`
	Weekly2nd  uint16 `json:"weekly2nd"`
	Weekly3rd  uint16 `json:"weekly3rd"`
	Weeklypart uint16 `json:"weeklypart"`
	Daily1st   uint16 `json:"daily1st"`
	Daily2nd   uint16 `json:"daily2nd"`
	Daily3rd   uint16 `json:"daily3rd"`
	Dailypart  uint16 `json:"dailypart"`
}

type TallyBoards map[string]TallyBoard

var PlayerTallyBoards TallyBoards

// This is where I create a global data structure where i can grab stuff rather than

func MakeTallyBoard(dayBoards ScoreBoards, weeklyBoards ScoreBoards) {
	loc, err := time.LoadLocation(constant.TimeZone)
	if err != nil {
		log.Println(err)
	}

	PlayerTallyBoards = make(TallyBoards)
	year, month, day := time.Now().In(loc).AddDate(0, 0, -constant.TailyDays).Date()
	tallyDate := time.Date(year, month, day, 0, 0, 0, 0, loc)

	// Rather than iterate directly through the Boards I need to do it in reverse so the newer playerNames are used
	// when the structures are initiated.  i simply grab the keys of the map and reverse them and then use them as index.

	dayKeys := make([]string, 0, len(DailyBoards))
	for k := range DailyBoards {
		dayKeys = append(dayKeys, k)
	}

	sort.Sort(sort.Reverse(sort.StringSlice(dayKeys)))

	for _, k := range dayKeys {
		if DailyBoards[k].ScoreDate.After(tallyDate) {
			for place, score := range DailyBoards[k].Scores {

				tallyPlayer("daily", place, score.PlayerID, score.PlayerName, score.Score)

			}

		}

	}

	weekKeys := make([]string, 0, len(weeklyBoards))
	for k := range weeklyBoards {
		weekKeys = append(weekKeys, k)
	}

	sort.Sort(sort.Reverse(sort.StringSlice(weekKeys)))

	for _, k := range weekKeys {
		if weeklyBoards[k].ScoreDate.After(tallyDate) {
			for place, score := range weeklyBoards[k].Scores {

				tallyPlayer("weekly", place, score.PlayerID, score.PlayerName, score.Score)
			}

		}

	}

}

func tallyPlayer(scoreType string, place int, playerID string, PlayerName string, score int) {

	if _, ok := PlayerTallyBoards[playerID]; !ok {
		PlayerTallyBoards[playerID] = TallyBoard{
			PlayerID:   playerID,
			Score:      0,
			PlayerName: PlayerName,
		}
	}

	PlayerTallyBoard := PlayerTallyBoards[playerID]

	if scoreType == "daily" {
		switch place {
		case 0:
			PlayerTallyBoard.Score = PlayerTallyBoards[playerID].Score + 5
			PlayerTallyBoard.Tally.Daily1st = PlayerTallyBoards[playerID].Tally.Daily1st + 1
		case 1:
			PlayerTallyBoard.Score = PlayerTallyBoards[playerID].Score + 3
			PlayerTallyBoard.Tally.Daily2nd = PlayerTallyBoards[playerID].Tally.Daily2nd + 1

		case 2:
			PlayerTallyBoard.Score = PlayerTallyBoards[playerID].Score + 2
			PlayerTallyBoard.Tally.Daily3rd = PlayerTallyBoards[playerID].Tally.Daily3rd + 1

		default:
			PlayerTallyBoard.Score = PlayerTallyBoards[playerID].Score + 1
			PlayerTallyBoard.Tally.Dailypart = PlayerTallyBoards[playerID].Tally.Dailypart + 1

		}
	}

	if scoreType == "weekly" {
		switch place {
		case 0:
			PlayerTallyBoard.Score = PlayerTallyBoards[playerID].Score + 10
			PlayerTallyBoard.Tally.Weekly1st = PlayerTallyBoards[playerID].Tally.Weekly1st + 1
		case 1:
			PlayerTallyBoard.Score = PlayerTallyBoards[playerID].Score + 6
			PlayerTallyBoard.Tally.Weekly2nd = PlayerTallyBoards[playerID].Tally.Weekly2nd + 1

		case 2:
			PlayerTallyBoard.Score = PlayerTallyBoards[playerID].Score + 4
			PlayerTallyBoard.Tally.Weekly3rd = PlayerTallyBoards[playerID].Tally.Weekly3rd + 1

		default:
			PlayerTallyBoard.Score = PlayerTallyBoards[playerID].Score + 2
			PlayerTallyBoard.Tally.Weeklypart = PlayerTallyBoards[playerID].Tally.Weeklypart + 1

		}
	}

	PlayerTallyBoards[playerID] = PlayerTallyBoard
}

func FindHighTally(PlayerTallyBoards *TallyBoards) []TallyBoard {
	// TODO THere's such better ways to do this than create a new dataset just for sorting and returning

	TB := make([]TallyBoard, 0)
	for _, value := range *PlayerTallyBoards {
		TB = append(TB, value)
	}

	sort.Slice(TB, func(i, j int) bool {
		return TB[i].Score > TB[j].Score
	})

	return TB

}
