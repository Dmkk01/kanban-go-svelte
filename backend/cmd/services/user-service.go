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
	defer db.Close()
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
	defer db.Close()
	return user, nil
}

func GetUsers() ([]models.User, error) {
	db, err := db.Connect()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT id, email, name, password, created_at, updated_at, role, inactive_status FROM user_table")
	if err != nil {
		return nil, err
	}

	checks := []models.User{}

	for rows.Next() {
		var check models.User
		err := rows.Scan(&check.ID, &check.Email, &check.Username, &check.Password, &check.CreatedAt, &check.UpdatedAt, &check.Role, &check.InactiveStatus)
		if err != nil {
			return nil, err
		}
		checks = append(checks, check)
	}
	defer db.Close()
	return checks, nil
}

func GetUser(id int) (models.User, error) {
	db, err := db.Connect()
	if err != nil {
		return models.User{}, err
	}

	var user models.User
	err = db.QueryRow("SELECT id, email, name, password, created_at, updated_at, role, inactive_status FROM user_table WHERE id = $1", id).Scan(&user.ID, &user.Email, &user.Username, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.Role, &user.InactiveStatus)
	if err != nil {
		return models.User{}, err
	}
	defer db.Close()
	return user, nil
}

func UpdateUserPassword(id int, password string) error {
	db, err := db.Connect()
	if err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	updatedAt := time.Now()

	_, err = db.Exec("UPDATE user_table SET password = $1, updated_at = $2 WHERE id = $3", string(hashedPassword), updatedAt, id)
	defer db.Close()
	return err
}

func UpdateUserEmail(id int, email string) error {
	db, err := db.Connect()
	if err != nil {
		return err
	}

	updatedAt := time.Now()
	_, err = db.Exec("UPDATE user_table SET email = $1, updated_at = $2 WHERE id = $3", email, updatedAt, id)
	defer db.Close()
	return err
}

func UpdateUserName(id int, username string) error {
	db, err := db.Connect()
	if err != nil {
		return err
	}

	updatedAt := time.Now()
	_, err = db.Exec("UPDATE user_table SET name = $1, updated_at = $2 WHERE id = $3", username, updatedAt, id)
	defer db.Close()
	return err
}

func ChangeUserStatus(id int, status bool) error {
	db, err := db.Connect()
	if err != nil {
		return err
	}

	updatedAt := time.Now()
	_, err = db.Exec("UPDATE user_table SET inactive_status = $1, updated_at = $2 WHERE id = $3", status, updatedAt, id)
	defer db.Close()
	return err
}

func GetUserSettings(userID int) (models.UserSettings, error) {
	db, err := db.Connect()
	if err != nil {
		return models.UserSettings{}, err
	}
	var settings models.UserSettings
	err = db.QueryRow("SELECT user_id, primary_board_id, app_name, app_emoji, date_format FROM user_settings WHERE user_id = $1", userID).Scan(&settings.UserID, &settings.PrimaryBoardID, &settings.AppName, &settings.AppEmoji, &settings.DateFormat)

	if err != nil {
		// primary_board_id could be null -> database/sql converting NULL to int is unsupported
		err = db.QueryRow("SELECT user_id, app_name, app_emoji, date_format FROM user_settings WHERE user_id = $1", userID).Scan(&settings.UserID, &settings.AppName, &settings.AppEmoji, &settings.DateFormat)
		if err != nil {
			return models.UserSettings{}, err
		}
	}
	defer db.Close()
	return settings, nil
}

func UpdateUserSettings(userID int, settings models.UpdateUserSettings) error {
	db, err := db.Connect()
	if err != nil {
		return err
	}

	_, err = db.Exec("UPDATE user_settings SET app_name = $1, app_emoji = $2, date_format = $3 WHERE user_id = $4", settings.AppName, settings.AppEmoji, settings.DateFormat, userID)
	defer db.Close()
	return err
}

func UserGettingStarted(userID int, gs models.UserGettingStarted) error {
	db, err := db.Connect()
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO user_settings (user_id, app_name, app_emoji) VALUES ($1, $2, $3)", userID, gs.AppName, gs.AppEmoji)
	defer db.Close()
	return err
}

func UpdateUserPrimaryBoard(userID, boardID int) error {
	db, err := db.Connect()
	if err != nil {
		return err
	}

	_, err = db.Exec("UPDATE user_settings SET primary_board_id = $1 WHERE user_id = $2", boardID, userID)
	defer db.Close()
	return err
}
