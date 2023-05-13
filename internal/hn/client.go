package hn

import (
	"fmt"
	"strconv"

	"github.com/go-resty/resty/v2"
)

const (
	baseURL        = "https://hacker-news.firebaseio.com"
	version        = "v0"
	maxItemPath    = "/maxitem.json"
	itemPath       = "/item"
	topStoriesPath = "/topstories.json"
)

type Client interface {
	MaxItem() (int, error)
	Item(int) (Item, error)
	TopStories() ([]int, error)
}

type hnClient struct {
	client *resty.Client
}

type Item struct {
	ID          int    `json:"id"`
	Deleted     bool   `json:"deleted"`
	Type        string `json:"type,omitempty"`
	By          string `json:"by,omitempty"`
	Time        int    `json:"time,omitempty"`
	Text        string `json:"text,omitempty"`
	Dead        bool   `json:"dead,omitempty"`
	Parent      int    `json:"parent,omitempty"`
	Poll        int    `json:"poll,omitempty"`
	Kids        []int  `json:"kids,omitempty"`
	URL         string `json:"url,omitempty"`
	Score       int    `json:"score,omitempty"`
	Title       string `json:"title,omitempty"`
	Parts       []int  `json:"parts,omitempty"`
	Descendants int    `json:"descendants,omitempty"`
}

func NewClient() Client {
	client := resty.New()
	client.SetBaseURL(fmt.Sprintf("%s/%s", baseURL, version))

	return &hnClient{
		client: client,
	}
}

func (c *hnClient) MaxItem() (int, error) {
	var itemID int

	_, err := c.client.R().
		SetResult(&itemID).
		Get(maxItemPath)

	return itemID, err
}

func (c *hnClient) Item(id int) (Item, error) {
	var item Item

	_, err := c.client.R().
		SetPathParam("itemID", strconv.Itoa(id)).
		SetResult(&item).
		Get(itemPath + "/{itemID}.json")

	return item, err
}

func (c *hnClient) TopStories() ([]int, error) {
	var itemIDs []int

	_, err := c.client.R().
		SetResult(&itemIDs).
		Get(topStoriesPath)

	return itemIDs, err
}
