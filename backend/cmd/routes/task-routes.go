package routes

import (
	"net/http"

	"github.com/Dmkk01/kanban-go-svelte/cmd/controllers"
	"github.com/Dmkk01/kanban-go-svelte/cmd/middlewares"
	"github.com/Dmkk01/kanban-go-svelte/cmd/models"
	"github.com/pangpanglabs/echoswagger/v2"
)

func taskRoutes(e echoswagger.ApiRoot) {
	taskGroup := e.Group("Tasks", "/task", middlewares.AuthMiddleware).SetSecurity("Authorization")

	taskGroup.GET("", controllers.GetAllTasks).
		SetSummary("Get All Tasks").
		AddResponse(http.StatusOK, "Tasks Found", []models.Task{}, nil).
		AddResponse(http.StatusInternalServerError, "Server Error", models.MessageResponse{}, nil).
		AddResponse(http.StatusNotFound, "Tasks not Found", models.MessageResponse{}, nil)

	taskGroup.PUT("/position", controllers.UpdateTaskPosition).
		SetSummary("Update Tasks Position").
		AddParamBody(models.TaskPositionUpdate{}, "body", "Task Position Body", true).
		AddResponse(http.StatusOK, "Tasks Updated", models.StatusResponse{}, nil).
		AddResponse(http.StatusInternalServerError, "Server Error", models.MessageResponse{}, nil).
		AddResponse(http.StatusBadRequest, "Bad Request", models.MessageResponse{}, nil).
		AddResponse(http.StatusNotFound, "Task not Found", models.MessageResponse{}, nil).
		AddResponse(http.StatusForbidden, "Forbidden", models.MessageResponse{}, nil)

	taskGroup.PUT("/:task_id", controllers.UpdateTaskByID).
		SetSummary("Update Task by ID").
		AddParamPath("", "task_id", "Task ID").
		AddParamBody(models.TaskUpdate{}, "body", "Task Body", true).
		AddResponse(http.StatusOK, "Task Updated", models.StatusResponse{}, nil).
		AddResponse(http.StatusInternalServerError, "Server Error", models.MessageResponse{}, nil).
		AddResponse(http.StatusBadRequest, "Bad Request", models.MessageResponse{}, nil).
		AddResponse(http.StatusNotFound, "Task not Found", models.MessageResponse{}, nil).
		AddResponse(http.StatusForbidden, "Forbidden", models.MessageResponse{}, nil)

	taskGroup.GET("/:task_id", controllers.GetTaskByID).
		SetSummary("Get Task by ID").
		AddParamPath("", "task_id", "Task ID").
		AddResponse(http.StatusOK, "Task Found", models.Task{}, nil).
		AddResponse(http.StatusInternalServerError, "Server Error", models.MessageResponse{}, nil).
		AddResponse(http.StatusBadRequest, "Bad Request", models.MessageResponse{}, nil).
		AddResponse(http.StatusNotFound, "Task not Found", models.MessageResponse{}, nil)

	taskGroup.DELETE("/:task_id", controllers.DeleteTaskByID).
		SetSummary("Delete Task by ID").
		AddParamPath("", "task_id", "Task ID").
		AddResponse(http.StatusOK, "Task Deleted", models.StatusResponse{}, nil).
		AddResponse(http.StatusInternalServerError, "Server Error", models.MessageResponse{}, nil).
		AddResponse(http.StatusBadRequest, "Bad Request", models.MessageResponse{}, nil).
		AddResponse(http.StatusNotFound, "Task not Found", models.MessageResponse{}, nil).
		AddResponse(http.StatusForbidden, "Forbidden", models.MessageResponse{}, nil)

}
