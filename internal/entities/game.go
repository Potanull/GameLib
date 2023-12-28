package entities

import (
	"time"
)

type Game struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`

	ImageURL *string `json:"image_url"`

	CreateDt time.Time `json:"create_dt"`
	UpdateDt time.Time `json:"update_dt"`
}

type CreateGame struct {
	Name  string  `json:"name"`
	Done  bool    `json:"done"`
	Image *string `json:"image"`

	FindGrid bool `json:"find_grid"`
}

type UpdateGame struct {
	Name     string  `json:"name"`
	Done     *bool   `json:"done"`
	ImageURL *string `json:"image_url"`
}
