package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/Dmkk01/kanban-go-svelte/cmd/models"
	"github.com/Dmkk01/kanban-go-svelte/cmd/services"
	"github.com/labstack/echo/v4"
)

func getBoardTagByContext(c echo.Context) (models.BoardTag, error) {
	claims := c.Get("claims").(*models.Claims)
	tagID, err := strconv.Atoi(c.Param("tag_id"))
	if err != nil {
		return models.BoardTag{}, echo.NewHTTPError(http.StatusBadRequest, "Invalid tag id type")
	}

	tag, err := services.GetTag(tagID)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.BoardTag{}, echo.NewHTTPError(http.StatusNotFound, "No tag found")
		}

		return models.BoardTag{}, echo.NewHTTPError(http.StatusInternalServerError, "Failed to get tag")
	}

	_, err = checkUserBoardByID(tag.BoardID, claims.Id)
	if err != nil {
		return models.BoardTag{}, err
	}

	return tag, nil
}

func GetTagsBoardID(c echo.Context) error {
	board, err := getCheckUserBoardById(c)
	if err != nil {
		return err
	}

	tags, err := services.GetBoardTagsBoardID(board.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(http.StatusNotFound, "No tags found")
		}

		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get tags")
	}

	return c.JSON(http.StatusOK, tags)
}

func GetTag(c echo.Context) error {
	tag, err := getBoardTagByContext(c)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, tag)
}

func CreateTag(c echo.Context) error {
	_, err := getCheckUserBoardById(c)
	if err != nil {
		return err
	}

	var tag models.BoardTagUpdate
	if err := c.Bind(&tag); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid tag body data")
	}

	err = services.CreateBoardTag(tag)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create tag")
	}

	return c.JSON(http.StatusCreated, models.StatusResponse{Status: "OK"})
}

func AddTagToTask(c echo.Context) error {
	tag, err := getBoardTagByContext(c)
	if err != nil {
		return err
	}

	task, err := getCheckTaskByIdContext(c)
	if err != nil {
		return err
	}

	err = services.AddBoardTagToTaskTag(task.ID, tag.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to add tag to task")
	}

	return c.JSON(http.StatusOK, models.StatusResponse{Status: "OK"})
}

func UpdateTag(c echo.Context) error {
	tag, err := getBoardTagByContext(c)
	if err != nil {
		return err
	}

	var tagUpdate models.BoardTagUpdate
	if err := c.Bind(&tagUpdate); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid tag body data")
	}

	err = services.UpdateBoardTag(tag.ID, tagUpdate)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update tag")
	}

	return c.JSON(http.StatusOK, models.StatusResponse{Status: "OK"})
}

func DeleteTag(c echo.Context) error {
	tag, err := getBoardTagByContext(c)
	if err != nil {
		return err
	}

	err = services.DeleteBoardTag(tag.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete tag")
	}

	return c.JSON(http.StatusOK, models.StatusResponse{Status: "OK"})
}

func GetTagTasks(c echo.Context) error {
	task, err := getCheckTaskByIdContext(c)
	if err != nil {
		return err
	}

	tags, err := services.GetTagsTask(task.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(http.StatusNotFound, "No tags found")
		}

		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get tags")
	}

	return c.JSON(http.StatusOK, tags)
}

func RemoveTaskTagFromTask(c echo.Context) error {
	task, err := getCheckTaskByIdContext(c)
	if err != nil {
		return err
	}

	tagID, err := strconv.Atoi(c.Param("tag_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid tag id type")
	}

	err = services.RemoveTaskTagFromTask(task.ID, tagID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to remove tag from task")
	}
	return c.JSON(http.StatusOK, models.StatusResponse{Status: "OK"})
}
