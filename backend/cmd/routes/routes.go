package routes

import (
	"net/http"

	"github.com/Dmkk01/kanban-go-svelte/cmd/middlewares"
	"github.com/Dmkk01/kanban-go-svelte/cmd/routes/controllers"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	checkGroup := e.Group("/check")
	checkGroup.GET("", controllers.GetChecks)
	checkGroup.GET("/:id", controllers.GetCheck)
	checkGroup.POST("", controllers.CreateCheck)

	authGroup := e.Group("/auth")
	authGroup.POST("/login", controllers.Login)
	authGroup.POST("/register", controllers.Register)
	authGroup.POST("/refresh-token", controllers.RefreshToken)

	userGroup := e.Group("/user", middlewares.AuthMiddleware)
	userGroup.GET("", controllers.GetUsers)
	userGroup.GET("/:id", controllers.GetUser)
	userGroup.PUT("/me/activate", controllers.ActivateUser)
	userGroup.PUT("/me/deactivate", controllers.DeactivateUser)
	userGroup.PUT("/me/password", controllers.UpdateUserPassword)
	userGroup.PUT("/me/email", controllers.UpdateUserEmail)
	userGroup.PUT("/me/username", controllers.UpdateUserUsername)

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Helasdasdlo, asdasd! <3")
	})

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})
}
