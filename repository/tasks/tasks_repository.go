package tasks_repository

import "keeper-crud/model"

type TasksRepository interface {
	Create(task *model.Task) error
	FindById(taskId int) (*model.Task, error)
	FindCreatedByUserId(userId int, offset int, limit int) ([]model.Task, error)
	Update(task *model.Task) error
	Delete(task *model.Task) error
}
