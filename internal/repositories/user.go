package repositories

import (
	"fmt"
	"github.com/Nkez/date/internal/model"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) UserInfo(user *model.UserInfo) error {
	query := fmt.Sprintf(`INSERT INTO users (
	type_request,
	browser,
	os,
	device,
	city,
	country
	)
	VALUES ($1,$2,$3,$4,$5,$6)`)
	result := r.db.QueryRow(query, user.TypeRequest, user.Browser, user.OS, user.Device, user.City, user.Country)
	if err := result.Err(); err != nil {
		pgError, _ := err.(*pq.Error)
		return pgError
	}
	return nil
}
