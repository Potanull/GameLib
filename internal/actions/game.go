package actions

import (
	"gamelib/internal/entities"
	"gamelib/internal/storage/db"
	"github.com/gin-gonic/gin"
)

func GetGame(ctx *gin.Context, id int64, storage *db.Storage) (*entities.Game, error) {
	return db.GetGame(ctx, id, storage.DataBase)
}

func CreateGame(ctx *gin.Context, game *entities.CreateGame, storage *db.Storage) (*entities.Game, error) {
	return db.CreateGame(ctx, game, storage.DataBase)
}

func PutGame(ctx *gin.Context, id int64, update *entities.UpdateGame, storage *db.Storage) (*entities.Game, error) {
	return db.PutGame(ctx, id, update, storage.DataBase)
}

func DeleteGame(ctx *gin.Context, id int64, storage *db.Storage) (*entities.Game, error) {
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

func CheckGame(ctx *gin.Context, id int64, storage *db.Storage) (bool, *entities.Game, error) {
	return db.CheckGame(ctx, id, storage.DataBase)
}

func CheckGameByName(ctx *gin.Context, name string, storage *db.Storage) (bool, *entities.Game, error) {
	return db.CheckGameByName(ctx, name, storage.DataBase)
}

func ReverseDoneStatus(ctx *gin.Context, id int64, storage *db.Storage) error {
	return db.ReverseDoneStatus(ctx, id, storage.DataBase)
}
