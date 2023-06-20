package routes

import (
	"net/http"

	"github.com/Dmkk01/kanban-go-svelte/cmd/controllers"
	"github.com/Dmkk01/kanban-go-svelte/cmd/middlewares"
	"github.com/Dmkk01/kanban-go-svelte/cmd/models"
	"github.com/pangpanglabs/echoswagger/v2"
)

func userRoutes(e echoswagger.ApiRoot) {
	userGroup := e.Group("User", "/user", middlewares.AuthMiddleware).SetSecurity("Authorization")

	userGroup.GET("", controllers.GetUsers).
		SetSummary("Get Many Users").
		AddResponse(http.StatusOK, "Users Found", []models.User{}, nil).
		AddResponse(http.StatusInternalServerError, "Users Not Found", models.MessageResponse{}, nil)

	userGroup.GET("/:id", controllers.GetUser).
		AddParamPath("", "id", "User ID").
		SetSummary("Get Single User").
		AddResponse(http.StatusOK, "User Found", models.User{}, nil).
		AddResponse(http.StatusNotFound, "User Not Found", models.MessageResponse{}, nil).
		AddResponse(http.StatusBadRequest, "Invalid Id", models.MessageResponse{}, nil)

	userGroup.GET("/me", controllers.GetMe).
		SetSummary("Get Current User").
		SetDescription("Get current user, same as /:id").
		AddResponse(http.StatusOK, "User Found", models.User{}, nil).
		AddResponse(http.StatusNotFound, "User Not Found", models.MessageResponse{}, nil).
		AddResponse(http.StatusBadRequest, "Invalid Id", models.MessageResponse{}, nil)

	userGroup.PUT("/me/activate", controllers.ActivateUser).
		SetSummary("Active Current User").
		AddResponse(http.StatusOK, "User is activated", models.MessageResponse{}, nil).
		AddResponse(http.StatusInternalServerError, "Server Error", models.MessageResponse{}, nil)

	userGroup.PUT("/me/deactivate", controllers.DeactivateUser).
		SetSummary("Deactivate Current User").
		AddResponse(http.StatusOK, "User is deactivated", models.MessageResponse{}, nil).
		AddResponse(http.StatusInternalServerError, "Server Error", models.MessageResponse{}, nil)

	userGroup.PUT("/me/password", controllers.UpdateUserPassword).
		SetSummary("Change Password").
		AddParamBody(controllers.UpdatePasswordData{}, "data", "Data", true).
		AddResponse(http.StatusOK, "Password Updated", models.MessageResponse{}, nil).
		AddResponse(http.StatusInternalServerError, "Server Error", models.MessageResponse{}, nil).
		AddResponse(http.StatusBadRequest, "Invalid data", models.MessageResponse{}, nil)

	userGroup.PUT("/me/email", controllers.UpdateUserEmail).
		SetSummary("Change Email").
		AddParamBody(controllers.UpdateUserEmailData{}, "data", "Data", true).
		AddResponse(http.StatusOK, "Email Updated", models.MessageResponse{}, nil).
		AddResponse(http.StatusInternalServerError, "Server Error", models.MessageResponse{}, nil).
		AddResponse(http.StatusBadRequest, "Invalid data", models.MessageResponse{}, nil)

	userGroup.PUT("/me/username", controllers.UpdateUserUsername).
		SetSummary("Change Username").
		AddParamBody(controllers.UpdateUserUsernameData{}, "data", "Data", true).
		AddResponse(http.StatusOK, "Email Updated", models.MessageResponse{}, nil).
		AddResponse(http.StatusInternalServerError, "Server Error", models.MessageResponse{}, nil).
		AddResponse(http.StatusBadRequest, "Invalid data", models.MessageResponse{}, nil)

	userGroup.GET("/me/settings", controllers.GetUserSettings).
		SetSummary("Get User Settings")

	userGroup.POST("/getting-started", controllers.UserGettingStarted).
		SetSummary("User Getting Started")
}
