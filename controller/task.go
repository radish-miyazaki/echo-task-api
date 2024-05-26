package controller

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/radish-miyazaki/echo-task-api/usecase"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type TaskController interface {
	Get(c echo.Context) error
	Create(c echo.Context) error
}

type taskController struct {
	u usecase.TaskUsecase
}

func NewTaskController(u usecase.TaskUsecase) TaskController {
	return &taskController{u: u}
}

func (t *taskController) Get(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{"message": fmt.Sprintf("parse error: %v", err.Error())},
		)
	}

	task, err := t.u.Get(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.JSON(http.StatusNotFound, map[string]string{"message": "task not found"})
		}
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, task)
}

func (t *taskController) Create(c echo.Context) error {
	var task Task
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	id, err := t.u.Create(task.Title)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, map[string]int{"id": id})
}
