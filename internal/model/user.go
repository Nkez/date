package model

type UserInfo struct {
	TypeRequest string `db:"type_request" json:"type_request"`
	OS          string `db:"os" json:"os"`
	Browser     string `db:"browser" json:"browser"`
	Device      string `db:"device" json:"device"`
	City        string `db:"city" json:"city"`
	Country     string `db:"country" json:"country"`
}
