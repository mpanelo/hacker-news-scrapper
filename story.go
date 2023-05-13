package main

import (
	"github.com/mpanelo/hacker-news-scraper/internal/hn"
	"github.com/rs/zerolog/log"
)

func (a *application) TopStories() ([]hn.Item, error) {
	itemIDs, err := a.hnClient.TopStories()
	if err != nil {
		return nil, err
	}

	var items []hn.Item
	for _, itemID := range itemIDs {
		item, err := a.hnClient.Item(itemID)
		if err != nil {
			log.Err(err).Msgf("failed to get item %d", itemID)
		} else {
			items = append(items, item)
		}
	}

	return items, nil
}
