package server

import (
	"gamelib/internal/server/handlers"
	"github.com/gin-gonic/gin"
)

func (s *Server) configureRoutes() *gin.Engine {
	router := gin.New()

	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*")

	h := handlers.NewHandler(s.Storage)

	router.GET("/", h.MainPage)
	router.POST("/api/image/:name", h.PostImage)

	accountChanges := router.Group("/api/v1/")
	{
		gameChanges := accountChanges.Group("/game/")
		{
			gameChanges.GET("/:id", h.GetGame)
			gameChanges.GET("/name/:name", h.GetGameByName)

			gameChanges.POST("/", h.PostGame)
			gameChanges.PUT("/:id", h.PutGame)
			gameChanges.DELETE("/:id", h.DeleteGame)

			gameChanges.GET("/all", h.GetAllGames)
			gameChanges.GET("/random", h.GetRandomGame)
			gameChanges.GET("/random/list", h.GetRandomListGames)
			gameChanges.PUT("/reverse/status/:id", h.ReverseDoneStatus)
		}
	}

	return router
}