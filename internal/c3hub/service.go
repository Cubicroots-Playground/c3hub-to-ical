package c3hub

import "time"

type Config struct {
	BaseURL string
	Day1    time.Time

	SessionCookie string
}

type service struct {
	config Config
}

func New(conf Config) Service {
	return &service{
		config: conf,
	}
}
