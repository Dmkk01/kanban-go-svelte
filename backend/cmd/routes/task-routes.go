package routes

import (
	"github.com/Dmkk01/kanban-go-svelte/cmd/controllers"
	"github.com/Dmkk01/kanban-go-svelte/cmd/middlewares"
	"github.com/pangpanglabs/echoswagger/v2"
)

func taskRoutes(e echoswagger.ApiRoot) {
	taskGroup := e.Group("Tasks", "/task", middlewares.AuthMiddleware).SetSecurity("Authorization")

	taskGroup.GET("", controllers.GetAllTasks)
	taskGroup.PUT("/:task_id", controllers.UpdateTaskByID)
	taskGroup.GET("/:task_id", controllers.GetTaskByID)
	taskGroup.DELETE("/:task_id", controllers.DeleteTaskByID)

	// columnGroup.DELETE("/:column_id", controllers.DeleteColumn).
	// 	AddParamPath("", "column_id", "Column ID").
	// 	SetSummary("Delete Column").
	// 	AddResponse(http.StatusOK, "Column Deleted", models.StatusResponse{}, nil).
	// 	AddResponse(http.StatusInternalServerError, "Server Error", models.MessageResponse{}, nil).
	// 	AddResponse(http.StatusBadRequest, "Bad Request", models.MessageResponse{}, nil).
	// 	AddResponse(http.StatusNotFound, "Column not Found", models.MessageResponse{}, nil).
	// 	AddResponse(http.StatusForbidden, "Forbidden Access", models.MessageResponse{}, nil)
}
