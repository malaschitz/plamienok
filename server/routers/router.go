package routers

import (
	"github.com/labstack/echo"
	"github.com/malaschitz/plamienok/server/controllers"
)

func InitAccesible(e *echo.Echo) {
	//intro index.html
	e.File("/", "server/static/intro/index.html")
	e.Static("/assets", "server/static/intro/assets")

	//appka html
	e.File("/app", "server/static/app/dist/index.html")
	e.Static("/css", "server/static/app/dist/css")
	e.Static("/js", "server/static/app/dist/js")

	//public api
	e.GET("/api/info", controllers.Info)
	e.POST("/api/login", controllers.Login)
	e.GET("/api/stats", controllers.Stats)

}

func InitRestricted(e *echo.Group) {
	//api
	e.GET("/api/status", controllers.Status)
}
