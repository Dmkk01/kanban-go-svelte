package routes

import (
	"github.com/Dmkk01/kanban-go-svelte/cmd/controllers"
	"github.com/Dmkk01/kanban-go-svelte/cmd/middlewares"
	"github.com/pangpanglabs/echoswagger/v2"
)

func columnRoutes(e echoswagger.ApiRoot) {
	columnGroup := e.Group("Columns", "/column", middlewares.AuthMiddleware).SetSecurity("Authorization")

	columnGroup.GET("/:column_id", controllers.GetColumn)
	columnGroup.PUT("/:column_id", controllers.UpdateColumn)
	columnGroup.DELETE("/:column_id", controllers.DeleteColumn)

	columnGroup.GET("/board/:board_id", controllers.GetColumnsBoardID)
	columnGroup.POST("/board/:board_id", controllers.CreateColumn)
}
