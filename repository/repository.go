package repository

import (
	"github.com/Ruslan706/go-news/commit/86c72d46a63efd17c0da10c052fdbd455e194421/internal/domain/models"
)

type Storage struct {
	db map[string]models.User
}

func New() *Storage {
	db := make(map[string]models.User)
	return &Storage{
		db: db,
	}
}
