package tasks_service

import (
	"context"
	"keeper-crud/model"
	repository "keeper-crud/repository/tasks"
)

type TasksServiceImplementation struct {
	tasksRepository repository.TasksRepository
}

func NewTasksServiceImplementation(repo repository.TasksRepository) TasksService {
	return &TasksServiceImplementation{tasksRepository: repo}
}

func (s *TasksServiceImplementation) CreateTask(ctx context.Context, task *model.Task) error {
	return s.tasksRepository.Create(task)
}

func (s *TasksServiceImplementation) GetTaskById(ctx context.Context, taskId int) (*model.Task, error) {
	return s.tasksRepository.FindById(taskId)
}

func (s *TasksServiceImplementation) GetTasksByUserId(ctx context.Context, userId int, page int, pageSize int) ([]model.Task, error) {
	offset := (page - 1) * pageSize
	return s.tasksRepository.FindCreatedByUserId(userId, offset, pageSize)
}

func (s *TasksServiceImplementation) UpdateTask(ctx context.Context, task *model.Task) error {
	return s.tasksRepository.Update(task)
}

func (s *TasksServiceImplementation) DeleteTask(ctx context.Context, taskId int) error {
	task, err := s.tasksRepository.FindById(taskId)
	if err != nil {
		return err
	}
	return s.tasksRepository.Delete(task)
}
