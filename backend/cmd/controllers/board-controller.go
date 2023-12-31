package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/Dmkk01/kanban-go-svelte/cmd/models"
	"github.com/Dmkk01/kanban-go-svelte/cmd/services"
	"github.com/labstack/echo/v4"
)

func checkUserBoardByID(boardID int, userID int) (models.Board, error) {
	board, err := services.GetBoard(boardID)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Board{}, echo.NewHTTPError(http.StatusNotFound, "Board not found")
		}
		return models.Board{}, echo.NewHTTPError(http.StatusInternalServerError, "There was an error getting the board")
	}

	if board.UserId != userID {
		return models.Board{}, echo.NewHTTPError(http.StatusForbidden, "You don't have access to this board")
	}

	return board, nil
}

func getCheckUserBoardById(c echo.Context) (models.Board, error) {
	claims := c.Get("claims").(*models.Claims)
	id := claims.Id

	boardId, err := strconv.Atoi(c.Param("board_id"))
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

	boardID := services.CreateBoard(data, id)
	if boardID == 0 {
		return echo.NewHTTPError(http.StatusInternalServerError, "There was an error creating the board")
	}

	userSettings, _ := services.GetUserSettings(id)

	if userSettings.PrimaryBoardID == 0 {
		boards, _ := services.GetBoards(id)
		if len(boards) > 0 {
			services.UpdateUserPrimaryBoard(id, boards[0].Id)
		}
	}

	return c.JSON(http.StatusCreated, models.BoardCreateResponse{Id: boardID})
}

func GetBoard(c echo.Context) error {
	board, err := getCheckUserBoardById(c)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, board)
}

func GetBoardFull(c echo.Context) error {
	final := models.BoardFullResponse{}
	board, err := getCheckUserBoardById(c)
	if err != nil {
		return err
	}

	final.Board = board
	final.Columns = []models.BoardColumnFullResponse{}

	columns, _ := services.GetColumnsFromBoard(board.Id)

	for _, column := range columns {
		var columnFull models.BoardColumnFullResponse
		tasks, err := services.GetTaskByColumnID(column.ID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "There was an error getting the tasks")
		}

		columnFull.BoardColumn = column
		columnFull.Tasks = tasks

		final.Columns = append(final.Columns, columnFull)
	}

	tags, err := services.GetBoardTagsBoardID(board.Id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "There was an error getting the tags")
	}
	final.Tags = tags

	return c.JSON(http.StatusOK, final)
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
