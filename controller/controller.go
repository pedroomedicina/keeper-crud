package controller

import (
	taskscontroller "keeper-crud/controller/tasks"
	userscontroller "keeper-crud/controller/users"
	tasksservice "keeper-crud/services/tasks"
	usersservice "keeper-crud/services/users"
)

type Controller struct {
	UsersController *userscontroller.UsersController
	TasksController *taskscontroller.TasksController
}

func NewController(tasksService *tasksservice.TasksService, usersService *usersservice.UsersService) *Controller {
	return &Controller{
		TasksController: taskscontroller.NewTasksController(*tasksService),
		UsersController: userscontroller.NewUsersController(*usersService),
	}
}
