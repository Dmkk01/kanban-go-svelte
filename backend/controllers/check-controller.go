package controllers

import (
	"github.com/Dmkk01/kanban-go-svelte/models"
	"github.com/Dmkk01/kanban-go-svelte/services"
	"github.com/labstack/echo/v4"
)

func GetChecks(c echo.Context) error {
	checks, err := services.GetChecks()
	if err != nil {
		return err
	}

	return c.JSON(200, checks)
}

func GetCheck(c echo.Context) error {
	id := c.Param("id")
	check, err := services.GetCheck(id)
	if err != nil {
		return err
	}

	return c.JSON(200, check)
}

func CreateCheck(c echo.Context) error {
	newCheck := models.Check{
		Name:        "randome name",
		Description: "random description",
	}

	check, err := services.CreateCheck(newCheck)
	if err != nil {
		return err
	}

	return c.JSON(200, check)
}
