package echo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//request struct for the echo post function
type request struct {
	Message string `json:"message" binding:"required"`
}

//response struct for the echo post function.
type response struct {
	Message string `json:"message"`
}

//login function does magic with the request struct and return the response struct.
func echo(request request) (response, error) {
	echo := response{"Success: " + request.Message}
	return echo, nil
}

// CreateEchoService initializes the post method to receive posts
func CreateEchoService(router *gin.Engine) {
	router.POST("/Echo", func(c *gin.Context) {

		var request request
		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, err.Error)
			return
		}

		response, err := echo(request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error)
			return
		}

		c.JSON(http.StatusOK, response)
	})
}
