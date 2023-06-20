package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Dmkk01/kanban-go-svelte/cmd/models"
	"github.com/Dmkk01/kanban-go-svelte/cmd/services"
	"github.com/labstack/echo/v4"
)

func GetColumnsBoardID(c echo.Context) error {
	board, err := getCheckUserBoardById(c)
	if err != nil {
		return err
	}

	columns, _ := services.GetColumns(board.Id)

	log.Println(columns)
	return c.JSON(http.StatusOK, columns)
}

func GetColumn(c echo.Context) error {
	columnId, err := strconv.Atoi(c.Param("column_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid id")
	}

	column, err := services.GetColumn(columnId)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Column not found")
	}
	return c.JSON(http.StatusOK, column)
}

func CreateColumn(c echo.Context) error {
	board, err := getCheckUserBoardById(c)
	if err != nil {
		return err
	}

	var data models.ColumnCreate
	c.Bind(&data)

	if data.Emoji == "" || data.Name == "" || data.Position <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "All fields are required")
	}

	err = services.CreateColumn(board.Id, data)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Server Error")
	}

	return c.JSON(http.StatusCreated, models.StatusResponse{Status: "OK"})
}

func UpdateColumn(c echo.Context) error {
	return nil
}

func DeleteColumn(c echo.Context) error {
	return nil
}
