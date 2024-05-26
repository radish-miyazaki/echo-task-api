package usecase

import (
	"fmt"

	"github.com/radish-miyazaki/echo-task-api/repository"
)

type TaskUsecase interface {
	Create(title string) error
	Get(id int) (*repository.Task, error)
	Update(id int, title string) error
	Delete(id int) error
}

type taskUseCase struct {
	r repository.TaskRepository
}

func NewTaskUsecase(r repository.TaskRepository) TaskUsecase {
	return &taskUseCase{r}
}

func (u *taskUseCase) Create(title string) error {
	task := repository.Task{Title: title}
	id, err := u.r.Create(&task)
	fmt.Println(id)

	return err
}

func (u *taskUseCase) Get(id int) (*repository.Task, error) {
	return u.r.Read(id)
}

func (u *taskUseCase) Update(id int, title string) error {
	task := repository.Task{ID: id, Title: title}
	err := u.r.Update(&task)
	return err
}

func (u *taskUseCase) Delete(id int) error {
	return u.r.Delete(id)
}
