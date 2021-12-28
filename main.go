package main

import (
	"maikurabu-robit/processes"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", homeGET)
	r.GET("/version", homeGET)
	r.GET("/bc", bcGET)

	r.Run(":" + os.Getenv("PORT"))
}

func homeGET(c *gin.Context) {
	c.String(http.StatusOK, "Maikurabu Robit v1.3")
}

func bcGET(c *gin.Context) {
	if c.Query("password") == os.Getenv("BC_PASSWORD") {
		_ = processes.BoardCast(discord, c.Query("message"))
		c.String(http.StatusOK, "OK")
		return
	}

	c.String(http.StatusUnauthorized, "401 Unauthorized")
}
