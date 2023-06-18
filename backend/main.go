package main

import (
	"os"

	"github.com/pangpanglabs/echoswagger/v2"

	"github.com/Dmkk01/kanban-go-svelte/cmd/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	r := echoswagger.New(echo.New(), "/doc", nil)

	r.AddSecurityAPIKey("Authorization", "JWT Token", echoswagger.SecurityInHeader)

	r.Echo().Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	r.Echo().Use(middleware.Recover())
	r.Echo().Use(middleware.CORS())

	routes.InitRoutes(r)

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	r.Echo().Logger.Fatal(r.Echo().Start(":" + httpPort))
}
