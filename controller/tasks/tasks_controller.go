package tasks_controller

import (
	"keeper-crud/data/response"
	"keeper-crud/model"
	tasksservice "keeper-crud/services/tasks"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{Error: "Invalid request"})
		return
	}

	err := ctrl.tasksService.CreateTask(ctx.Request.Context(), &task)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Error: "Failed to create task"})
		return
	}

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   task,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

// GetTaskById Get Task by ID godoc
//
//	@Summary		Get Task by ID
//	@Description	Get task data by ID.
//	@Param			tasks	query	request.GetTaskByIdRequest	true	"get a task by id"
//	@Produce		application/json
//	@Tasks			tasks
//	@Success		200	{object}	response.Response{}
//	@Router			/tasks/:id [get]
func (ctrl *TasksController) GetTaskById(ctx *gin.Context) {
	taskId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{Error: "Invalid task ID"})
		return
	}

	task, err := ctrl.tasksService.GetTaskById(ctx.Request.Context(), taskId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Error: "Failed to retrieve task"})
		return
	}

	if task == nil {
		ctx.JSON(http.StatusNotFound, response.ErrorResponse{Error: "Task not found"})
		return
	}

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   task,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// GetTasksByUserId Get Tasks created by User ID godoc
//
//	@Summary		Get Tasks created by User providing its ID
//	@Description	Get tasks created by User data by User ID.
//	@Param			tasks	query	request.GetTasksByUserIdRequest	true	"get a user tasks data"
//	@Produce		application/json
//	@Tasks			tasks
//	@Success		200	{object}	response.Response{}
//	@Router			/user/:id/tasks [get]
func (ctrl *TasksController) GetTasksByUserId(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{Error: "Invalid user ID"})
		return
	}

	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{Error: "Invalid page number"})
		return
	}

	pageSize, err := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	if err != nil || pageSize < 1 {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{Error: "Invalid page size"})
		return
	}

	tasks, err := ctrl.tasksService.GetTasksByUserId(ctx.Request.Context(), userId, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Error: "Failed to retrieve tasks"})
		return
	}

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   tasks,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// UpdateTask Update Task godoc
//
//	@Summary		Update Task
//	@Description	Update task data by ID.
//	@Param			tasks	body	request.UpdateTaskRequest	true	"update a task"
//	@Produce		application/json
//	@Tasks			tasks
//	@Success		200	{object}	response.Response{}
//	@Router			/tasks/:id [put]
func (ctrl *TasksController) UpdateTask(ctx *gin.Context) {
	var task model.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ctrl.tasksService.UpdateTask(ctx.Request.Context(), &task)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Error: "Failed to update task"})
		return
	}

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   task,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
