package routes

import (
	"net/http"

	"github.com/Dmkk01/kanban-go-svelte/cmd/controllers"
	"github.com/Dmkk01/kanban-go-svelte/cmd/middlewares"
	"github.com/Dmkk01/kanban-go-svelte/cmd/models"
	"github.com/pangpanglabs/echoswagger/v2"
)

func columnRoutes(e echoswagger.ApiRoot) {
	columnGroup := e.Group("Columns", "/column", middlewares.AuthMiddleware).SetSecurity("Authorization")

	columnGroup.GET("", controllers.GetAllColumns).
		SetSummary("Get All Column").
		AddResponse(http.StatusOK, "Columns Found", []models.BoardColumn{}, nil).
		AddResponse(http.StatusInternalServerError, "Server Error", models.MessageResponse{}, nil)

	columnGroup.GET("/:column_id", controllers.GetColumn).
		AddParamPath("", "column_id", "Column ID").
		SetSummary("Get Single Column").
		AddResponse(http.StatusOK, "Board Found", models.BoardColumn{}, nil).
		AddResponse(http.StatusInternalServerError, "Server Error", models.MessageResponse{}, nil).
		AddResponse(http.StatusBadRequest, "Bad Request", models.MessageResponse{}, nil).
		AddResponse(http.StatusNotFound, "Column not Found", models.MessageResponse{}, nil).
		AddResponse(http.StatusForbidden, "Forbidden Access", models.MessageResponse{}, nil)

	columnGroup.PUT("/:column_id", controllers.UpdateColumn).
		AddParamPath("", "column_id", "Column ID").
		SetSummary("Update Column").
		AddParamBody(models.ColumnUpdate{}, "body", "Column Update Request", true).
		AddResponse(http.StatusOK, "Column Updated", models.StatusResponse{}, nil).
		AddResponse(http.StatusInternalServerError, "Server Error", models.MessageResponse{}, nil).
		AddResponse(http.StatusBadRequest, "Bad Request", models.MessageResponse{}, nil).
		AddResponse(http.StatusNotFound, "Column not Found", models.MessageResponse{}, nil).
		AddResponse(http.StatusForbidden, "Forbidden Access", models.MessageResponse{}, nil)

	columnGroup.DELETE("/:column_id", controllers.DeleteColumn).
		AddParamPath("", "column_id", "Column ID").
		SetSummary("Delete Column").
		AddResponse(http.StatusOK, "Column Deleted", models.StatusResponse{}, nil).
		AddResponse(http.StatusInternalServerError, "Server Error", models.MessageResponse{}, nil).
		AddResponse(http.StatusBadRequest, "Bad Request", models.MessageResponse{}, nil).
		AddResponse(http.StatusNotFound, "Column not Found", models.MessageResponse{}, nil).
		AddResponse(http.StatusForbidden, "Forbidden Access", models.MessageResponse{}, nil)

	columnGroup.GET("/:column_id/task", controllers.GetTasksByColumnID)
	columnGroup.POST("/:column_id/task", controllers.CreateTaskByColumnID)
	columnGroup.PUT("/:column_id/task/position", controllers.UpdateTaskPosition)
}
