package tasks_controller

import (
	"github.com/gin-gonic/gin"
	"keeper-crud/model"
	tasksservice "keeper-crud/services/tasks"
	"net/http"
)

type TasksController struct {
	tasksService tasksservice.TasksService
}

func NewTasksController(service tasksservice.TasksService) *TasksController {
	return &TasksController{tasksService: service}
}

// CreateTask Create Task		godoc
//
//	@Summary		Create Tasks
//	@Description	Save tasks data in Db.
//	@Param			tasks	body	request.CreateTaskRequest	true	"Create a task"
//	@Produce		application/json
//	@Tasks			tasks
//	@Success		200	{object}	response.Response{}
//	@Router			/tasks [post]
func (ctrl *TasksController) CreateTask(ctx *gin.Context) {
	var task model.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ctrl.tasksService.CreateTask(ctx.Request.Context(), &task)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Task created successfully", "task": task})
}
