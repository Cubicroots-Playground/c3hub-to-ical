package c3hub

import (
	"context"
	"net/http"
	"time"
)

type Config struct {
	BaseURL string
	Day1    time.Time

	SessionCookie string
	APIToken      string
}

type service struct {
	config Config
}

func New(conf Config) Service {
	return &service{
		config: conf,
	}
}

func (service *service) injectAuth(ctx context.Context, r *http.Request) *http.Request {
	if service.config.APIToken != "" {
		r.Header.Add("Authorization", "Token "+service.config.APIToken)
	} else {
		r.Header.Add("Cookie", "38C3_SESSION="+service.config.SessionCookie)
	}

	return r.WithContext(ctx)
}
