package routes

import (
	"net/http"

	"github.com/Dmkk01/kanban-go-svelte/cmd/controllers"
	"github.com/Dmkk01/kanban-go-svelte/cmd/middlewares"
	"github.com/Dmkk01/kanban-go-svelte/cmd/models"
	"github.com/pangpanglabs/echoswagger/v2"
)

func tagsRoutes(e echoswagger.ApiRoot) {
	tagGroup := e.Group("Tags", "/tag", middlewares.AuthMiddleware).SetSecurity("Authorization")

	tagGroup.GET("/:tag_id", controllers.GetTag).
		AddParamPath("", "tag_id", "Tag ID").
		SetSummary("Get Board Tag").
		AddResponse(http.StatusOK, "Tag Found", models.BoardTag{}, nil).
		AddResponse(http.StatusInternalServerError, "Server Error", models.MessageResponse{}, nil).
		AddResponse(http.StatusBadRequest, "Bad Request", models.MessageResponse{}, nil).
		AddResponse(http.StatusNotFound, "Not Found", models.MessageResponse{}, nil).
		AddResponse(http.StatusForbidden, "Forbidden", models.MessageResponse{}, nil)

	tagGroup.DELETE("/:tag_id", controllers.DeleteTag).
		AddParamPath("", "tag_id", "Tag ID").
		SetSummary("Delete Board Tag").
		AddResponse(http.StatusOK, "Tag Deleted", models.StatusResponse{}, nil).
		AddResponse(http.StatusInternalServerError, "Server Error", models.MessageResponse{}, nil).
		AddResponse(http.StatusBadRequest, "Bad Request", models.MessageResponse{}, nil).
		AddResponse(http.StatusNotFound, "Not Found", models.MessageResponse{}, nil).
		AddResponse(http.StatusForbidden, "Forbidden", models.MessageResponse{}, nil)

	tagGroup.PUT("/:tag_id", controllers.UpdateTag).
		AddParamPath("", "tag_id", "Board ID").
		AddParamBody(models.BoardTagUpdate{}, "body", "Tag Update Request", true).
		SetSummary("Update Board Tag").
		AddResponse(http.StatusOK, "Tags Found", models.StatusResponse{}, nil).
		AddResponse(http.StatusInternalServerError, "Server Error", models.MessageResponse{}, nil).
		AddResponse(http.StatusBadRequest, "Bad Request", models.MessageResponse{}, nil).
		AddResponse(http.StatusNotFound, "Not Found", models.MessageResponse{}, nil).
		AddResponse(http.StatusForbidden, "Forbidden", models.MessageResponse{}, nil)

}
