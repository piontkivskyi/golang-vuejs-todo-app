package main

import (
	"os"
	"todo-app/controllers"
	"todo-app/middlewares"
	"todo-app/models/connection"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	connection.InitConnection()
	e := echo.New()
	initRoutes(e)
	e.Start(":1234")
	defer connection.DB.Close()
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
