package server

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
	"singo/api"
	"singo/middleware"
)

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

		users := v1.Group("/users")
		{
			users.GET(":id", api.UserFetch)
			//users.GET("", api.UserList)
			users.POST("login", api.UserLogin)
			users.POST("register", api.UserRegister)
			auth := users.Group("")
			auth.Use(middleware.AuthRequired())
			{
				auth.GET("me", api.UserMe)
				auth.DELETE("logout", api.UserLogout)
			}
		}

		problems := v1.Group("/problems")
		{
			problems.GET(":id", api.ProblemFetch)
			problems.GET("", api.ProblemList)
			auth := problems.Group("")
			auth.Use(middleware.AuthRequired())
			{
				auth.POST("submit", api.ProblemSubmit)
				admin := auth.Group("")
				admin.Use(middleware.AdminRequired())
				{
					admin.POST("add", api.ProblemAdd)
					admin.POST("delete", api.ProblemDelete)
					admin.POST("update", api.ProblemUpdate)
				}
			}
		}

		submissions := v1.Group("/submissions")
		{
			submissions.GET(":id", api.SubmissionFetch)
			submissions.POST("rejudge", api.SubmissionRejudge)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
