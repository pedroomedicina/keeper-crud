package tasks_service

import (
	"context"
	"keeper-crud/model"
)

type TasksService interface {
	CreateTask(ctx context.Context, task *model.Task) error
	GetTaskById(ctx context.Context, taskId int) (*model.Task, error)
	GetTasksByUserId(ctx context.Context, userId int, page int, pageSize int) ([]model.Task, error)
	UpdateTask(ctx context.Context, task *model.Task) error
	DeleteTask(ctx context.Context, taskId int) error
}
