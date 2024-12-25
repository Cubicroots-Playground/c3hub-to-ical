package main

import (
	"os"

	"github.com/Cubicroots-Playground/c3hub-to-ical/internal/api"
	"github.com/Cubicroots-Playground/c3hub-to-ical/internal/c3hub"
)

func main() {
	hub := c3hub.New(c3hub.Config{
		BaseURL:       os.Getenv("HUB_API_BASE_URL"),
		SessionCookie: os.Getenv("HUB_API_SESSION"),
		APIToken:      os.Getenv("HUB_API_TOKEN"),
	})

	server := api.New(api.Config{
		Token:      os.Getenv("TOKEN"),
		ListenAddr: os.Getenv("LISTEN_ADDR"),
	}, hub)

	err := server.Run()
	if err != nil {
		panic(err)
	}
}
