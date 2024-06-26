package handlers

import (
	"fmt"
	"net/http"

	"gamelib/internal/actions"
	"gamelib/internal/entities"
	"gamelib/pkg/web"
	"github.com/gin-gonic/gin"
)

const hltbGroup = "games"

func (h *Handler) SearchGameByName(ctx *gin.Context) {
	name, exists := ctx.GetQuery("name")
	if !exists {
		ctx.JSON(http.StatusBadRequest, web.ErrorResponse(fmt.Errorf("name is empty")))
		return
	}

	if len(name) == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"data": nil,
		})
		return
	}

	searchGame := &entities.CreateGame{
		Name: name,
	}

	games, err := actions.FindHltbGames(ctx, searchGame.Name, h.HLTB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.ErrorResponse(err))
		return
	}

	for ind := range games {
		games[ind].GameImage = *actions.ParseHltbImage(games[ind].GameImage)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": games,
	})
}
