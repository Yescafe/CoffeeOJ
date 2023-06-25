package server

import (
	"os"
	"singo/api"
	"singo/middleware"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("ping", api.Ping)

		// 用户登录
		v1.POST("user/register", api.UserRegister)

		// 用户登录
		v1.POST("user/login", api.UserLogin)

		// 需要登录保护的
		auth := v1.Group("")
		auth.Use(middleware.AuthRequired())
		{
			// User Routing
			auth.GET("user/me", api.UserMe)
			auth.DELETE("user/logout", api.UserLogout)
			auth.POST("problem/submit", api.ProblemSubmit)
		}

		// 需要登录保护同时需要管理员
		adminAuth := v1.Group("")
		adminAuth.Use(middleware.AuthRequired())
		adminAuth.Use(middleware.AdminRequired())
		{
			// 添加题目
			adminAuth.POST("problem/add", api.ProblemAdd)
			// 删除题目
			adminAuth.POST("problem/delete", api.ProblemDelete)
		}

		v1.GET("problem/fetch", api.ProblemFetch)
	}
	return r
}
