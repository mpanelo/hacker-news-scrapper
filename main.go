package main

import (
	"context"
	"database/sql"
	"time"

	"github.com/joho/godotenv"
	"github.com/mpanelo/hacker-news-scraper/internal/data"
	"github.com/mpanelo/hacker-news-scraper/internal/env"
	"github.com/mpanelo/hacker-news-scraper/internal/hn"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	_ "github.com/lib/pq"
)

type application struct {
	cfg      config
	hnClient hn.Client
	models   data.Models
}

type config struct {
	db struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
}

func main() {
	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// TODO: set via command line or env
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	err := godotenv.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load .env file")
	}

	var cfg config
	cfg.db.dsn = env.ReadString("HACKER_NEWS_DB_DSN")
	cfg.db.maxOpenConns = env.ReadInt("DB_MAX_OPEN_CONNS")
	cfg.db.maxIdleConns = env.ReadInt("DB_MAX_IDLE_CONNS")
	cfg.db.maxIdleTime = env.ReadString("DB_MAX_IDLE_TIME")

	db, err := openDB(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to open DB")
	}

	defer db.Close()

	app := NewApplication(cfg, db)
	if err := app.SaveTopStories(); err != nil {
		log.Fatal().Err(err).Msg("failed to save top hacker news stories")
	}
}

func NewApplication(cfg config, db *sql.DB) *application {
	return &application{
		cfg:      cfg,
		hnClient: hn.NewClient(),
		models:   data.NewModels(db),
	}
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	// Set the maximum number of open (in-use + idle) connections in the pool. Note that
	// passing a value less than or equal to 0 will mean there is no limit.
	db.SetMaxOpenConns(cfg.db.maxOpenConns)

	// Set the maximum number of idle connections in the pool. Again, passing a value
	// less than or equal to 0 will mean there is no limit.
	db.SetMaxIdleConns(cfg.db.maxIdleConns)

	duration, err := time.ParseDuration(cfg.db.maxIdleTime)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
