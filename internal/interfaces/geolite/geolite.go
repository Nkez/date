package geolite

import (
	"github.com/oschwald/geoip2-golang"
	"log"
)

type Config struct {
	DB string
}

func NewGeoDB(cfg Config) *geoip2.Reader {
	reader, err := geoip2.Open(cfg.DB)
	if err != nil {
		log.Fatal(err)
	}
	return reader
}
