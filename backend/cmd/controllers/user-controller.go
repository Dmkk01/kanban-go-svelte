package controllers

import (
	"log"

	"github.com/Dmkk01/kanban-go-svelte/cmd/services"
	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	users, err := services.GetUsers()

	log.Println("claims", c.Get("claims"))
	if err != nil {
		return err
	}

	return c.JSON(200, users)
}
