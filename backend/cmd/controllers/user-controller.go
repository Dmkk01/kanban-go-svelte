package controllers

import (
	"net/http"
	"strconv"

	"github.com/Dmkk01/kanban-go-svelte/cmd/models"
	"github.com/Dmkk01/kanban-go-svelte/cmd/services"
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers(c echo.Context) error {
	users, err := services.GetUsers()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "There was an error getting the users")
	}

	return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid id")
	}

	user, err := getUserById(idInt)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

type UpdatePasswordData struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

func UpdateUserPassword(c echo.Context) error {
	claims := c.Get("claims").(*models.Claims)
	id := claims.Id

	user, err := getUserById(id)
	if err != nil {
		return err
	}

	var data UpdatePasswordData

	err = c.Bind(&data)
	if err != nil {
		return err
	}

	if data.OldPassword == "" || data.NewPassword == "" || data.OldPassword == data.NewPassword {
		return echo.NewHTTPError(http.StatusBadRequest, "Both fields are required and must be different")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.OldPassword)); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Incorrect Old password")
	}

	//TODO better password validation
	if len(data.NewPassword) < 6 {
		return echo.NewHTTPError(http.StatusBadRequest, "Password must be at least 8 characters long")
	}

	err = services.UpdateUserPassword(id, data.NewPassword)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "There was an error updating the password")
	}

	return c.JSON(http.StatusOK, models.MessageResponse{Message: "Password updated successfully"})
}

func changeUserStatus(status bool, c echo.Context) error {
	claims := c.Get("claims").(*models.Claims)
	id := claims.Id

	err := services.ChangeUserStatus(id, status)
	if err != nil {
		return err
	}

	return nil
}

func DeactivateUser(c echo.Context) error {
	err := changeUserStatus(true, c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "There was an error deactivating the user")
	}
	return c.JSON(http.StatusOK, models.MessageResponse{Message: "User deactivated successfully"})
}

func ActivateUser(c echo.Context) error {
	err := changeUserStatus(false, c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "There was an error activating the user")
	}
	return c.JSON(http.StatusOK, models.MessageResponse{Message: "User activated successfully"})
}

type UpdateUserEmailData struct {
	Email string `json:"email"`
}

func UpdateUserEmail(c echo.Context) error {
	claims := c.Get("claims").(*models.Claims)
	id := claims.Id

	var data UpdateUserEmailData

	err := c.Bind(&data)
	if err != nil {
		return err
	}

	if data.Email == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Email is required")
	}
	if !govalidator.IsEmail(data.Email) {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid email")
	}

	err = services.UpdateUserEmail(id, data.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "There was an error updating the email")
	}

	return c.JSON(http.StatusOK, models.MessageResponse{Message: "Email updated successfully"})
}

type UpdateUserUsernameData struct {
	Username string `json:"username"`
}

func UpdateUserUsername(c echo.Context) error {
	claims := c.Get("claims").(*models.Claims)
	id := claims.Id

	var data UpdateUserUsernameData

	err := c.Bind(&data)
	if err != nil {
		return err
	}

	if data.Username == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Username is required")
	}

	err = services.UpdateUserName(id, data.Username)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "There was an error updating the username")
	}

	return c.JSON(http.StatusOK, models.MessageResponse{Message: "Username updated successfully"})
}

func GetMe(c echo.Context) error {
	claims := c.Get("claims").(*models.Claims)
	id := claims.Id

	user, err := getUserById(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func GetUserSettings(c echo.Context) error {
	return nil
}

func UserGettingStarted(c echo.Context) error {
	return nil
}
