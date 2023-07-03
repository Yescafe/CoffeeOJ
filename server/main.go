package main

import (
	"singo/conf"
	_ "singo/docs"
	"singo/server"
)

// @title						CoffeeOJ RUSTful APIs
// @version					1.0
// @contact.name				Ivan Chien
// @contact.email				qyc027@gmail.com
// @license.name				GPL-3.0 License
// @license.url				https://www.gnu.org/licenses/gpl-3.0.html
// @host						127.0.0.1:3000
// @BasePath					/api/v1
// @securityDefinitions.apikey	SetCookie
// @in							header
// @name						Cookie
// @description				cookie
func main() {
	// 从配置文件读取配置
	conf.Init()

	// 装载路由
	r := server.NewRouter()
	r.Run(":3000")
}
