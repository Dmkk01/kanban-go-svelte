package services

import (
	"github.com/Dmkk01/kanban-go-svelte/cmd/db"
	"github.com/Dmkk01/kanban-go-svelte/cmd/models"
)

func GetChecks() ([]models.Check, error) {
	db, err := db.Connect()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT * FROM check_test")
	if err != nil {
		return nil, err
	}

	checks := []models.Check{}

	for rows.Next() {
		var check models.Check
		err := rows.Scan(&check.ID, &check.Name, &check.Description)
		if err != nil {
			return nil, err
		}
		checks = append(checks, check)
	}

	return checks, nil
}

func GetCheck(id string) (models.Check, error) {
	db, err := db.Connect()
	if err != nil {
		return models.Check{}, err
	}

	var check models.Check

	err = db.QueryRow("SELECT * FROM check_test WHERE id = $1", id).Scan(&check.ID, &check.Name, &check.Description)
	if err != nil {
		return models.Check{}, err
	}

	return check, nil
}

func CreateCheck(check models.Check) (models.Check, error) {
	db, err := db.Connect()
	if err != nil {
		return models.Check{}, err
	}

	err = db.QueryRow("INSERT INTO check_test (name, description) VALUES ($1, $2) RETURNING id", check.Name, check.Description).Scan(&check.ID)
	if err != nil {
		return models.Check{}, err
	}

	return check, nil
}
