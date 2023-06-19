package routes

import (
	"github.com/Dmkk01/kanban-go-svelte/cmd/controllers"
	"github.com/Dmkk01/kanban-go-svelte/cmd/middlewares"
	"github.com/pangpanglabs/echoswagger/v2"
)

func boardRoutes(e echoswagger.ApiRoot) {
	boardGroup := e.Group("Board", "/board", middlewares.AuthMiddleware).SetSecurity("Authorization")

	boardGroup.GET("", controllers.GetBoards)
	boardGroup.POST("", controllers.CreateBoard)
	boardGroup.GET("/:id", controllers.GetBoard)
	boardGroup.PUT("/:id", controllers.UpdateBoard)
	boardGroup.DELETE("/:id", controllers.DeleteBoard)
}
