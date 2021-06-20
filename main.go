package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", homeGET)
	r.GET("/version", homeGET)

	r.Run(":" + os.Getenv("PORT"))
}

func homeGET(c *gin.Context) {
	c.String(http.StatusOK, "Maikurabu Robit v1.2.1")
}
