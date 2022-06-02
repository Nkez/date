package service

import (
	"github.com/Nkez/date/internal/model"
	"github.com/Nkez/date/internal/repositories"
	"github.com/oschwald/geoip2-golang"
)

type User interface {
	GetTime(info string, city *geoip2.City) (*model.TimeInfo, error)
	GetDate(info string, city *geoip2.City) (*model.TimeInfo, error)
}

type Repository struct {
	User
}

type Service struct {
	User
}

func NewService(repository *repositories.Repository) *Service {
	return &Service{
		User: NewUserService(repository.User),
	}
}
