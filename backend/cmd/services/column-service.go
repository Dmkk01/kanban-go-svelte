package services

import (
	"log"
	"time"

	"github.com/Dmkk01/kanban-go-svelte/cmd/db"
	"github.com/Dmkk01/kanban-go-svelte/cmd/models"
)

func GetColumnsFromBoard(boardID int) ([]models.BoardColumn, error) {
	db, err := db.Connect()
	if err != nil {
		return []models.BoardColumn{}, err
	}

	columns := []models.BoardColumn{}

	rows, _ := db.Query("SELECT id, board_id, name, emoji, position, created_at, updated_at FROM board_column WHERE board_id = $1 order by position asc", boardID)
	for rows.Next() {
		var column models.BoardColumn
		err := rows.Scan(&column.ID, &column.BoardID, &column.Name, &column.Emoji, &column.Position, &column.CreatedAt, &column.UpdatedAt)
		if err != nil {
			log.Println(err)
			return columns, err
		}

		columns = append(columns, column)
	}
	defer db.Close()
	return columns, nil
}

func GetColumn(columnID int) (models.BoardColumn, error) {
	db, err := db.Connect()
	if err != nil {
		return models.BoardColumn{}, err
	}

	var data models.BoardColumn
	err = db.QueryRow("SELECT id, board_id, name, emoji, position, created_at, updated_at FROM board_column WHERE id = $1", columnID).Scan(&data.ID, &data.BoardID, &data.Name, &data.Emoji, &data.Position, &data.CreatedAt, &data.UpdatedAt)

	if err != nil {
		return models.BoardColumn{}, err
	}
	defer db.Close()
	return data, nil
}

func GetAllColumns(userID int) ([]models.BoardColumn, error) {
	db, err := db.Connect()
	if err != nil {
		return []models.BoardColumn{}, err
	}
	defer db.Close()

	columns := []models.BoardColumn{}

	rows, _ := db.Query("SELECT id, board_id, name, emoji, position, created_at, updated_at FROM board_column WHERE board_id IN (SELECT id FROM board WHERE user_id = $1)", userID)
	for rows.Next() {
		var column models.BoardColumn
		err := rows.Scan(&column.ID, &column.BoardID, &column.Name, &column.Emoji, &column.Position, &column.CreatedAt, &column.UpdatedAt)
		if err != nil {
			log.Println(err)
			return columns, err
		}

		columns = append(columns, column)
	}

	return columns, nil
}

func CreateColumn(boardID int, column models.ColumnCreate) error {
	db, err := db.Connect()
	if err != nil {
		return err
	}

	createdAt := time.Now()
	updatedAt := time.Now()
	_, err = db.Exec("INSERT INTO board_column (board_id, name, emoji, position, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)", boardID, column.Name, column.Emoji, column.Position, createdAt, updatedAt)
	defer db.Close()
	return err
}

func UpdateColumn(columnID int, column models.ColumnUpdate) error {
	db, err := db.Connect()
	if err != nil {
		return err
	}
	defer db.Close()
	updatedAt := time.Now()
	_, err = db.Exec("UPDATE board_column SET name = $1, emoji = $2, updated_at = $3 WHERE id = $4", column.Name, column.Emoji, updatedAt, columnID)
	return err
}

func UpdateColumnPositions(data []models.ColumnUpdatePosition) error {
	db, err := db.Connect()
	if err != nil {
		return err
	}
	defer db.Close()
	updatedAt := time.Now()

	for _, item := range data {
		_, err = db.Exec("UPDATE board_column SET position = $1, updated_at = $2 WHERE id = $3", item.Position, updatedAt, item.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

func DeleteColumn(columnID int) error {
	db, err := db.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM board_column WHERE id = $1", columnID)

	return err
}
