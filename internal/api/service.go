package api

import (
	"log/slog"
	"net/http"

	"github.com/Cubicroots-Playground/c3hub-to-ical/internal/c3hub"
)

type Config struct {
	Token      string
	ListenAddr string
}

type service struct {
	conf       Config
	hubService c3hub.Service
}

func New(conf Config, hubService c3hub.Service) Service {
	return &service{
		conf:       conf,
		hubService: hubService,
	}
}

func (service *service) Run() error {
	slog.Info("listening on " + service.conf.ListenAddr)
	http.HandleFunc("/ical", service.MySchedule)
	return http.ListenAndServe(service.conf.ListenAddr, nil)
}
