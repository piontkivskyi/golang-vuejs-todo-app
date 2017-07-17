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
	api := e.Group("/api")
	api.POST("/token", controllers.GetToken)
	api.GET("/tasks", controllers.GetTasks, restricted)
	api.GET("/tasks/:id", controllers.GetTask, restricted)
	api.POST("/tasks", controllers.CreateTask, restricted)
	api.POST("/tasks/:id", controllers.UpdateTask, restricted)
	api.DELETE("/tasks/:id", controllers.DeleteTask, restricted)
	// return index while other routes not matching
	e.Static("/dist", "./views/dist")
	e.File("*", "./views")
}
