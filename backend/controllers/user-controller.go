package controllers

import (
	"log"

	"github.com/Dmkk01/kanban-go-svelte/services"
	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	users, err := services.GetUsers()
	log.Println("Found users:", users)
	if err != nil {
		return err
	}

	return c.JSON(200, users)
}
