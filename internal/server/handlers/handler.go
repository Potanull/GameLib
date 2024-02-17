package handlers

import (
	"gamelib/internal/storage/minio"
	"gamelib/internal/storage/postgres"
	"github.com/forbiddencoding/howlongtobeat"
)

type Handler struct {
	Storage *postgres.Storage
	Minio   *minio.Client
	HLTB    *howlongtobeat.Client
}

func NewHandler(storage *postgres.Storage, minio *minio.Client, hltb *howlongtobeat.Client) Handler {
	return Handler{
		Storage: storage,
		Minio:   minio,
		HLTB:    hltb,
	}
}
