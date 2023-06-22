package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Dmkk01/kanban-go-svelte/cmd/models"
	"github.com/Dmkk01/kanban-go-svelte/cmd/services"
	"github.com/labstack/echo/v4"
)

func getCheckTaskById(taskID int, userID int) (models.Task, error) {
	task, err := getTaskByID(taskID)
	if err != nil {
		return models.Task{}, err
	}

	_, err = checkUserBoardByID(task.BoardID, userID)
	if err != nil {
		return models.Task{}, err
	}

	return task, nil
}

func getTaskByID(taskID int) (models.Task, error) {
	task, err := services.GetTaskByID(taskID)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Task{}, echo.NewHTTPError(http.StatusNotFound, "Task not Found")
		}
		return models.Task{}, echo.NewHTTPError(http.StatusInternalServerError, "Server Error")
	}
	return task, nil
}

func GetAllTasks(c echo.Context) error {
	claims := c.Get("claims").(*models.Claims)
	userID := claims.Id

	tasks, err := services.GetAllTasks(userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(http.StatusNotFound, "Tasks not Found")
		}

		return echo.NewHTTPError(http.StatusInternalServerError, "Server Error")
	}

	return c.JSON(http.StatusOK, tasks)
}

func GetTaskByID(c echo.Context) error {
	taskID, err := strconv.Atoi(c.Param("task_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid  ID")
	}

	task, err := getTaskByID(taskID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, task)
}

func GetTaskByBoardID(c echo.Context) error {
	board, err := getCheckUserBoardById(c)
	if err != nil {
		return err
	}

	tasks, err := services.GetTaskByBoardID(board.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(http.StatusNotFound, "Tasks not Found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Server Error")
	}

	return c.JSON(http.StatusOK, tasks)
}

func GetTasksByColumnID(c echo.Context) error {
	column, err := getCheckColumnById(c)
	if err != nil {
		return err
	}

	tasks, err := services.GetTaskByColumnID(column.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(http.StatusNotFound, "Tasks not Found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Server Error")
	}

	return c.JSON(http.StatusOK, tasks)
}

func CreateTaskByColumnID(c echo.Context) error {
	_, err := getCheckColumnById(c)
	if err != nil {
		return err
	}

	var data models.TaskCreate
	if err := c.Bind(&data); err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Request Body")
	}

	_, err = services.CreateTaskByColumnID(data)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Server Error")
	}

	err = cleanUpTasksPositionInColumn(data.ColumnID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, models.StatusResponse{Status: "OK"})
}

func UpdateTaskByID(c echo.Context) error {
	claims := c.Get("claims").(*models.Claims)
	taskID, err := strconv.Atoi(c.Param("task_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	_, err = getCheckTaskById(taskID, claims.Id)
	if err != nil {
		return err
	}

	var data models.TaskUpdate
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Request Body")
	}

	err = services.UpdateTaskByID(taskID, data)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Server Error")
	}

	return c.JSON(http.StatusOK, models.StatusResponse{Status: "OK"})
}

func cleanUpTasksPositionInColumn(columnID int) error {
	tasks, err := services.GetTaskByColumnID(columnID)
	if err != nil {
		return err
	}

	for i, task := range tasks {
		if task.Position != i+1 {
			var update models.TaskPositionUpdate
			update.TaskID = task.ID
			update.Position = i + 1
			update.ColumnID = columnID

			err = services.UpdateTaskPosition(update)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "Server Error Updating Column Position")
			}
		}
	}

	return nil
}

func UpdateTaskPosition(c echo.Context) error {
	claims := c.Get("claims").(*models.Claims)
	var data []models.TaskPositionUpdate
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Request Body")
	}

	columns := make(map[int]bool)

	for _, item := range data {
		task, err := getCheckTaskById(item.TaskID, claims.Id)
		if err != nil {
			return err
		}
		err = services.UpdateTaskPosition(item)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Server Error")
		}

		columns[task.ColumnID] = true
	}

	for columnID := range columns {
		err := cleanUpTasksPositionInColumn(columnID)
		if err != nil {
			return err
		}
	}

	return c.JSON(http.StatusOK, models.StatusResponse{Status: "OK"})
}

func DeleteTaskByID(c echo.Context) error {
	claims := c.Get("claims").(*models.Claims)
	taskID, err := strconv.Atoi(c.Param("task_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid  ID")
	}

	_, err = getCheckTaskById(taskID, claims.Id)
	if err != nil {
		return err
	}

	err = services.DeleteTaskByID(taskID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Server Error")
	}

	return c.JSON(http.StatusOK, models.StatusResponse{Status: "OK"})
}
