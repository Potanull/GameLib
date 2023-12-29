package actions

import (
	"gamelib/internal/entities"
	"github.com/forbiddencoding/howlongtobeat"
	"github.com/gin-gonic/gin"
)

const imagePath = "https://howlongtobeat.com/games/"

func GetImage(path string) *string {
	result := imagePath + path
	return &result
}

func FindGame(ctx *gin.Context, game *entities.CreateGame, hltb *howlongtobeat.Client) (*howlongtobeat.SearchGameData, error) {
	searchResults, err := hltb.Search(ctx, game.Name, howlongtobeat.SearchModifierNone, nil)
	if err != nil {
		return nil, err
	}

	if len(searchResults.Data) > 0 {
		return &searchResults.Data[0], nil
	}
	return nil, err
}
