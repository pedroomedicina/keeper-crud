package router

import (
	taskscontroller "keeper-crud/controller/tasks"
	userscontroller "keeper-crud/controller/users"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(tasksController *taskscontroller.TasksController, usersController *userscontroller.UsersController) *gin.Engine {
	router := gin.Default()

	// add swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})
	baseRouter := router.Group("/api")
	SetupTasksRouter(baseRouter, tasksController)
	SetupUsersRouter(baseRouter, usersController)

	return router
}
