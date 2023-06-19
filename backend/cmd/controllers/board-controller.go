package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/Dmkk01/kanban-go-svelte/cmd/models"
	"github.com/Dmkk01/kanban-go-svelte/cmd/services"
	"github.com/labstack/echo/v4"
)

func getCheckUserBoardById(c echo.Context) (models.Board, error) {
	claims := c.Get("claims").(*models.Claims)
	id := claims.Id

	boardId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return models.Board{}, echo.NewHTTPError(http.StatusBadRequest, "Invalid id")
	}

	board, err := services.GetBoard(boardId)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Board{}, echo.NewHTTPError(http.StatusNotFound, "Board not found")
		}
		return models.Board{}, echo.NewHTTPError(http.StatusInternalServerError, "There was an error getting the board")
	}

	if board.UserId != id {
		return models.Board{}, echo.NewHTTPError(http.StatusForbidden, "You don't have access to this board")
	}

	return board, nil
}

func GetBoards(c echo.Context) error {
	claims := c.Get("claims").(*models.Claims)
	id := claims.Id

	boards, err := services.GetBoards(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "There was an error getting the boards")
	}

	return c.JSON(http.StatusOK, boards)
}

func CreateBoard(c echo.Context) error {
	claims := c.Get("claims").(*models.Claims)
	id := claims.Id

	var data models.CreateBoardRequest
	c.Bind(&data)

	if data.Name == "" || data.Emoji == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Both fields are required")
	}

	ok := services.CreateBoard(data, id)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "There was an error creating the board")
	}

	return c.JSON(http.StatusCreated, models.StatusResponse{Status: "ok"})
}

func GetBoard(c echo.Context) error {
	board, err := getCheckUserBoardById(c)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, board)
}

func UpdateBoard(c echo.Context) error {
	board, err := getCheckUserBoardById(c)
	if err != nil {
		return err
	}

	var data models.BoardUpdate
	c.Bind(&data)

	if data.Name == "" && data.Emoji == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Nothing to update")
	}

	ok := services.UpdateBoard(board.Id, data)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "There was an error updating the board")
	}

	return c.JSON(http.StatusOK, models.StatusResponse{Status: "ok"})
}

func DeleteBoard(c echo.Context) error {
	board, err := getCheckUserBoardById(c)
	if err != nil {
		return err
	}

	ok := services.DeleteBoard(board.Id)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "There was an error deleting the board")
	}

	return c.JSON(http.StatusOK, models.StatusResponse{Status: "ok"})
}
