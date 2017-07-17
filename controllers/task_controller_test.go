package controllers

import (
	"testing"
	"todo-app/models/connection"
)

func TestCretaeTask(t *testing.T) {
	connection.InitConnection()
	defer connection.DB.Close()

}

func TestGetTasks(t *testing.T) {
	connection.InitConnection()
	defer connection.DB.Close()

}
