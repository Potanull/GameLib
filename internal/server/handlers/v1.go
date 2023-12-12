package handlers

import (
	"net/http"
	"strconv"

	"gamelib/internal/actions"
	"gamelib/internal/entities"
	"gamelib/pkg/web"
	"github.com/gin-gonic/gin"
)

func (h *Handler) MainPage(ctx *gin.Context) {
	base, err := actions.GetAllGames(ctx, h.Storage)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.ErrorResponse(err))
		return
	}

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"gameList":   base,
		"randomGame": "",
	})
}

func (h *Handler) GetGame(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.ErrorResponse(err))
		return
	}

	game, err := actions.GetGame(ctx, id, h.Storage)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": game,
	})
}

func (h *Handler) GetGameByName(ctx *gin.Context) {
	name := ctx.Param("name")
	game, err := actions.GetGameByName(ctx, name, h.Storage)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": game,
	})
}

func (h *Handler) GetAllGames(ctx *gin.Context) {
	games, err := actions.GetAllGames(ctx, h.Storage)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": games,
	})
}

func (h *Handler) GetRandomGame(ctx *gin.Context) {
	done := ctx.GetBool("done")
	randomGame, err := actions.GetRandomGame(ctx, done, h.Storage)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": randomGame,
	})
}

func (h *Handler) GetRandomListGames(ctx *gin.Context) {
	done := ctx.GetBool("done")
	randimGames, err := actions.GetRandomListGames(ctx, done, h.Storage)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": randimGames,
	})
}

func (h *Handler) PostGame(ctx *gin.Context) {
	var game *entities.CreateGame
	if err := ctx.ShouldBindJSON(&game); err != nil {
		ctx.JSON(http.StatusBadRequest, web.ErrorResponse(err))
		return
	}

	check, _, err := actions.CheckGameByName(ctx, game.Name, h.Storage)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.ErrorResponse(err))
		return
	}

	if check {
		ctx.JSON(http.StatusOK, web.ExistResponse())
		return
	}

	result, err := actions.CreateGame(ctx, game, h.Storage)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": result,
	})
}

func (h *Handler) PutGame(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.ErrorResponse(err))
		return
	}

	var game *entities.UpdateGame
	if err := ctx.ShouldBindJSON(&game); err != nil {
		ctx.JSON(http.StatusBadRequest, web.ErrorResponse(err))
		return
	}

	result, err := actions.PutGame(ctx, id, game, h.Storage)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": result,
	})
}

func (h *Handler) ReverseDoneStatus(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.ErrorResponse(err))
		return
	}

	err = actions.ReverseDoneStatus(ctx, id, h.Storage)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

//func (h *Handler) deleteGameRequest(c *gin.Context) {
//	game := c.Request.FormValue("deleteGame")
//
//	if err := h.Game.DeleteGameRequest(game, c); err != nil {
//		log.Fatalf("request invalid: %s", err)
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"deleteGame": game,
//	})
//}
//
//func (h *Handler) updateGameDoneRequest(c *gin.Context) {
//	game := c.Request.FormValue("updateDoneGame")
//	if err := h.Game.UpdateGameDoneRequest(game, c); err != nil {
//		log.Fatalf("request invalid: %s", err)
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"updateDoneGame": game,
//	})
//}
