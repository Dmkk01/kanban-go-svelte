package main

import (
	"net/http"
	"os"

	"github.com/Dmkk01/kanban-go-svelte/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	checkGroup := e.Group("/check")
	checkGroup.GET("", controllers.GetChecks)
	checkGroup.GET("/:id", controllers.GetCheck)
	checkGroup.POST("", controllers.CreateCheck)

	userGroup := e.Group("/user")
	userGroup.GET("", controllers.GetUsers)

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Helasdasdlo, asdasd! <3")
	})

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
