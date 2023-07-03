package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
	"strconv"
)

// SubmissionFetch godoc
//
//	@Summary	submit code
//	@Tags		submissions
//	@Accept		json
//	@Produce	json
//	@Param		id					path	int	true	"submission id"
//	@Router		/submissions/{id}	[post]
func SubmissionFetch(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, ErrorResponse(err))
		return
	}
	serv := service.SubmissionFetchService{
		ID: uint(id),
	}
	res := serv.Fetch()
	c.JSON(200, res)
}

// SubmissionRejudge godoc
//
//	@Summary	rejudge submission
//	@Tags		submissions
//	@Accept		json
//	@Produce	json
//	@Param		id						query	int	true	"submission id"
//	@Router		/submissions/rejudge	[post]
//	@Security	SetCookie
func SubmissionRejudge(c *gin.Context) {
	var serv service.SubmissionRejudgeService
	if err := c.ShouldBind(&serv); err == nil {
		res := serv.Rejudge()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
