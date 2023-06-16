package services

import (
	"github.com/Dmkk01/kanban-go-svelte/db"
	"github.com/Dmkk01/kanban-go-svelte/models"
)

func GetUsers() ([]models.User, error) {
	db, err := db.Connect()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT * FROM user_table")
	if err != nil {
		return nil, err
	}

	checks := []models.User{}

	for rows.Next() {
		var check models.User
		err := rows.Scan(&check.ID, &check.Username, &check.Password, &check.CreatedAt, &check.UpdatedAt)
		if err != nil {
			return nil, err
		}
		checks = append(checks, check)
	}

	return checks, nil
}
