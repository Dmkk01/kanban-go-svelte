package routes

import (
	"net/http"

	"github.com/Dmkk01/kanban-go-svelte/cmd/controllers"
	"github.com/labstack/echo/v4"
	"github.com/pangpanglabs/echoswagger/v2"
)

func InitRoutes(e echoswagger.ApiRoot) {
	checkGroup := e.Group("Check", "/check")
	checkGroup.GET("", controllers.GetChecks)
	checkGroup.GET("/:id", controllers.GetCheck)
	checkGroup.POST("", controllers.CreateCheck)

	authRoutes(e)
	userRoutes(e)
	boardRoutes(e)
	columnRoutes(e)
	taskRoutes(e)

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Helasdasdlo, asdasd! <3")
	})

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})
}
