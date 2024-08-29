package services

import (
	"github.com/go-playground/validator/v10"
	tasksrepository "keeper-crud/repository/tasks"
	usersrepository "keeper-crud/repository/users"
	tasksservice "keeper-crud/services/tasks"
	usersservice "keeper-crud/services/users"
)

type Service struct {
	TasksService tasksservice.TasksService
	UsersService usersservice.UsersService
}

func NewService(taskRepository tasksrepository.TasksRepository, usersRepository usersrepository.UsersRepository,
	validate *validator.Validate) *Service {
	return &Service{
		TasksService: tasksservice.NewTasksServiceImplementation(taskRepository),
		UsersService: usersservice.NewUsersServiceImplementation(usersRepository, validate),
	}
}
