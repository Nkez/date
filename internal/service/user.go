package service

import (
	"errors"
	"fmt"
	"github.com/Nkez/date/internal/model"
	"github.com/Nkez/date/internal/repositories"
	"github.com/mssola/user_agent"
	"github.com/oschwald/geoip2-golang"
	"github.com/sirupsen/logrus"
	"time"
)

const (
	requestTypeTime  = "get time"
	requestTypeDate  = "get date"
	deviceTypeMobile = "mobile phone"
	deviceTypePc     = "PC"
)

type UserService struct {
	repository repositories.User
	geo        *geoip2.Reader
}

func NewUserService(repository repositories.User) *UserService {
	return &UserService{repository: repository}
}

func (s *UserService) GetTime(info string, city *geoip2.City) (*model.TimeInfo, error) {
	cities, err := TimeResponse()
	if err != nil {
		return nil, err
	}
	userInfo, _ := s.ParseUserInfo(info, city)
	userInfo.TypeRequest = requestTypeTime
	defer func() {
		if err := s.repository.UserInfo(userInfo); err != nil {
			logrus.Info(err.Error())
		}
	}()
	return cities, nil
}

func (s *UserService) GetDate(info string, city *geoip2.City) (*model.TimeInfo, error) {
	cities, err := TimeResponse()
	if err != nil {
		return nil, err
	}
	cities = DateResponse(cities)
	userInfo, _ := s.ParseUserInfo(info, city)
	userInfo.TypeRequest = requestTypeDate
	defer func() {
		if err := s.repository.UserInfo(userInfo); err != nil {
			logrus.Info(err.Error())
		}
	}()
	return cities, nil
}

func (s *UserService) ParseUserInfo(info string, city *geoip2.City) (*model.UserInfo, error) {
	parsed := user_agent.New(info)
	browser, _ := parsed.Browser()
	os := parsed.OS()
	device := parsed.Mobile()
	user := &model.UserInfo{
		TypeRequest: "",
		OS:          os,
		Browser:     browser,
		Device:      "",
		City:        city.Country.Names["en"],
		Country:     city.City.Names["en"],
	}
	switch device {
	case true:
		user.Device = deviceTypeMobile
	case false:
		user.Device = deviceTypePc
	}

	return user, nil
}

func TimeResponse() (*model.TimeInfo, error) {
	cities := &model.TimeInfo{
		Minsk:   "Europe/Minsk",
		Tbilisi: "Asia/Tbilisi",
		Warsaw:  "Europe/Warsaw",
		NewYork: "America/New_York",
		Vilnius: "Europe/Vilnius",
	}
	loc, err := time.LoadLocation(cities.Warsaw)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error %s", cities.Warsaw))
	}
	cities.Warsaw = time.Now().In(loc).Format("15.04.05")
	loc, err = time.LoadLocation(cities.Minsk)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error %s", cities.Minsk))
	}
	cities.Minsk = time.Now().In(loc).Format("15.04.05")
	loc, err = time.LoadLocation(cities.NewYork)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error %s", cities.NewYork))
	}
	cities.NewYork = time.Now().In(loc).Format("15.04.05")
	loc, err = time.LoadLocation(cities.Tbilisi)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error %s", cities.Tbilisi))
	}
	cities.Tbilisi = time.Now().In(loc).Format("15.04.05")
	loc, err = time.LoadLocation(cities.Vilnius)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error %s", cities.Vilnius))
	}
	cities.Vilnius = time.Now().In(loc).Format("15.04.05")
	return cities, nil
}

func DateResponse(date *model.TimeInfo) *model.TimeInfo {
	tm := time.Now().Format("02.01.2006")
	date.Minsk = date.Minsk + " " + tm
	date.Tbilisi = date.Tbilisi + " " + tm
	date.NewYork = date.NewYork + " " + tm
	date.Vilnius = date.Vilnius + " " + tm
	date.Warsaw = date.Warsaw + " " + tm
	return date
}
