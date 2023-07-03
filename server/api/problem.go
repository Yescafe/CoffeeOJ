package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
	"strconv"
)

// ProblemAdd godoc
//
//	@Summary	add a new problem
//	@Tags		problems
//	@Accept		json
//	@Produce	json
//	@Param		title			query	string	true	"problem title"
//	@Param		memo_limit		query	int		true	"problem memory limitation"
//	@Param		time_limit		query	int		true	"problem time limitation"
//	@Param		text			query	string	true	"problem text"
//	@Router		/problems/add	[post]
//	@Security	SetCookie
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

// ProblemDelete godoc
//
//	@Summary	delete a problem
//	@Tags		problems
//	@Accept		json
//	@Produce	json
//	@Param		id					query	int	true	"problem memory limitation"
//	@Router		/problems/delete	[post]
//	@Security	SetCookie
func ProblemDelete(c *gin.Context) {
	var serv service.ProblemDeleteService
	if err := c.ShouldBind(&serv); err == nil {
		res := serv.Delete()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ProblemFetch godoc
//
//	@Summary	fetch an existed problem
//	@Tags		problems
//	@Accept		json
//	@Produce	json
//	@Param		id				path	int	true	"problem id"
//	@Router		/problems/{id}	[get]
func ProblemFetch(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, ErrorResponse(err))
		return
	}
	serv := service.ProblemFetchService{
		ID: uint(id),
	}
	res := serv.Fetch()
	c.JSON(200, res)
}

// ProblemSubmit godoc
//
//	@Summary	submit code
//	@Tags		problems, submissions
//	@Accept		json
//	@Produce	json
//	@Param		problem_id			query	int		true	"problem id"
//	@Param		source_code			query	string	true	"submission source code"
//	@Param		lang				query	string	true	"submission language"	Enums(c, cpp, rust, python)
//	@Router		/problems/submit	[post]
//	@Security	SetCookie
func ProblemSubmit(c *gin.Context) {
	var serv service.ProblemSubmitService
	if err := c.ShouldBind(&serv); err == nil {
		res := serv.Submit()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ProblemList godoc
//
//	@Summary	list problems
//	@Tags		problems
//	@Produce	json
//	@Router		/problems	[get]
func ProblemList(c *gin.Context) {
	var serv service.ProblemListService
	if err := c.ShouldBind(&serv); err == nil {
		res := serv.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ProblemUpdate godoc
//
//	@Summary	update a problem
//	@Tags		problems
//	@Accept		json
//	@Produce	json
//	@Param		id					query	int		true	"problem id"
//	@Param		title				query	string	true	"problem title"
//	@Param		memo_limit			query	int		true	"problem memory limitation"
//	@Param		time_limit			query	int		true	"problem time limitation"
//	@Param		text				query	string	true	"problem text"
//	@Router		/problems/update	[post]
//	@Security	SetCookie
func ProblemUpdate(c *gin.Context) {
	var serv service.ProblemUpdateService
	if err := c.ShouldBind(&serv); err == nil {
		res := serv.Update()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
