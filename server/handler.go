package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleGetRandomJoke(app App, log Logger) func(c *gin.Context) {
	return func(c *gin.Context) {
		result, err := app.FetchRandomNameJoke()
		if err != nil {
			log.Err(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, result)
	}
}
