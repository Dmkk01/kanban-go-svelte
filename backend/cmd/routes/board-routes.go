package routes

import (
	"net/http"

	"github.com/Dmkk01/kanban-go-svelte/cmd/controllers"
	"github.com/Dmkk01/kanban-go-svelte/cmd/middlewares"
	"github.com/Dmkk01/kanban-go-svelte/cmd/models"
	"github.com/pangpanglabs/echoswagger/v2"
)

func boardRoutes(e echoswagger.ApiRoot) {
	boardGroup := e.Group("Board", "/board", middlewares.AuthMiddleware).SetSecurity("Authorization")

	boardGroup.GET("", controllers.GetBoards).
		SetSummary("Get Boards").
		AddResponse(http.StatusOK, "Boards Found", []models.Board{}, nil).
		AddResponse(http.StatusInternalServerError, "Server Error", models.MessageResponse{}, nil)

	boardGroup.POST("", controllers.CreateBoard).
		SetSummary("Create New Board").
		AddParamBody(models.CreateBoardRequest{}, "body", "Create Board Request", true).
		AddResponse(http.StatusCreated, "Board Created", models.StatusResponse{}, nil).
		AddResponse(http.StatusInternalServerError, "Server Error", models.MessageResponse{}, nil).
		AddResponse(http.StatusBadRequest, "Bad Request", models.MessageResponse{}, nil)

	boardGroup.GET("/:id", controllers.GetBoard).
		AddParamPath("", "id", "Board ID").
		SetSummary("Get Single Board").
		AddResponse(http.StatusOK, "Board Found", []models.Board{}, nil).
		AddResponse(http.StatusInternalServerError, "Server Error", models.MessageResponse{}, nil).
		AddResponse(http.StatusBadRequest, "Bad Request", models.MessageResponse{}, nil).
		AddResponse(http.StatusNotFound, "Board not Found", models.MessageResponse{}, nil).
		AddResponse(http.StatusForbidden, "Forbidden Access", models.MessageResponse{}, nil)

	boardGroup.PUT("/:id", controllers.UpdateBoard).
		AddParamPath("", "id", "Board ID").
		SetSummary("Update Single Board").
		AddParamBody(models.BoardUpdate{}, "body", "Update Board Request", true).
		AddResponse(http.StatusOK, "Board Updated", []models.StatusResponse{}, nil).
		AddResponse(http.StatusInternalServerError, "Server Error", models.MessageResponse{}, nil).
		AddResponse(http.StatusBadRequest, "Bad Request", models.MessageResponse{}, nil).
		AddResponse(http.StatusNotFound, "Board not Found", models.MessageResponse{}, nil).
		AddResponse(http.StatusForbidden, "Forbidden Access", models.MessageResponse{}, nil)

	boardGroup.DELETE("/:id", controllers.DeleteBoard).
		AddParamPath("", "id", "Board ID").
		SetSummary("Delete Single Board").
		AddResponse(http.StatusOK, "Board Deleted", []models.StatusResponse{}, nil).
		AddResponse(http.StatusInternalServerError, "Server Error", models.MessageResponse{}, nil).
		AddResponse(http.StatusBadRequest, "Bad Request", models.MessageResponse{}, nil).
		AddResponse(http.StatusNotFound, "Board not Found", models.MessageResponse{}, nil).
		AddResponse(http.StatusForbidden, "Forbidden Access", models.MessageResponse{}, nil)
}
