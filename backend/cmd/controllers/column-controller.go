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

	columns, _ := services.GetColumnsFromBoard(board.Id)

	log.Println(columns)
	return c.JSON(http.StatusOK, columns)
}

func GetColumn(c echo.Context) error {
	claims := c.Get("claims").(*models.Claims)

	columnId, err := strconv.Atoi(c.Param("column_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid id")
	}

	column, err := services.GetColumn(columnId)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Column not found")
	}

	board, _ := services.GetBoardFromColumn(columnId)
	if board.UserId != claims.Id {
		return echo.NewHTTPError(http.StatusForbidden, "Unauthorized")
	}

	return c.JSON(http.StatusOK, column)
}

func GetAllColumns(c echo.Context) error {
	claims := c.Get("claims").(*models.Claims)

	columns, err := services.GetAllColumns(claims.Id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Server Error")
	}

	return c.JSON(http.StatusOK, columns)
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

	cleanColumnPositions(board.Id)

	return c.JSON(http.StatusCreated, models.StatusResponse{Status: "OK"})
}

func UpdateColumn(c echo.Context) error {
	claims := c.Get("claims").(*models.Claims)
	columnID, err := strconv.Atoi(c.Param("column_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid id")
	}

	var data models.ColumnUpdate
	c.Bind(&data)

	_, err = services.GetColumn(columnID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Column not found")
	}

	board, _ := services.GetBoardFromColumn(columnID)
	if board.UserId != claims.Id {
		return echo.NewHTTPError(http.StatusForbidden, "Unauthorized")
	}

	err = services.UpdateColumn(columnID, data)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Server Error")
	}

	return c.JSON(http.StatusOK, models.StatusResponse{Status: "OK"})
}

func cleanColumnPositions(boardID int) {
	columns, _ := services.GetColumnsFromBoard(boardID)

	var updates []models.ColumnUpdatePosition

	for i, column := range columns {
		if column.Position != i+1 {
			var update models.ColumnUpdatePosition
			update.ID = column.ID
			update.Position = i + 1
			updates = append(updates, update)
		}
	}

	if len(updates) > 0 {
		services.UpdateColumnPositions(updates)
	}
}

func UpdateColumnPositions(c echo.Context) error {
	board, err := getCheckUserBoardById(c)
	if err != nil {
		return err
	}

	var data []models.ColumnUpdatePosition
	c.Bind(&data)

	err = services.UpdateColumnPositions(data)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Server Error")
	}

	cleanColumnPositions(board.Id)

	return c.JSON(http.StatusOK, models.StatusResponse{Status: "OK"})
}

func DeleteColumn(c echo.Context) error {
	claims := c.Get("claims").(*models.Claims)

	columnId, err := strconv.Atoi(c.Param("column_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid id")
	}

	_, err = services.GetColumn(columnId)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Column not found")
	}

	board, _ := services.GetBoardFromColumn(columnId)
	if board.UserId != claims.Id {
		return echo.NewHTTPError(http.StatusForbidden, "Unauthorized")
	}

	err = services.DeleteColumn(columnId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Server Error")
	}

	return c.JSON(http.StatusOK, models.StatusResponse{Status: "OK"})
}
