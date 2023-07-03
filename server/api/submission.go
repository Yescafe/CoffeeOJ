package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

// SubmissionFetch godoc
//
//	@Summary	submit code
//	@Accept		json
//	@Produce	json
//	@Param		id					query	int	true	"submission id"
//	@Router		/submission/fetch	[post]
func SubmissionFetch(c *gin.Context) {
	var serv service.SubmissionFetchService
	if err := c.ShouldBind(&serv); err == nil {
		res := serv.Fetch()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// SubmissionRejudge godoc
//
//	@Summary	rejudge submission
//	@Accept		json
//	@Produce	json
//	@Param		id					query	int	true	"submission id"
//	@Router		/submission/rejudge	[post]
func SubmissionRejudge(c *gin.Context) {
	var serv service.SubmissionRejudgeService
	if err := c.ShouldBind(&serv); err == nil {
		res := serv.Rejudge()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
