package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/radish-miyazaki/echo-task-api/controller"
	"github.com/radish-miyazaki/echo-task-api/repository"
	"github.com/radish-miyazaki/echo-task-api/usecase"
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
	if _, err = db.Exec("CREATE TABLE IF NOT EXISTS tasks (id INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT);"); err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	taskRepository := repository.NewTaskRepository(db)
	taskUsecase := usecase.NewTaskUsecase(taskRepository)
	taskController := controller.NewTaskController(taskUsecase)
	e.GET("/tasks/:id", taskController.Get)
	e.POST("/tasks", taskController.Create)

	e.Start(":8080")
}
