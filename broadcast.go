package main

import (
	"maikurabu-robit/common"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func broadcastGET(c *gin.Context) {
	pass := c.Query("password")
	channelID := c.Param("channel")
	message := c.Param("message")

	if pass != os.Getenv("BC_PASSWORD") {
		c.AbortWithStatusJSON(401, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	message = strings.ReplaceAll(message, "<br>", "\n")

	_, err := common.RobitState.Primary.Conn.ChannelMessageSend(
		channelID,
		message,
	)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"status": "ok",
	})
}
