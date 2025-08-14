package router

import (
	tasks_controller "keeper-crud/controller/tasks"

	"github.com/gin-gonic/gin"
)

func NewTasksRouter(baseRouter *gin.RouterGroup) *gin.RouterGroup {
	return baseRouter.Group("/tasks")
}

func SetupTasksRouter(baseRouter *gin.RouterGroup, tasksController *tasks_controller.TasksController) {
	tasksRouter := NewTasksRouter(baseRouter)
	tasksRouter.POST("/", tasksController.CreateTask)
	tasksRouter.GET("/:id", tasksController.GetTaskById)
	tasksRouter.GET("/", tasksController.GetTasksByUserId)
	tasksRouter.PUT("/:id", tasksController.UpdateTask)
	tasksRouter.DELETE("/:id", tasksController.DeleteTask)
}
