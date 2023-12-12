package handlers

import (
	"gamelib/internal/storage/db"
)

type Handler struct {
	Storage *db.Storage
}

func NewHandler(storage *db.Storage) Handler {
	return Handler{
		Storage: storage,
	}
}
