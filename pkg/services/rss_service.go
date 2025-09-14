package services

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"vuka-api/pkg/models"

	"github.com/google/uuid"
)

type RssService struct {
	articleService *ArticleService
}

func NewRssService(articleService *ArticleService) *RssService {
	return &RssService{
		articleService: articleService,
	}
}

func (s *RssService) IngestRSSFeed(url string) error {
	return s.IngestRSSFeedWithSource(url, nil)
}

func (s *RssService) IngestRSSFeedWithSource(url string, sourceID *uuid.UUID) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch RSS feed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	var rss models.RSS
	err = xml.Unmarshal(body, &rss)
	if err != nil {
		return fmt.Errorf("failed to parse RSS feed: %w", err)
	}

	fmt.Printf("Feed Title: %s\n", rss.Channel.Title)
	fmt.Printf("Number of Items: %d\n", len(rss.Channel.Items))

	savedCount := 0
	duplicateCount := 0

	// Save all articles from the RSS feed
	for i, item := range rss.Channel.Items {
		article, err := item.ToArticle()
		if err != nil {
			log.Printf("Failed to convert RSS item %d to article: %v", i, err)
			continue
		}

		// Set the source ID if provided
		if sourceID != nil {
			article.SourceID = sourceID
		}

		created, err := s.articleService.CreateArticleIfNotExists(article)
		if err != nil {
			log.Printf("Failed to save article '%s': %v", article.Title, err)
			continue
		}

		if created {
			fmt.Printf("Successfully saved article: %s\n", article.Title)
			savedCount++
		} else {
			log.Printf("Article already exists, skipping: %s", article.Title)
			duplicateCount++
		}
	}

	fmt.Printf("RSS ingestion completed. New articles: %d, Duplicates skipped: %d\n", savedCount, duplicateCount)

	return nil
}
