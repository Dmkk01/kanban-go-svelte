package controllers

import (
	"database/sql"
	"net/http"

	"github.com/Dmkk01/kanban-go-svelte/cmd/models"
	"github.com/Dmkk01/kanban-go-svelte/cmd/services"
	"github.com/labstack/echo/v4"
)

func getUserById(id int) (models.User, error) {
	user, err := services.GetUser(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, echo.NewHTTPError(http.StatusNotFound, "User not found")
		}
		return models.User{}, err
	}

	return user, nil
}
