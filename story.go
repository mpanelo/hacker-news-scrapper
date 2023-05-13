package main

import (
	"github.com/mpanelo/hacker-news-scraper/internal/conv"
	"github.com/mpanelo/hacker-news-scraper/internal/hn"
	"github.com/rs/zerolog/log"
)

func (a application) SaveTopStories() error {
	items, err := a.topStories()
	if err != nil {
		return err
	}

	for _, item := range items {
		story := conv.ItemToStory(item)

		err := a.models.Stories.Insert(story)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a application) topStories() ([]hn.Item, error) {
	itemIDs, err := a.hnClient.TopStories()
	if err != nil {
		return nil, err
	}

	var items []hn.Item
	for i, itemID := range itemIDs {
		log.Info().Msgf("fetching item %d (%d out of %d)", itemID, i+1, len(itemIDs))
		item, err := a.hnClient.Item(itemID)
		if err != nil {
			log.Err(err).Msgf("failed to get item %d", itemID)
		} else {
			items = append(items, item)
		}
	}

	return items, nil
}
