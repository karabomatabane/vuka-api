package services

import (
	"fmt"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

type CronService struct {
	cron              *cron.Cron
	rssService        *RssService
	sourceService     *SourceService
	newsletterService *NewsletterService
}

func NewCronService(rssService *RssService, sourceService *SourceService, newsletterService *NewsletterService) *CronService {
	// Create cron with second precision and logging
	c := cron.New(cron.WithSeconds(), cron.WithLogger(cron.VerbosePrintfLogger(log.New(log.Writer(), "CRON: ", log.LstdFlags))))

	return &CronService{
		cron:              c,
		rssService:        rssService,
		sourceService:     sourceService,
		newsletterService: newsletterService,
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

// ScheduleNewsletterWeekly schedules newsletter to be sent weekly
func (s *CronService) ScheduleNewsletterWeekly(dayOfWeek time.Weekday, hour, minute int) error {
	// Cron day of week: 0 = Sunday, 6 = Saturday
	cronSpec := fmt.Sprintf("0 %d %d * * %d", minute, hour, dayOfWeek)
	_, err := s.cron.AddFunc(cronSpec, s.sendWeeklyNewsletter)
	if err != nil {
		return err
	}

	log.Printf("Weekly newsletter scheduled for %s at %02d:%02d", dayOfWeek.String(), hour, minute)
	return nil
}

// ScheduleNewsletterDaily schedules newsletter to be sent daily
func (s *CronService) ScheduleNewsletterDaily(hour, minute int) error {
	cronSpec := fmt.Sprintf("0 %d %d * * *", minute, hour)
	_, err := s.cron.AddFunc(cronSpec, s.sendDailyNewsletter)
	if err != nil {
		return err
	}

	log.Printf("Daily newsletter scheduled at %02d:%02d", hour, minute)
	return nil
}

// ScheduleNewsletterMonthly schedules newsletter to be sent monthly on a specific day
func (s *CronService) ScheduleNewsletterMonthly(dayOfMonth, hour, minute int) error {
	cronSpec := fmt.Sprintf("0 %d %d %d * *", minute, hour, dayOfMonth)
	_, err := s.cron.AddFunc(cronSpec, s.sendMonthlyNewsletter)
	if err != nil {
		return err
	}

	log.Printf("Monthly newsletter scheduled for day %d at %02d:%02d", dayOfMonth, hour, minute)
	return nil
}

// sendWeeklyNewsletter sends the weekly newsletter with featured articles
func (s *CronService) sendWeeklyNewsletter() {
	log.Println("Starting scheduled weekly newsletter...")
	start := time.Now()

	subject := fmt.Sprintf("Vuka Weekly Newsletter - %s", time.Now().Format("January 2, 2006"))
	err := s.newsletterService.SendNewsletterWithLatestArticles(subject, 10)
	if err != nil {
		log.Printf("Failed to send weekly newsletter: %v", err)
		return
	}

	duration := time.Since(start)
	log.Printf("Weekly newsletter sent successfully in %v", duration)
}

// sendDailyNewsletter sends the daily newsletter with featured articles
func (s *CronService) sendDailyNewsletter() {
	log.Println("Starting scheduled daily newsletter...")
	start := time.Now()

	subject := fmt.Sprintf("Vuka Daily Newsletter - %s", time.Now().Format("January 2, 2006"))
	err := s.newsletterService.SendNewsletterWithLatestArticles(subject, 5)
	if err != nil {
		log.Printf("Failed to send daily newsletter: %v", err)
		return
	}

	duration := time.Since(start)
	log.Printf("Daily newsletter sent successfully in %v", duration)
}

// sendMonthlyNewsletter sends the monthly newsletter with featured articles
func (s *CronService) sendMonthlyNewsletter() {
	log.Println("Starting scheduled monthly newsletter...")
	start := time.Now()

	subject := fmt.Sprintf("Vuka Monthly Newsletter - %s", time.Now().Format("January 2006"))
	err := s.newsletterService.SendNewsletterWithLatestArticles(subject, 20)
	if err != nil {
		log.Printf("Failed to send monthly newsletter: %v", err)
		return
	}

	duration := time.Since(start)
	log.Printf("Monthly newsletter sent successfully in %v", duration)
}

// TriggerNewsletterNow manually triggers newsletter sending
func (s *CronService) TriggerNewsletterNow(subject string, articleLimit int) error {
	log.Println("Manually triggering newsletter...")
	return s.newsletterService.SendNewsletterWithLatestArticles(subject, articleLimit)
}
