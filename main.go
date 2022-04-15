package main

import (
	"maikurabu-robit/common"
	"maikurabu-robit/primary"
	"maikurabu-robit/secondary"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	common.RobitState = common.Robit{
		Version: "v1.4",
		Primary: &common.RobitSession{
			AppID: os.Getenv("PRIMARY_BOT_ID"),
			Token: os.Getenv("PRIMARY_BOT_TOKEN"),
			Conn:  nil,
			Stop:  make(chan bool),
		},
		Secondary: &common.RobitSession{
			AppID: os.Getenv("SECONDARY_BOT_ID"),
			Token: os.Getenv("SECONDARY_BOT_TOKEN"),
			Conn:  nil,
			Stop:  make(chan bool),
		},
		Start: &common.StartProcess{
			MinVoter:            3,
			VotePeriod:          3 * time.Minute,
			ReactionCheckPeriod: 3 * time.Second,
			StopVote:            make(chan bool),
		},
		MaxClassmateNum:     83,
		ActivityInterval:    10 * time.Second,
		AutoClosingWaitTime: 15 * time.Minute,
	}

	// Launch bots
	go primary.Start()
	go secondary.Start()

	// HTTP server routing
	router := gin.Default()

	router.GET("/", rootGET)
	router.GET("/version", rootGET)

	router.Run(":" + os.Getenv("PORT"))
}

func rootGET(c *gin.Context) {
	c.String(http.StatusOK, "Maikurabu Robit "+common.RobitState.Version)
}
