package models

import (
	"encoding/xml"
	"time"
	"vuka-api/pkg/models/db"
)

// The root of the RSS feed
type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

// The channel element containing feed metadata and items
type Channel struct {
	XMLName     xml.Name `xml:"channel"`
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	Description string   `xml:"description"`
	Items       []Item   `xml:"item"`
}

// An item represents a single article or post in the feed
type Item struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	Description string   `xml:"description"`
	GUID        string   `xml:"guid"`
	Author      string   `xml:"author"`
	PubDate     string   `xml:"pubDate"` // You might want to parse this into a time.Time later
	Categories  []string `xml:"category"`
}

func (feed *Item) ToArticle() (*db.Article, error) {
	publishedAt, err := time.Parse(time.RFC1123Z, feed.PubDate)
	if err != nil {
		return nil, err
	}
	article := &db.Article{
		Title:       feed.Title,
		Language:    "",
		OriginalUrl: feed.Link,
		ContentBody: feed.Description,
		PublishedAt: publishedAt,
		IsFeatured:  false,
	}

	return article, nil
}