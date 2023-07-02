package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

func SubmissionFetch(c *gin.Context) {
	var serv service.SubmissionFetchService
	if err := c.ShouldBind(&serv); err == nil {
		res := serv.Fetch()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func SubmissionRejudge(c *gin.Context) {
	var serv service.SubmissionRejudgeService
	if err := c.ShouldBind(&serv); err == nil {
		res := serv.Rejudge()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
