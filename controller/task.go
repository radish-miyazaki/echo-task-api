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

func (t *TaskController) List(c echo.Context) error {
	return c.JSON(http.StatusOK, "get task")
}

func (t *TaskController) Create(c echo.Context) error {
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
