package routers

import (
	"github.com/labstack/echo"
	"github.com/malaschitz/plamienok/server/controllers"
)

func InitAccesible(e *echo.Echo) {
	//intro index.html
	e.Static("/", "www/dist")
	//e.File("/", "www/dist/index.html")
	//e.File("/favicon.ico", "www/dist/favicon.ico")

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
	e.GET("/api/relationships", controllers.Relationships)
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
	e.GET("/api/persons/tasks", controllers.PersonsTasks)
	e.GET("/api/person/:id", controllers.Person)
	e.POST("/api/person", controllers.PersonPost)
	e.PUT("/api/person", controllers.PersonPut)
	e.DELETE("/api/person/relative/:id", controllers.PersonRelationDelete)
	e.POST("/api/person/relative/:id", controllers.PersonRelationPost)
	e.GET("/api/meniny/:date", controllers.Meniny)

	//sessions
	e.GET("/api/sessions", controllers.Sessions)
	e.POST("/api/session", controllers.SessionPost)

	//visits
	e.GET("/api/visits/:id", controllers.Visits)
	e.GET("/api/visitcall/:id", controllers.VisitCall)
	e.POST("/api/visitcall", controllers.VisitCallPost)
	e.GET("/api/visitcalls", controllers.VisitCalls)

	e.GET("/api/visithome/:id", controllers.VisitHome)
	e.POST("/api/visithome", controllers.VisitHomePost)
	e.GET("/api/visithomes", controllers.VisitHomes)

	//tasks
	e.GET("/api/tasks/:id", controllers.Ciele)
	e.POST("/api/task", controllers.CielPost)
}
