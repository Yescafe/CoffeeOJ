package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

func SubmissionFetch(c *gin.Context) {
	var service service.SubmissionFetchService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Fetch()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func SubmissionRejudge(c *gin.Context) {
	var service service.SubmissionRejudgeService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Rejudge()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
