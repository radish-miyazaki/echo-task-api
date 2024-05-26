package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/radish-miyazaki/echo-task-api/model"
)

type taskRepositoryMock struct {
	mock.Mock
}

func (m *taskRepositoryMock) Read(id int) (*model.Task, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Task), args.Error(1)
}

func (m *taskRepositoryMock) Update(task *model.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *taskRepositoryMock) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *taskRepositoryMock) Create(task *model.Task) (int, error) {
	args := m.Called(task)
	return args.Int(0), args.Error(1)
}

func TestTaskUsecase(t *testing.T) {
	mockRepo := new(taskRepositoryMock)
	taskUsecase := NewTaskUsecase(mockRepo)

	task := model.Task{Title: "test"}
	mockRepo.On("Create", &task).Return(1, nil)

	id, err := taskUsecase.Create(task.Title)
	assert.NoError(t, err)
	assert.Equal(t, 1, id)
}
