package tasks_service

import (
	"context"
	"errors"
	"keeper-crud/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTasksRepository struct {
	mock.Mock
}

func (m *MockTasksRepository) Create(task *model.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTasksRepository) FindById(taskId int) (*model.Task, error) {
	args := m.Called(taskId)
	return args.Get(0).(*model.Task), args.Error(1)
}

func (m *MockTasksRepository) FindCreatedByUserId(userId int, offset int, limit int) ([]model.Task, error) {
	args := m.Called(userId, offset, limit)
	return args.Get(0).([]model.Task), args.Error(1)
}

func (m *MockTasksRepository) Update(task *model.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTasksRepository) Delete(task *model.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func TestCreateTask(t *testing.T) {
	mockRepo := new(MockTasksRepository)
	task := &model.Task{Title: "New Task"}

	mockRepo.On("Create", task).Return(nil)

	service := NewTasksServiceImplementation(mockRepo)
	err := service.CreateTask(context.Background(), task)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetTaskById(t *testing.T) {
	mockRepo := new(MockTasksRepository)
	task := &model.Task{Title: "New Task"}

	mockRepo.On("FindById", 1).Return(task, nil)

	service := NewTasksServiceImplementation(mockRepo)
	result, err := service.GetTaskById(context.Background(), 1)

	assert.NoError(t, err)
	assert.Equal(t, task, result)
	mockRepo.AssertExpectations(t)
}

func TestGetTasksByUserId(t *testing.T) {
	mockRepo := new(MockTasksRepository)
	tasks := []model.Task{
		{Title: "Task 1"},
		{Title: "Task 2"},
	}

	mockRepo.On("FindCreatedByUserId", 1, 0, 10).Return(tasks, nil)

	service := NewTasksServiceImplementation(mockRepo)
	result, err := service.GetTasksByUserId(context.Background(), 1, 1, 10)

	assert.NoError(t, err)
	assert.Equal(t, tasks, result)
	mockRepo.AssertExpectations(t)
}

func TestUpdateTask(t *testing.T) {
	mockRepo := new(MockTasksRepository)
	task := &model.Task{Title: "Updated Task"}

	mockRepo.On("Update", task).Return(nil)

	service := NewTasksServiceImplementation(mockRepo)
	err := service.UpdateTask(context.Background(), task)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteTask(t *testing.T) {
	mockRepo := new(MockTasksRepository)
	task := &model.Task{Title: "Task to Delete"}

	mockRepo.On("FindById", 1).Return(task, nil)
	mockRepo.On("Delete", task).Return(nil)

	service := NewTasksServiceImplementation(mockRepo)
	err := service.DeleteTask(context.Background(), 1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteTask_NotFound(t *testing.T) {
	mockRepo := new(MockTasksRepository)

	mockRepo.On("FindById", 1).Return((*model.Task)(nil), errors.New("not found"))

	service := NewTasksServiceImplementation(mockRepo)
	err := service.DeleteTask(context.Background(), 1)

	assert.Error(t, err)
	assert.Equal(t, "not found", err.Error())
	mockRepo.AssertExpectations(t)
}
