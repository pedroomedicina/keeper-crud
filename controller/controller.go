package controller

import (
	userscontroller "keeper-crud/controller/users"
	tasksservice "keeper-crud/services/tasks"
	usersservice "keeper-crud/services/users"
)

type Controller struct {
	UsersController *userscontroller.UsersController
}

func NewController(tasksService *tasksservice.TasksService, usersService *usersservice.UsersService) *Controller {
	return &Controller{
		// TasksController: ,
		UsersController: userscontroller.NewUsersController(*usersService),
	}
}
