package repositories

import (
	"fmt"
	"github.com/Nkez/date/internal/model"
	"github.com/jmoiron/sqlx"
)

const (
	defaultPageSize = uint64(10)
)

type EventPostgres struct {
	db *sqlx.DB
}

func NewEventPostgres(db *sqlx.DB) *EventPostgres {
	return &EventPostgres{db: db}
}

func (r *EventPostgres) Get(id string) (*model.Event, error) {
	event := &model.Event{}
	if err := r.db.Get(event, `SELECT * FROM users WHERE id=$1`, id); err != nil {
		return nil, err
	}
	return event, nil
}

func (r *EventPostgres) List(filter *model.Filter) ([]*model.Event, error) {
	events := make([]*model.Event, 0, defaultPageSize)
	query := "SELECT * FROM users"
	query, args := r.decodeFilter(query, filter)
	query = r.paginateFilter(query, filter)
	if err := r.db.Select(&events, query, args...); err != nil {
		return nil, err
	}
	return events, nil
}

func (r *EventPostgres) decodeFilter(query string, _ *model.Filter) (string, []interface{}) {
	query = fmt.Sprintf("%s WHERE 1=1", query)
	args := make([]interface{}, 0)
	query = r.db.Rebind(query)
	return query, args
}

func (r *EventPostgres) paginateFilter(query string, filter *model.Filter) string {
	size := defaultPageSize
	number := uint64(1)
	if filter.Page == nil {
		filter.Page = &size
	}
	if filter.Size == nil {
		filter.Size = &number
	}
	if *filter.Page > 1 {
		query = fmt.Sprintf("%s OFFSET %d", query, (*filter.Page-1)**filter.Size)
	}
	return fmt.Sprintf("%s LIMIT %d", query, *filter.Page)
}
