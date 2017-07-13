package controllers

import (
	"testing"
	"todo-app/models/connection"
)

func init() {
	connection.InitConnection()
}

func TestGetTasks(t *testing.T) {
	defer connection.DB.Close()
}
