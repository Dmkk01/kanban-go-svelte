package services

import (
	"log"
	"time"

	"github.com/Dmkk01/kanban-go-svelte/cmd/db"
	"github.com/Dmkk01/kanban-go-svelte/cmd/models"
)

func GetColumns(boardID int) ([]models.BoardColumn, error) {
	db, err := db.Connect()
	if err != nil {
		return []models.BoardColumn{}, err
	}

	columns := []models.BoardColumn{}

	rows, _ := db.Query("SELECT id, board_id, name, emoji, position, created_at, updated_at FROM board_column WHERE board_id = $1", boardID)
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

	return data, nil
}

func CreateColumn(boardID int, column models.ColumnCreate) error {
	db, err := db.Connect()
	if err != nil {
		return err
	}

	createdAt := time.Now()
	updatedAt := time.Now()
	_, err = db.Exec("INSERT INTO board_column (board_id, name, emoji, position, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)", boardID, column.Name, column.Emoji, column.Position, createdAt, updatedAt)

	return err
}
