package main

import "github.com/gorilla/feeds"

// FeedParser convert the feeds to RSS format
type FeedParser struct {
	feeds Feeds
	link  string
}

// Parse all feeds and convert it to a know format like RSS
func (f *FeedParser) Parse() (string, error) {
	feed := &feeds.Feed{
		Title:       "OSPG links",
		Link:        &feeds.Link{Href: f.link},
		Description: "Interesting links about tech, hacking and more!",
		Author:      &feeds.Author{Name: "Multiple authors"},
	}
	var items []*feeds.Item
	for _, value := range f.feeds {
		item := &feeds.Item{
			Title:       value.Name,
			Link:        &feeds.Link{Href: value.Name},
			Description: value.Author,
			Created:     value.CreatedAt,
			Author:      value.Author,
		}
		items = append(items, item)
	}
	feed.Items = items
	return feed.ToRss()
}
