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

	accountChanges := router.Group("/api/v1/")
	{
		gameChanges := accountChanges.Group("/game/")
		{
			gameChanges.POST("/", h.PostGame)
			gameChanges.GET("/:id", h.GetGame)
			gameChanges.PUT("/:id", h.PutGame)
			//gameChanges.DELETE("/", h.deleteGameRequest)
			//gameChanges.PATCH("/", h.updateGameDoneRequest)

			gameChanges.PUT("/reverse/status/:id", h.ReverseDoneStatus)

			gameChanges.GET("/name/:name", h.GetGameByName)
			gameChanges.GET("/all", h.GetAllGames)
			gameChanges.GET("/random", h.GetRandomGame)
			gameChanges.GET("/random/list", h.GetRandomListGames)
		}
	}

	return router
}
