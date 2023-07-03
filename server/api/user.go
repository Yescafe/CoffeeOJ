package api

import (
	"singo/serializer"
	"singo/service"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserRegister godoc
//
//	@Summary	login
//	@Tags		users
//	@Accept		json
//	@Produce	json
//	@Param		nickname			query	string	true	"nickname"
//	@Param		user_name			query	string	true	"username"
//	@Param		password			query	string	true	"password"
//	@Param		password_confirm	query	string	true	"confirmation password"
//	@Router		/users/register																																																																																																																																																				[post]
func UserRegister(c *gin.Context) {
	var serv service.UserRegisterService
	if err := c.ShouldBind(&serv); err == nil {
		res := serv.Register()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserLogin godoc
//
//	@Summary	login
//	@Tags		users
//	@Accept		json
//	@Produce	json
//	@Param		user_name		query	string	true	"username"
//	@Param		password		query	string	true	"password"
//	@Router		/users/login	[post]
func UserLogin(c *gin.Context) {
	var serv service.UserLoginService
	if err := c.ShouldBind(&serv); err == nil {
		res := serv.Login(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserMe godoc
//
//	@Summary	fetch user info
//	@Tags		users
//	@Accept		json
//	@Produce	json
//	@Router		/users/me	[get]
//	@Security	SetCookie
func UserMe(c *gin.Context) {
	user := CurrentUser(c)
	res := serializer.BuildUserResponse(*user)
	c.JSON(200, res)
}

// UserLogout godoc
//
//	@Summary	current user logout
//	@Tags		users
//	@Produce	json
//	@Router		/users/logout	[delete]
//	@Security	SetCookie
func UserLogout(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	s.Save()
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "登出成功",
	})
}

// UserFetch godoc
//
//	@Summary	fetch user info
//	@Tags		users
//	@Accept		json
//	@Produce	json
//	@Param		id			path	int	true	"user ID"
//	@Router		/users/{id}	[get]
func UserFetch(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, ErrorResponse(err))
		return
	}
	serv := service.UserFetchService{
		ID: uint(id),
	}
	res := serv.Fetch()
	c.JSON(200, res)
}
