package usecase

import (
	"github.com/radish-miyazaki/echo-task-api/model"
	"github.com/radish-miyazaki/echo-task-api/repository"
)

type TaskUsecase interface {
	Create(title string) (int, error)
	Get(id int) (*model.Task, error)
	Update(id int, title string) error
	Delete(id int) error
}

type taskUseCase struct {
	r repository.TaskRepository
}

func NewTaskUsecase(r repository.TaskRepository) TaskUsecase {
	return &taskUseCase{r}
}

func (u *taskUseCase) Create(title string) (int, error) {
	task := model.Task{Title: title}
	if err := task.Validate(); err != nil {
		return 0, err
	}

	id, err := u.r.Create(&task)

	return id, err
}

func (u *taskUseCase) Get(id int) (*model.Task, error) {
	return u.r.Read(id)
}

func (u *taskUseCase) Update(id int, title string) error {
	task := model.Task{ID: id, Title: title}
	if err := task.Validate(); err != nil {
		return err
	}

	err := u.r.Update(&task)
	return err
}

func (u *taskUseCase) Delete(id int) error {
	return u.r.Delete(id)
}
