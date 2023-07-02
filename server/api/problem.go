package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

// ProblemAdd godoc
//
// @Summary		add a new problem
// @Accept		json
// @Produce		json
// @Param		title		formData	string	true	"problem title"
// @param		memo_limit	formData	int		true	"problem memory limitation"
func ProblemAdd(c *gin.Context) {
	var serv service.ProblemAddService
	if err := c.ShouldBind(&serv); err == nil {
		user := CurrentUser(c)
		res := serv.Add(user.ID)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func ProblemDelete(c *gin.Context) {
	var serv service.ProblemDeleteService
	if err := c.ShouldBind(&serv); err == nil {
		res := serv.Delete()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func ProblemFetch(c *gin.Context) {
	var serv service.ProblemFetchService
	if err := c.ShouldBind(&serv); err == nil {
		res := serv.Fetch()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func ProblemSubmit(c *gin.Context) {
	var serv service.ProblemSubmitService
	if err := c.ShouldBind(&serv); err == nil {
		res := serv.Submit()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func ProblemList(c *gin.Context) {
	var serv service.ProblemListService
	if err := c.ShouldBind(&serv); err == nil {
		res := serv.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func ProblemUpdate(c *gin.Context) {
	var serv service.ProblemUpdateService
	if err := c.ShouldBind(&serv); err == nil {
		res := serv.Update()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
