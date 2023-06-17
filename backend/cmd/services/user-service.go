package services

import (
	"time"

	"github.com/Dmkk01/kanban-go-svelte/cmd/db"
	"github.com/Dmkk01/kanban-go-svelte/cmd/models"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(username, password, email string) bool {
	db, err := db.Connect()
	if err != nil {
		return false
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return false
	}
	createdAt := time.Now()
	updatedAt := time.Now()

	_, err = db.Exec("INSERT INTO user_table (email, name, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)", email, username, string(hashedPassword), createdAt, updatedAt)
	return err == nil
}

func GetUserByEmail(email string) (models.User, error) {
	db, err := db.Connect()
	if err != nil {
		return models.User{}, err
	}

	var user models.User
	err = db.QueryRow("SELECT id, email, name, password FROM user_table WHERE email = $1", email).Scan(&user.ID, &user.Email, &user.Username, &user.Password)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func GetUsers() ([]models.User, error) {
	db, err := db.Connect()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT id, email, name, password, created_at, updated_at FROM user_table")
	if err != nil {
		return nil, err
	}

	checks := []models.User{}

	for rows.Next() {
		var check models.User
		err := rows.Scan(&check.ID, &check.Email, &check.Username, &check.Password, &check.CreatedAt, &check.UpdatedAt)
		if err != nil {
			return nil, err
		}
		checks = append(checks, check)
	}

	return checks, nil
}
