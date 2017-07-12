package main

import (
	"os"
	"todo-app/controllers"
	"todo-app/middlewares"
	"todo-app/models"
	"todo-app/models/connection"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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

	migrateUser()
}

func migrateUser() {
	user := models.User{
		Name:     "Test User",
		Username: "admin",
		Password: "admin"}

	if res := connection.DB.Create(&user); res.Error != nil {
		panic(res.Error)
	}
}

func initRoutes(e *echo.Echo) {
	// config for middleware
	config := middleware.JWTConfig{
		Claims:     &middlewares.JwtCustomClaims{},
		SigningKey: []byte(os.Getenv("SECRET_KEY")),
	}
	restricted := middleware.JWTWithConfig(config)
	e.Static("/", "views")
	e.POST("/api/token", controllers.GetToken)
	e.GET("/api/tasks", controllers.GetTasks, restricted)
	e.GET("/api/tasks/:id", controllers.GetTask, restricted)
	e.POST("/api/tasks", controllers.CreateTask, restricted)
	e.POST("/api/tasks/:id", controllers.UpdateTask, restricted)
	e.DELETE("/api/tasks/:id", controllers.DeleteTask, restricted)
	// return index while other routes not matching
	e.Static("/dist", "./views/dist")
	e.File("*", "./views")
}
