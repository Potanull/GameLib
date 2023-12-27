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

	apiGroup := router.Group("/api/")
	{
		apiGroup.POST("/image/:name", h.PostImage)
		v1Group := apiGroup.Group("/v1/")
		{
			gameGroup := v1Group.Group("/game/")
			{
				gameGroup.GET("/:id", h.GetGame)
				gameGroup.GET("/name/:name", h.GetGameByName)

				gameGroup.POST("/", h.PostGame)
				gameGroup.PUT("/:id", h.PutGame)
				gameGroup.DELETE("/:id", h.DeleteGame)

				gameGroup.GET("/all", h.GetAllGames)
				gameGroup.GET("/random", h.GetRandomGame)
				gameGroup.GET("/random/list", h.GetRandomListGames)
				gameGroup.PUT("/reverse/status/:id", h.ReverseDoneStatus)
			}
		}
	}

	return router
}
