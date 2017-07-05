package main

import (
	"todo-app/controllers"
	"todo-app/models"
	"todo-app/models/connection"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

var connectionString = "root:root@tcp(localhost:3306)/todoapp?charset=utf8&parseTime=true"

func main() {
	initConnection()
	e := echo.New()
	initRoutes(e)
	e.Start(":1234")
	defer connection.DB.Close()
}

func initConnection() {
	con, err := gorm.Open("mysql", connectionString)
	if err != nil {
		log.Error(err.Error())
		panic(err)
	}

	// check if connection was really opened
	er := con.DB().Ping()
	if er != nil {
		panic(er)
	}

	con.DropTableIfExists(&models.Task{}, &models.User{})
	con.AutoMigrate(&models.Task{}, &models.User{})

	connection.DB = con
}

func initRoutes(e *echo.Echo) {
	e.Static("/", "views")
	e.GET("/api/tasks", controllers.GetTasks)
	e.GET("/api/tasks/:id", controllers.GetTask)
	e.POST("/api/tasks", controllers.CreateTask)
	e.POST("/api/tasks/:id", controllers.UpdateTask)
	e.DELETE("/api/tasks/:id", controllers.DeleteTask)
	// return index while other routes not matching
	e.Static("/dist", "./views/dist")
	e.File("*", "./views")
}
