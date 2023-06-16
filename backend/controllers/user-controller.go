package controllers

import (
	"github.com/Dmkk01/kanban-go-svelte/services"
	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	users, err := services.GetUsers()
	if err != nil {
		return err
	}

	return c.JSON(200, users)
}
