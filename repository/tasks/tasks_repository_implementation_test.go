package tasks_repository

import (
	"errors"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"keeper-crud/model"
	"testing"
)

type TasksRepositoryTestSuite struct {
	suite.Suite
	DB   *gorm.DB
	Repo TasksRepository // Use the interface type here
}

func (suite *TasksRepositoryTestSuite) SetupTest() {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	suite.Require().NoError(err)

	// Drop the existing tables to ensure a clean slate
	err = db.Migrator().DropTable(&model.Task{})
	suite.Require().NoError(err)

	// Recreate the schema
	err = db.AutoMigrate(&model.Task{})
	suite.Require().NoError(err)

	suite.DB = db
	suite.Repo = NewTasksRepositoryImplementation(db)
}

func (suite *TasksRepositoryTestSuite) TestCreateTask() {
	testTask := model.Task{
		Title:         "Sample Title",
		Description:   "Sample description",
		CreatedByUser: 1,
	}

	suite.Repo.Create(&testTask)

	var fetchedTask model.Task
	err := suite.DB.First(&fetchedTask, "title = ?", "Sample Title").Error
	suite.Require().NoError(err)

	suite.Equal(testTask.Title, fetchedTask.Title)
	suite.Equal(testTask.Description, fetchedTask.Description)
	suite.Equal(testTask.CreatedByUser, fetchedTask.CreatedByUser)
}

func (suite *TasksRepositoryTestSuite) TestFindById() {
	testTask := model.Task{
		Title:         "Sample Title",
		Description:   "Sample description",
		CreatedByUser: 1,
	}
	suite.DB.Create(&testTask)

	task, err := suite.Repo.FindById(int(testTask.ID))
	suite.Require().NoError(err)
	suite.NotNil(task)
	suite.Equal(testTask.Title, task.Title)
	suite.Equal(testTask.Description, task.Description)
	suite.Equal(testTask.CreatedByUser, task.CreatedByUser)
}

func (suite *TasksRepositoryTestSuite) TestFindCreatedByUserId() {
	testTask1 := model.Task{
		Title:         "Sample Title 1",
		Description:   "Sample description 1",
		CreatedByUser: 1,
	}
	testTask2 := model.Task{
		Title:         "Sample Title 2",
		Description:   "Sample description 2",
		CreatedByUser: 1,
	}
	suite.DB.Create(&testTask1)
	suite.DB.Create(&testTask2)

	tasks, err := suite.Repo.FindCreatedByUserId(1, 0, 0)
	suite.Require().NoError(err)
	suite.Len(tasks, 2)
}

func (suite *TasksRepositoryTestSuite) TestUpdateTask() {
	testTask := model.Task{
		Title:         "Sample Title",
		Description:   "Sample description",
		CreatedByUser: 1,
	}
	suite.DB.Create(&testTask)

	testTask.Description = "Updated description"
	err := suite.Repo.Update(&testTask)
	suite.Require().Nil(err)

	var fetchedTask model.Task
	suite.DB.First(&fetchedTask, testTask.ID)
	suite.Equal("Updated description", fetchedTask.Description)
}

func (suite *TasksRepositoryTestSuite) TestDeleteTask() {
	testTask := model.Task{
		Title:         "Sample Title",
		Description:   "Sample description",
		CreatedByUser: 1,
	}
	suite.DB.Create(&testTask)

	err := suite.Repo.Delete(&testTask)
	suite.Require().Nil(err)

	var fetchedTask model.Task
	result := suite.DB.First(&fetchedTask, testTask.ID)
	suite.Error(result.Error)
	suite.True(errors.Is(result.Error, gorm.ErrRecordNotFound))
}

func TestTasksRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(TasksRepositoryTestSuite))
}
