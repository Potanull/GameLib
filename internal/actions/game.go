package actions

import (
	"fmt"
	"gamelib/internal/entities"
	"gamelib/internal/storage/db"
	"gamelib/pkg/web"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	maxLenGameName = 150

	directUp  = "../"
	PathGrids = "assets/static/grids/"
)

func GetGame(ctx *gin.Context, id int64, storage *db.Storage) (*entities.Game, error) {
	return db.GetGame(ctx, id, storage.DataBase)
}

func CreateGame(ctx *gin.Context, game *entities.CreateGame, storage *db.Storage) (*entities.Game, error) {
	if len(game.Name) == 0 {
		return nil, fmt.Errorf("name can't be empty")
	}

	if len(game.Name) > maxLenGameName {
		return nil, fmt.Errorf("name of game is too long")
	}

	check, result, err := CheckGameByName(ctx, game.Name, storage)
	if err != nil {
		return nil, err
	}

	if check {
		ctx.JSON(http.StatusOK, web.ExistResponse())
		return result, nil
	}

	if !game.FindGrid && game.Image != nil {
		newName := directUp + PathGrids + *game.Image
		game.Image = &newName
	}

	return db.CreateGame(ctx, game, storage.DataBase)
}

func PutGame(ctx *gin.Context, id int64, update *entities.UpdateGame, storage *db.Storage) (*entities.Game, error) {
	return db.PutGame(ctx, id, update, storage.DataBase)
}

func DeleteGame(ctx *gin.Context, id int64, storage *db.Storage) (*entities.Game, error) {
	game, err := GetGame(ctx, id, storage)
	if err != nil {
		return nil, err
	}

	if game.ImageURL != nil {
		err = os.Remove(strings.TrimPrefix(*game.ImageURL, directUp))
		if err != nil {
			log.Println(err)
		}
	}

	return db.DeleteGame(ctx, id, storage.DataBase)
}

func GetGameByName(ctx *gin.Context, name string, storage *db.Storage) (*entities.Game, error) {
	return db.GetGameByName(ctx, name, storage.DataBase)
}

func GetAllGames(ctx *gin.Context, storage *db.Storage) ([]*entities.Game, error) {
	return db.GetAllGames(ctx, storage.DataBase)
}

func GetRandomGame(ctx *gin.Context, done bool, storage *db.Storage) (*entities.Game, error) {
	return db.GetRandomGame(ctx, done, storage.DataBase)
}

func GetRandomListGames(ctx *gin.Context, done bool, storage *db.Storage) ([]*entities.Game, error) {
	return db.GetRandomListGames(ctx, done, storage.DataBase)
}

func GetRandomListGamesWithImage(ctx *gin.Context, done bool, storage *db.Storage) ([]*entities.Game, error) {
	return db.GetRandomListGamesWithImage(ctx, done, storage.DataBase)
}

func CheckGame(ctx *gin.Context, id int64, storage *db.Storage) (bool, *entities.Game, error) {
	return db.CheckGame(ctx, id, storage.DataBase)
}

func CheckGameByName(ctx *gin.Context, name string, storage *db.Storage) (bool, *entities.Game, error) {
	return db.CheckGameByName(ctx, name, storage.DataBase)
}

func ReverseDoneStatus(ctx *gin.Context, id int64, storage *db.Storage) error {
	return db.ReverseDoneStatus(ctx, id, storage.DataBase)
}

func ReverseFavoriteStatus(ctx *gin.Context, id int64, storage *db.Storage) error {
	return db.ReverseFavoriteStatus(ctx, id, storage.DataBase)
}
