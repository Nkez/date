package handler

import (
	"github.com/Nkez/date/internal/service"
	"github.com/oschwald/geoip2-golang"
	routing "github.com/qiangxue/fasthttp-routing"
)

type Handler struct {
	service *service.Service
	geo     *geoip2.Reader
}

func NewHandler(service *service.Service, geo *geoip2.Reader) *Handler {
	return &Handler{service: service, geo: geo}
}

func (h *Handler) InitRouter() *routing.Router {
	r := routing.New()
	r.Get("/date", h.GetDate)
	r.Get("/time", h.GetTime)
	return r
}
