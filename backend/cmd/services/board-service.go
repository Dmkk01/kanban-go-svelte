package services

import (
	"time"

	"github.com/Dmkk01/kanban-go-svelte/cmd/db"
	"github.com/Dmkk01/kanban-go-svelte/cmd/models"
)

func GetBoards(userId int) ([]models.Board, error) {
	db, err := db.Connect()
	if err != nil {
		return nil, err
	}

	boards := []models.Board{}

	rows, err := db.Query("SELECT id, user_id, name, emoji, created_at, updated_at FROM board WHERE user_id = $1", userId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var board models.Board
		err := rows.Scan(&board.Id, &board.UserId, &board.Name, &board.Emoji, &board.CreatedAt, &board.UpdatedAt)
		if err != nil {
			return nil, err
		}

		boards = append(boards, board)
	}
	defer db.Close()
	return boards, nil
}

func GetBoardFromColumn(columnId int) (models.Board, error) {
	db, err := db.Connect()
	if err != nil {
		return models.Board{}, err
	}
	defer db.Close()

	var board models.Board

	err = db.QueryRow("SELECT board.id, board.user_id, board.name, board.emoji, board.created_at, board.updated_at FROM board_column INNER JOIN board ON board_column.board_id = board.id WHERE board_column.id = $1", columnId).Scan(&board.Id, &board.UserId, &board.Name, &board.Emoji, &board.CreatedAt, &board.UpdatedAt)
	if err != nil {
		return models.Board{}, err
	}

	return board, nil
}

func GetBoard(id int) (models.Board, error) {
	db, err := db.Connect()
	if err != nil {
		return models.Board{}, err
	}

	var board models.Board

	err = db.QueryRow("SELECT id, user_id, name, emoji, created_at, updated_at FROM board WHERE id = $1", id).Scan(&board.Id, &board.UserId, &board.Name, &board.Emoji, &board.CreatedAt, &board.UpdatedAt)
	if err != nil {
		return models.Board{}, err
	}

	defer db.Close()

	return board, nil
}

func CreateBoard(board models.CreateBoardRequest, userId int) int {
	db, err := db.Connect()
	if err != nil {
		return 0
	}
	defer db.Close()

	boardID := 0

	err = db.QueryRow("INSERT INTO board (user_id, name, emoji) VALUES ($1, $2, $3) RETURNING id", userId, board.Name, board.Emoji).Scan(&boardID)
	if err != nil {
		return 0
	}

	for _, column := range board.Columns {
		err := CreateColumn(boardID, column)
		if err != nil {
			return boardID
		}
	}

	return boardID
}

func UpdateBoard(id int, board models.BoardUpdate) bool {
	db, err := db.Connect()
	if err != nil {
		return false
	}

	updatedAt := time.Now()
	_, err = db.Exec("UPDATE board SET name = $1, emoji = $2, updated_at = $3 WHERE id = $4", board.Name, board.Emoji, updatedAt, id)
	defer db.Close()
	return err == nil
}

func DeleteBoard(id int) bool {
	db, err := db.Connect()
	if err != nil {
		return false
	}

	_, err = db.Exec("DELETE FROM board WHERE id = $1", id)
	defer db.Close()
	return err == nil
}
