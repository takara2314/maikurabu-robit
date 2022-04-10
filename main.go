package main

import (
	"maikurabu-robit/common"
	"maikurabu-robit/primary"
	"maikurabu-robit/secondary"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	robit := common.Robit{
		Primary: &common.RobitSession{
			Conn: nil,
			Stop: make(chan bool),
		},
		Secondary: &common.RobitSession{
			Conn: nil,
			Stop: make(chan bool),
		},
	}

	// Launch bots
	go primary.Start(&robit)
	go secondary.Start(&robit)

	// HTTP server routing
	router := gin.Default()

	router.GET("/", rootGET)
	router.GET("/version", rootGET)

	router.Run(":" + os.Getenv("PORT"))
}

func rootGET(c *gin.Context) {
	c.String(http.StatusOK, "Maikurabu Robit v1.4")
}
