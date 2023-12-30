package handlers

import (
	"gamelib/internal/actions"
	"gamelib/internal/entities"
	"gamelib/pkg/web"
	"github.com/forbiddencoding/howlongtobeat"
	"github.com/gin-gonic/gin"
	"log"
	"math"
	"net/http"
	"strconv"
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
	randomGames, err := actions.GetRandomListGames(ctx, done, h.Storage)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": randomGames,
	})
}

func (h *Handler) PostGame(ctx *gin.Context) {
	var game *entities.CreateGame
	if err := ctx.ShouldBindJSON(&game); err != nil {
		ctx.JSON(http.StatusBadRequest, web.ErrorResponse(err))
		return
	}

	var err error
	var resultHLTB *howlongtobeat.GameDetailSimple

	if game.HowLongToBeatID != 0 {
		resultHLTB, err = actions.GetHltbGame(ctx, game.HowLongToBeatID, h.HLTB)
		if err != nil {
			log.Println(err)
		}
	}

	if game.FindGrid && resultHLTB != nil {
		game.Image = actions.ParseHltbImage(resultHLTB.GameImage)
	} else if game.FindGrid {
		searchGame, err := actions.FindHltbGame(ctx, game, h.HLTB)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.ErrorResponse(err))
			return
		}

		if searchGame != nil {
			game.Image = actions.ParseHltbImage(searchGame.GameImage)
		}
	}

	if resultHLTB != nil {
		game.HowLongToBeatID = resultHLTB.GameID
		game.HowLongToBeatMainTime = int(math.Round(resultHLTB.CompMain))
		game.HowLongToBeatFullTime = int(math.Round(resultHLTB.CompPlus))
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

	if game.HowLongToBeatID != 0 {
		gameHLTB, err := actions.GetHltbGame(ctx, game.HowLongToBeatID, h.HLTB)
		if err != nil {
			return
		}

		game.HowLongToBeatMainTime = int(math.Round(gameHLTB.CompMain))
		game.HowLongToBeatFullTime = int(math.Round(gameHLTB.CompPlus))

		if *game.FindGrid {
			game.ImageURL = actions.ParseHltbImage(gameHLTB.GameImage)
		}
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

func (h *Handler) DeleteGame(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.ErrorResponse(err))
		return
	}

	result, err := actions.DeleteGame(ctx, id, h.Storage)
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
