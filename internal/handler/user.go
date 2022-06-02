package handler

import (
	"encoding/json"
	realip "github.com/Ferluci/fast-realip"
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/sirupsen/logrus"
	"net"
)

func (h *Handler) GetTime(context *routing.Context) error {
	clientIP := realip.FromRequest(context.RequestCtx)
	ua := string(context.Request.Header.UserAgent())
	city, _ := h.geo.City(net.ParseIP(clientIP))
	times, err := h.service.GetTime(ua, city)
	if err != nil {
		return err
	}
	output, err := json.Marshal(&times)
	if err != nil {
		logrus.WithError(err).Error("error marshaling all users")
	}
	_, err = context.Write(output)
	if err != nil {
		return err
	}
	return nil
}

func (h *Handler) GetDate(context *routing.Context) error {
	clientIP := realip.FromRequest(context.RequestCtx)
	ua := string(context.Request.Header.UserAgent())
	city, _ := h.geo.City(net.ParseIP(clientIP))
	date, err := h.service.GetDate(ua, city)
	if err != nil {
		return err
	}
	output, err := json.Marshal(&date)
	if err != nil {
		logrus.WithError(err).Error("error marshaling all users")
	}
	_, err = context.Write(output)
	if err != nil {
		return err
	}
	return nil
}
