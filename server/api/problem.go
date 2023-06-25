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
