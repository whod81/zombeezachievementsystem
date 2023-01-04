package scraper

// coolpoint.go file

import (
	"log"
	"regexp"
	"sort"
	"strconv"
)

type CoolBoard struct {
	PlayerID   string `json:"playerID"`
	PlayerName string `json:"playerName"`
	CoolPoints int    `json:"coolpoints"`
}

type CoolBoards map[string]CoolBoard

var PlayerCoolBoards CoolBoards

// This is where I create a global data structure where i can grab stuff rather than

func MakeCoolBoard(dayBoards ScoreBoards, weeklyBoards ScoreBoards) {

	PlayerCoolBoards = make(CoolBoards)

	// Rather than use the TallyBoard I decided to duplicate things in memory for CoolBoard
	// TODO take a nil argument for either of the boards and skip processing on that one

	// Rather than iterate directly through the Boards I need to do it in reverse so the newer playerNames are used
	// when the structures are initiated.  i simply grab the keys of the map and reverse them and then use them as index.

	dayKeys := make([]string, 0, len(DailyBoards))
	for k := range DailyBoards {
		dayKeys = append(dayKeys, k)
	}

	sort.Sort(sort.Reverse(sort.StringSlice(dayKeys)))

	for _, k := range dayKeys {
		for _, score := range DailyBoards[k].Scores {
			coolPoints := returnCoolPoints(score.Score)
			if coolPoints > 0 {
				log.Printf("Cool Point (%d) Added for %s -- %d -- %s", coolPoints, score.PlayerName, score.Score, DailyBoards[k].ScoreDate)
				coolPointPlayer(coolPoints, score.PlayerID, score.PlayerName)

			}
		}
	}

	weekKeys := make([]string, 0, len(weeklyBoards))
	for k := range weeklyBoards {
		weekKeys = append(weekKeys, k)
	}

	sort.Sort(sort.Reverse(sort.StringSlice(weekKeys)))

	for _, k := range weekKeys {
		for _, score := range weeklyBoards[k].Scores {
			coolPoints := returnCoolPoints(score.Score)
			if coolPoints > 0 {
				log.Printf("Cool Point (%d) Added for %s -- %d -- %s", coolPoints, score.PlayerName, score.Score, weeklyBoards[k].ScoreDate)
				coolPointPlayer(coolPoints, score.PlayerID, score.PlayerName)

			}

		}
	}

}

func coolPointPlayer(points int, playerID string, PlayerName string) {

	if _, ok := PlayerCoolBoards[playerID]; !ok {
		PlayerCoolBoards[playerID] = CoolBoard{
			PlayerID:   playerID,
			PlayerName: PlayerName,
			CoolPoints: 0,
		}
	}

	PlayerCoolBoard := PlayerCoolBoards[playerID]

	PlayerCoolBoard.CoolPoints = PlayerCoolBoards[playerID].CoolPoints + points

	PlayerCoolBoards[playerID] = PlayerCoolBoard
}

func returnCoolPoints(score int) int {
	bs := []byte(strconv.Itoa(score))

	re := regexp.MustCompile(`1337|69|420`)

	return len(re.FindAll(bs, -1))
}

func FindHighCoolPoint(PlayerCoolBoards *CoolBoards) []CoolBoard {
	// TODO THere's such better ways to do this than create a new dataset just for sorting and returning

	CB := make([]CoolBoard, 0)
	for _, value := range *PlayerCoolBoards {
		CB = append(CB, value)
	}

	sort.Slice(CB, func(i, j int) bool {
		return CB[i].CoolPoints > CB[j].CoolPoints
	})

	return CB

}
