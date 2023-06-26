package services

import (
	"time"

	"github.com/Dmkk01/kanban-go-svelte/cmd/db"
	"github.com/Dmkk01/kanban-go-svelte/cmd/models"
)

func GetAllTags(boardIDs []int) ([]models.BoardTag, error) {
	db, err := db.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	tags := []models.BoardTag{}

	for _, boardID := range boardIDs {
		boardTags, err := GetTagsBoardID(boardID)
		if err != nil {
			return nil, err
		}

		tags = append(tags, boardTags...)
	}

	return tags, nil
}

func GetTagsBoardID(boardId int) ([]models.BoardTag, error) {
	db, err := db.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	tags := []models.BoardTag{}

	rows, err := db.Query("SELECT id, board_id, title, color, created_at, updated_at FROM board_tags WHERE board_id = $1", boardId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var tag models.BoardTag
		err := rows.Scan(&tag.ID, &tag.BoardID, &tag.Title, &tag.Color, &tag.CreatedAt, &tag.UpdatedAt)
		if err != nil {
			return nil, err
		}

		tags = append(tags, tag)
	}

	return tags, nil
}

func GetTag(tagID int) (models.BoardTag, error) {
	db, err := db.Connect()
	if err != nil {
		return models.BoardTag{}, err
	}
	defer db.Close()

	var tag models.BoardTag
	err = db.QueryRow("SELECT id, board_id, title, color, created_at, updated_at FROM board_tags WHERE id = $1", tagID).Scan(&tag.ID, &tag.BoardID, &tag.Title, &tag.Color, &tag.CreatedAt, &tag.UpdatedAt)
	if err != nil {
		return models.BoardTag{}, err
	}

	return tag, nil
}

func GetTagsTask(taskID int) ([]models.TaskTag, error) {
	db, err := db.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	tags := []models.TaskTag{}

	rows, err := db.Query("SELECT task_id, board_tag_id, created_at, updated_at FROM tags WHERE task_id = $1", taskID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var tag models.TaskTag
		err := rows.Scan(&tag.TaskID, &tag.BoardTagID, &tag.CreatedAt, &tag.UpdatedAt)
		if err != nil {
			return nil, err
		}

		tags = append(tags, tag)
	}

	return tags, nil
}

func GetTagsTaskBoard(boardID int) ([]models.TaskTag, error) {
	db, err := db.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	tags := []models.TaskTag{}

	rows, err := db.Query("SELECT task_id, board_tag_id, created_at, updated_at FROM tags WHERE board_tag_id IN (SELECT id FROM board_tags WHERE board_id = $1)", boardID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var tag models.TaskTag
		err := rows.Scan(&tag.TaskID, &tag.BoardTagID, &tag.CreatedAt, &tag.UpdatedAt)
		if err != nil {
			return nil, err
		}

		tags = append(tags, tag)
	}

	return tags, nil
}

func AddBoardTagToTaskTag(taskID int, boardTagID int) error {
	db, err := db.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	createdAt := time.Now()
	updatedAt := time.Now()

	_, err = db.Exec("INSERT INTO tags (task_id, board_tag_id, created_at, updated_at) VALUES ($1, $2, $3, $4)", taskID, boardTagID, createdAt, updatedAt)

	return err
}

func CreateBoardTag(tag models.BoardTagUpdate) error {
	db, err := db.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	createdAt := time.Now()
	updatedAt := time.Now()

	_, err = db.Exec("INSERT INTO board_tags (board_id, title, color, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)", tag.BoardID, tag.Title, tag.Color, createdAt, updatedAt)

	return err
}

func UpdateBoardTag(tagID int, tag models.BoardTagUpdate) error {
	db, err := db.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	updatedAt := time.Now()

	_, err = db.Exec("UPDATE board_tags SET title = $1, color = $2, updated_at = $3, board_id = $4 WHERE id = $5", tag.Title, tag.Color, updatedAt, tag.BoardID, tagID)

	return err
}

func DeleteBoardTag(tagID int) error {
	db, err := db.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM board_tags WHERE id = $1", tagID)

	return err
}

func RemoveTaskTagFromTask(taskID int, boardTagID int) error {
	db, err := db.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM tags WHERE task_id = $1 AND board_tag_id = $2", taskID, boardTagID)

	return err
}
