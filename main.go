package main

import (
	"net/http"
	"zombeez/scraper"

	"github.com/gin-gonic/gin"
)

func returnToday(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, scraper.GetToday())
}

func returnYesterday(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, scraper.GetYesterday())
}

func returnWeekly(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, scraper.GetTodayWeekly())
}

func returnHighScore(c *gin.Context) {

	dailyHS := scraper.FindHighScore(&scraper.DailyBoards)
	weeklyHS := scraper.FindHighScore(&scraper.WeeklyBoards)

	if weeklyHS.Scores[0].Score > dailyHS.Scores[0].Score {
		c.IndentedJSON(http.StatusOK, weeklyHS)

	} else {
		c.IndentedJSON(http.StatusOK, dailyHS)
	}

}

func returnTallyBoard(c *gin.Context) {

	tallyBoardSorted := scraper.FindHighTally(&scraper.PlayerTallyBoards)

	c.IndentedJSON(http.StatusOK, tallyBoardSorted)

}

func returnCoolBoard(c *gin.Context) {

	coolBoardSorted := scraper.FindHighCoolPoint(&scraper.PlayerCoolBoards)

	c.IndentedJSON(http.StatusOK, coolBoardSorted)

}

func return420(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, scraper.Get420())
}

func return69(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, scraper.Get69())
}

func main() {

	router := gin.Default()

	router.GET("/today", returnToday)
	router.GET("/yesterday", returnYesterday)
	router.GET("/weekly", returnWeekly)
	router.GET("/high", returnHighScore)
	router.GET("/tally", returnTallyBoard)
	router.GET("/coolpoints", returnCoolBoard)
	router.GET("/custom420420", return420)
	router.GET("/custom696969", return69)

	router.Static("/server", "./server")

	//pp.Println(scraper.FindHighScore(&scraper.DailyBoards))

	//pp.Println(scraper.ReturnCoolPoints(&scraper.FindHighScore(&scraper.DailyBoards).Scores[0]))
	router.Run("localhost:8080")

}
