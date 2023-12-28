package handlers

import (
	"gamelib/internal/storage/db"
	"github.com/forbiddencoding/howlongtobeat"
)

type Handler struct {
	Storage *db.Storage
	HLTB    *howlongtobeat.Client
}

func NewHandler(storage *db.Storage, hltb *howlongtobeat.Client) Handler {
	return Handler{
		Storage: storage,
		HLTB:    hltb,
	}
}
