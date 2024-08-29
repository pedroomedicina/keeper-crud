package tasks_repository

import (
	"gorm.io/gorm"
	"keeper-crud/model"
)

type TasksRepositoryImplementation struct {
	Db *gorm.DB
}

func NewTasksRepositoryImplementation(Db *gorm.DB) TasksRepository {
	return &TasksRepositoryImplementation{Db: Db}
}

func (t *TasksRepositoryImplementation) Create(task *model.Task) error {
	return t.Db.Create(&task).Error
}

func (t *TasksRepositoryImplementation) FindById(id int) (*model.Task, error) {
	var task model.Task
	result := t.Db.First(&task, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &task, nil
}

func (t *TasksRepositoryImplementation) FindCreatedByUserId(userId int, offset int, limit int) ([]model.Task, error) {
	var tasks []model.Task
	err := t.Db.Where("user_id = ?", userId).Offset(offset).Limit(limit).Find(&tasks).Error
	return tasks, err
}

func (t *TasksRepositoryImplementation) Update(task *model.Task) error {
	return t.Db.Save(task).Error
}

func (t *TasksRepositoryImplementation) Delete(task *model.Task) error {
	return t.Db.Delete(task).Error
}
