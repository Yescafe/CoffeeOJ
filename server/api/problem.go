package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

func ProblemAdd(c *gin.Context) {
	var service service.ProblemAddService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Add()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func ProblemDelete(c *gin.Context) {
	var service service.ProblemDeleteService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Delete()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func ProblemFetch(c *gin.Context) {
	var service service.ProblemFetchService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Fetch()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func ProblemSubmit(c *gin.Context) {
	var service service.ProblemSubmitService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Submit()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
