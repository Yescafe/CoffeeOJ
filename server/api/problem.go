package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

// ProblemAdd godoc
//
//	@Summary	add a new problem
//	@Accept		json
//	@Produce	json
//	@Param		title			query	string	true	"problem title"
//	@Param		memo_limit		query	int		true	"problem memory limitation"
//	@Param		time_limit		query	int		true	"problem time limitation"
//	@Param		text			query	string	true	"problem text"
//	@Router		/problem/add	[post]
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
//	@Accept		json
//	@Produce	json
//	@Param		id				query	int	true	"problem memory limitation"
//	@Router		/problem/delete	[post]
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
//	@Accept		json
//	@Produce	json
//	@Param		id				query	int	true	"problem id"
//	@Router		/problem/fetch	[get]
func ProblemFetch(c *gin.Context) {
	var serv service.ProblemFetchService
	if err := c.ShouldBind(&serv); err == nil {
		res := serv.Fetch()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ProblemSubmit godoc
//
//	@Summary	submit code
//	@Accept		json
//	@Produce	json
//	@Param		problem_id		query	int		true	"problem id"
//	@Param		source_code		query	string	true	"submission source code"
//	@Param		lang			query	string	true	"submission language"	Enums(c, cpp, rust, python)
//	@Router		/problem/submit	[post]
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
//	@Produce	json
//	@Router		/problem/list	[get]
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
//	@Accept		json
//	@Produce	json
//	@Param		id				query	int		true	"problem id"
//	@Param		title			query	string	true	"problem title"
//	@Param		memo_limit		query	int		true	"problem memory limitation"
//	@Param		time_limit		query	int		true	"problem time limitation"
//	@Param		text			query	string	true	"problem text"
//	@Router		/problem/update	[post]
func ProblemUpdate(c *gin.Context) {
	var serv service.ProblemUpdateService
	if err := c.ShouldBind(&serv); err == nil {
		res := serv.Update()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
