package controllers

import (
	"net/http"
	"strconv"
	"todo-app/models"
	"todo-app/models/connection"

	"github.com/labstack/echo"
)

// GetTasks returns list of all tasks
func GetTasks(c echo.Context) error {
	tasks := []models.Task{}

	var (
		page, _    = strconv.Atoi(c.QueryParam("page"))
		perPage, _ = strconv.Atoi(c.QueryParam("per-page"))
		query      = connection.DB
	)

	if perPage != 0 {
		if page < 1 {
			page = 1
		}
		query.Limit(perPage).Offset((page - 1) * perPage)
	}

	if res := query.Find(&tasks); res.Error != nil {
		return c.JSON(http.StatusInternalServerError, res.Error)
	}

	return c.JSON(http.StatusOK, tasks)
}

// CreateTask get task info and save new Task
func CreateTask(c echo.Context) error {
	task := models.Task{}

	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if res := connection.DB.Create(&task); res.Error != nil {
		return c.JSON(http.StatusInternalServerError, res.Error)
	}

	return c.JSON(http.StatusCreated, task)
}

// UpdateTask by given id
func UpdateTask(c echo.Context) error {
	// get id from url
	var taskID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	task := models.Task{}
	// try to get task by given id
	if res := connection.DB.First(&task, taskID); res.Error != nil {
		return c.JSON(http.StatusNotFound, res.Error)
	}
	// update Task model
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	// try to save
	if res := connection.DB.Save(&task); res.Error != nil {
		return c.JSON(http.StatusInternalServerError, res.Error)
	}

	return c.JSON(http.StatusOK, task)
}

// DeleteTask by given id
func DeleteTask(c echo.Context) error {
	// get id from url
	var taskID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	//
	// if res := connection.DB.First(&task, taskID); res.Error != nil {
	// 	return c.JSON(http.StatusNotFound, res.Error)
	// }

	if res := connection.DB.Delete(models.Task{}, taskID); res.Error != nil {
		return c.JSON(http.StatusBadRequest, res.Error)
	}

	return c.JSON(http.StatusOK, taskID)
}

// GetTask by given id
func GetTask(c echo.Context) error {
	var taskID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	task := new(models.Task)
	// try to get task by given id
	if res := connection.DB.First(&task, taskID); res.Error != nil {
		return c.JSON(http.StatusNotFound, res.Error)
	}

	return c.JSON(http.StatusFound, task)
}
