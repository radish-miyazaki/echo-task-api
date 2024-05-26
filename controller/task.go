package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type TaskController struct{}

func (t *TaskController) GetTasks(c echo.Context) error {
	return c.JSON(http.StatusOK, "get all tasks")
}

func (t *TaskController) CreateTask(c echo.Context) error {
	var task Task
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	//created, err := usecase.CreateTask(&task)
	//if err != nil {
	//	return c.JSON(http.StatusInternalServerError, err)
	//}

	return c.JSON(http.StatusCreated, task)
}