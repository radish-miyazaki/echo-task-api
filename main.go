package main

import (
	"database/sql"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/radish-miyazaki/echo-task-api/controller"
)

func initDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "test.db")
	return db, err
}

func main() {
	db, err := initDB()
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	taskController := controller.TaskController{}

	e.GET("/tasks", taskController.GetTasks)
	e.POST("/tasks", taskController.CreateTask)

	e.Start(":8080")
}
