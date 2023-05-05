package main

import (
	"fmt"

	"github.com/mpanelo/hacker-news-scraper/internal/hn"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type application struct {
	hnClient hn.Client
}

func main() {
	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// TODO: set via command line or env
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	app := NewApplication()
	items, err := app.DiscoverStories()
	if err != nil {
		log.Err(err).Msg("failed to discover times")
	}

	for _, item := range items {
		fmt.Println(item.Title)
	}
}

func NewApplication() *application {
	return &application{
		hnClient: hn.NewClient(),
	}
}
