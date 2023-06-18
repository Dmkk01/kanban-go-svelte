package routes

import (
	"net/http"

	"github.com/Dmkk01/kanban-go-svelte/cmd/controllers"
	"github.com/Dmkk01/kanban-go-svelte/cmd/models"
	"github.com/pangpanglabs/echoswagger/v2"
)

func authRoutes(e echoswagger.ApiRoot) {
	authGroup := e.Group("Auth", "/auth")
	authGroup.POST("/login", controllers.Login).
		AddParamBody(controllers.LoginCredentials{}, "credentials", "Credentials", true).
		SetSummary("Login").
		AddResponse(http.StatusBadRequest, "Bad Request", models.MessageResponse{}, nil).
		AddResponse(http.StatusInternalServerError, "Internal Server Error", models.MessageResponse{}, nil).
		AddResponse(http.StatusOK, "Login completed", controllers.TokenResponse{}, nil)

	authGroup.POST("/register", controllers.Register).
		AddParamBody(controllers.RegisterCredentials{}, "credentials", "Credentials", true).
		SetSummary("Register").
		AddResponse(http.StatusBadRequest, "Bad Request", models.MessageResponse{}, nil).
		AddResponse(http.StatusInternalServerError, "Internal Server Error", models.MessageResponse{}, nil).
		AddResponse(http.StatusOK, "Register completed", models.StatusResponse{}, nil)

	authGroup.POST("/refresh-token", controllers.RefreshToken).
		SetSecurity("Authorization").
		SetSummary("Refresh JWT Token").
		AddResponse(http.StatusBadRequest, "Bad Request", models.MessageResponse{}, nil).
		AddResponse(http.StatusInternalServerError, "Internal Server Error", models.MessageResponse{}, nil).
		AddResponse(http.StatusOK, "Refresh completed", controllers.TokenResponse{}, nil)
}
