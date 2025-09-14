package services

import (
	"fmt"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

type CronService struct {
	cron          *cron.Cron
	rssService    *RssService
	sourceService *SourceService
}

func NewCronService(rssService *RssService, sourceService *SourceService) *CronService {
	// Create cron with second precision and logging
	c := cron.New(cron.WithSeconds(), cron.WithLogger(cron.VerbosePrintfLogger(log.New(log.Writer(), "CRON: ", log.LstdFlags))))

	return &CronService{
		cron:          c,
		rssService:    rssService,
		sourceService: sourceService,
	}
}

// Start starts the cron scheduler
func (s *CronService) Start() {
	log.Println("Starting cron service...")
	s.cron.Start()
}

// Stop stops the cron scheduler
func (s *CronService) Stop() {
	log.Println("Stopping cron service...")
	s.cron.Stop()
}

// ScheduleRSSIngestion schedules RSS feed ingestion to run at specified intervals
func (s *CronService) ScheduleRSSIngestion() error {
	// Schedule to run every hour
	_, err := s.cron.AddFunc("0 0 * * * *", s.ingestAllRSSFeeds)
	if err != nil {
		return err
	}

	log.Println("RSS ingestion scheduled to run every hour")
	return nil
}

// ScheduleRSSIngestionDaily schedules RSS feed ingestion to run daily at a specific time
func (s *CronService) ScheduleRSSIngestionDaily(hour, minute int) error {
	// Schedule to run daily at specified time (e.g., "0 30 8 * * *" for 8:30 AM daily)
	cronSpec := fmt.Sprintf("0 %d %d * * *", minute, hour)
	_, err := s.cron.AddFunc(cronSpec, s.ingestAllRSSFeeds)
	if err != nil {
		return err
	}

	log.Printf("RSS ingestion scheduled to run daily at %02d:%02d", hour, minute)
	return nil
}

// ingestAllRSSFeeds fetches all sources and ingests their RSS feeds
func (s *CronService) ingestAllRSSFeeds() {
	log.Println("Starting scheduled RSS feed ingestion...")
	start := time.Now()

	sources, err := s.sourceService.GetAllSources()
	if err != nil {
		log.Printf("Failed to get sources: %v", err)
		return
	}

	if len(sources) == 0 {
		log.Println("No sources found to ingest")
		return
	}

	log.Printf("Found %d sources to process", len(sources))

	successCount := 0
	errorCount := 0

	for _, source := range sources {
		if source.RssFeedUrl == "" {
			log.Printf("Skipping source '%s' - no RSS feed URL", source.Name)
			continue
		}

		log.Printf("Ingesting RSS feed for source: %s (%s)", source.Name, source.RssFeedUrl)

		err := s.rssService.IngestRSSFeedWithSource(source.RssFeedUrl, &source.ID)
		if err != nil {
			log.Printf("Failed to ingest RSS feed for source '%s': %v", source.Name, err)
			errorCount++
		} else {
			log.Printf("Successfully ingested RSS feed for source: %s", source.Name)
			successCount++
		}
	}

	duration := time.Since(start)
	log.Printf("RSS ingestion completed in %v. Success: %d, Errors: %d", duration, successCount, errorCount)
}

// TriggerRSSIngestionNow manually triggers RSS ingestion for all sources
func (s *CronService) TriggerRSSIngestionNow() {
	log.Println("Manually triggering RSS feed ingestion...")
	go s.ingestAllRSSFeeds()
}
