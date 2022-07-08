package model

import "time"

type Event struct {
	ID          string    `db:"id" json:"id"`
	TypeRequest string    `db:"type_request" json:"type_request"`
	OS          string    `db:"os" json:"os"`
	Browser     string    `db:"browser" json:"browser"`
	Device      string    `db:"device" json:"device"`
	City        string    `db:"city" json:"city"`
	Country     string    `db:"country" json:"country"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}
