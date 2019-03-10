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
	e.POST("/api/forgot", controllers.ForgottenPassword)
	e.POST("/api/checkCode6", controllers.CheckCode6)
	e.GET("/api/stats", controllers.Stats)

	//ciselniky
	e.GET("/api/meds", controllers.Meds)
	e.GET("/api/diagnoses", controllers.Diagnoses)
	e.GET("/api/diagnosesAll", controllers.DiagnosesAll)
}

func InitRestricted(e *echo.Group) {
	//api
	e.GET("/api/status", controllers.Status)

	//users
	e.GET("/api/users", controllers.Users)
	e.POST("/api/user", controllers.UserPost)
	e.DELETE("/api/user/:id", controllers.UserDelete)

	//cars
	e.GET("/api/cars", controllers.Cars)
	e.POST("/api/car", controllers.CarPost)
	e.DELETE("/api/car/:id", controllers.CarDelete)

	//persons
	e.POST("/api/persons", controllers.Persons)
	e.GET("/api/persons", controllers.PersonsAll)
	e.GET("/api/person/:id", controllers.Person)
	e.POST("/api/person", controllers.PersonPost)
	e.PUT("/api/person", controllers.PersonPut)

	//sessions
	e.GET("/api/sessions", controllers.Sessions)
	e.POST("/api/session", controllers.SessionPost)
}
