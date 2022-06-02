package repositories

import (
	"github.com/Nkez/date/internal/model"
	"github.com/jmoiron/sqlx"
)

type User interface {
	UserInfo(config *model.UserInfo) error
}

type Repository struct {
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User: NewUserPostgres(db),
	}
}
