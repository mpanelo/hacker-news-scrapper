package data

import (
	"context"
	"database/sql"
	"time"

	"github.com/lib/pq"
)

type Story struct {
	ID          int
	Type        string
	By          string
	Time        time.Time
	Kids        []int
	URL         string
	Score       int
	Title       string
	Descendants int
}

type StoryModel struct {
	DB *sql.DB
}

func (sm StoryModel) Insert(story Story) error {
	query := `
        INSERT INTO stories (id, type, by, time, kids, url, score, title, descendants)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	args := []any{
		story.ID,
		story.Type,
		story.By,
		story.Time,
		pq.Array(story.Kids),
		story.URL,
		story.Score,
		story.Title,
		story.Descendants,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := sm.DB.ExecContext(ctx, query, args...)
	return err
}
