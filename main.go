package main

import (
	"keeper-crud/config"
	"keeper-crud/controller"
	_ "keeper-crud/docs"
	"keeper-crud/helper"
	tasksrepository "keeper-crud/repository/tasks"
	usersrepository "keeper-crud/repository/users"
	"keeper-crud/router"
	"keeper-crud/services"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

//	@title			Keeper API
//	@version		1.0
//	@description	A Keeper API in Go using Gin framework

// @host		localhost:8888
// @BasePath	/api
func main() {
	log.Info().Msg("Started Server!")
	// Database
	db := config.DatabaseConnection()

	// Repository
	tasksRepository := tasksrepository.NewTasksRepositoryImplementation(db)
	usersRepository := usersrepository.NewUsersRepositoryImplementation(db)

	// Service
	validate := validator.New()
	mainService := services.NewService(tasksRepository, usersRepository, validate)

	// Controllers
	mainController := controller.NewController(&mainService.TasksService, &mainService.UsersService)

	// Router
	routes := router.NewRouter(mainController.TasksController, mainController.UsersController)

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
