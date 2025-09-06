package services

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"vuka-api/pkg/models"
)

type RssService struct{}

func NewRssService() *RssService {
	return &RssService{}
}

func (s *RssService) IngestRSSFeed(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var rss models.RSS
	err = xml.Unmarshal(body, &rss)
	if err != nil {
		// Handle error
		return
	}

	fmt.Printf("Feed Title: %s\n", rss.Channel.Title)
	fmt.Printf("Number of Items: %d\n", len(rss.Channel.Items))
	article, err := rss.Channel.Items[0].ToArticle()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Article desc: %s", article.OriginalUrl)
}
