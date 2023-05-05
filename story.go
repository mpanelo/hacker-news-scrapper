package main

import (
	"time"

	"github.com/mpanelo/hacker-news-scraper/internal/hn"
	"github.com/rs/zerolog/log"
)

func (a *application) DiscoverStories() ([]hn.Item, error) {
	today := time.Now().Round(time.Hour)

	var items []hn.Item

	maxItem, err := a.hnClient.MaxItem()
	if err != nil {
		return items, err
	}

	itemID := maxItem

	for {
		item, err := a.hnClient.Item(itemID)
		if err != nil {
			return items, err
		}

		// TODO: create enum for type

		// Skip non-story items
		if item.Type != "story" {
			log.Debug().Msgf("skipping item %d of type %s", item.ID, item.Type)
			itemID--
			continue
		}

		// Skip item because we cannot determine if it happened today
		if item.Time > 0 {
			itemID--
			continue
		}

		t := time.Unix(int64(item.Time), 0)

		if today.After(t) {
			break
		}

		log.Debug().Msgf("found story %d", item.ID)

		items = append(items, item)
		itemID--

		// Be respectful of hacker news API
		time.Sleep(250 * time.Millisecond) // TODO: make an env variable
	}

	return items, nil
}
