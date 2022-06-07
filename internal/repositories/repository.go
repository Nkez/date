package repositories

import (
	"github.com/Nkez/date/internal/model"
	"github.com/jmoiron/sqlx"
)

type User interface {
	UserInfo(config *model.UserInfo) error
}

type Event interface {
	Get(id string) (*model.Event, error)
	List(filter *model.Filter) ([]*model.Event, error)
}

type Repository struct {
	User
	Event
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User:  NewUserPostgres(db),
		Event: NewEventPostgres(db),
	}
}
