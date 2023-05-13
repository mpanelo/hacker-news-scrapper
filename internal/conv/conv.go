package conv

import (
	"time"

	"github.com/mpanelo/hacker-news-scraper/internal/data"
	"github.com/mpanelo/hacker-news-scraper/internal/hn"
)

func ItemToStory(item hn.Item) data.Story {
	t := time.Unix(int64(item.Time), 0)

	story := data.Story{
		ID:          item.ID,
		Type:        item.Type,
		By:          item.By,
		Time:        t,
		Kids:        item.Kids,
		URL:         item.URL,
		Score:       item.Score,
		Title:       item.Title,
		Descendants: item.Descendants,
	}

	if story.Kids == nil {
		story.Kids = []int{}
	}

	return story
}
