package main

import (
	"log"
	"time"
	"vuka-api/pkg/config"
	"vuka-api/pkg/services"
)

// Example demonstrating how to use the newsletter service programmatically

func main() {
	// Initialize
	config.LoadEnvVariables()
	config.Connect()

	// Create service manager
	serviceManager := services.NewServices(config.GetDB())
	newsletterService := serviceManager.Newsletter

	// Example 1: Test SMTP Connection
	log.Println("Testing SMTP connection...")
	if err := newsletterService.TestSMTPConnection(); err != nil {
		log.Fatalf("SMTP connection failed: %v", err)
	}
	log.Println("✓ SMTP connection successful")

	// Example 2: Send Test Email
	log.Println("\nSending test email...")
	if err := newsletterService.SendTestEmail("test@example.com", "Test User"); err != nil {
		log.Fatalf("Failed to send test email: %v", err)
	}
	log.Println("✓ Test email sent")

	// Example 3: Send Newsletter with Featured Articles
	log.Println("\nSending newsletter with featured articles...")
	subject := "Vuka Newsletter - " + time.Now().Format("January 2, 2006")
	if err := newsletterService.SendNewsletterWithLatestArticles(subject, 5); err != nil {
		log.Fatalf("Failed to send newsletter: %v", err)
	}
	log.Println("✓ Newsletter sent")

	// Example 4: Send Custom Newsletter
	log.Println("\nSending custom newsletter...")
	customContent := `
		<h1>Special Announcement</h1>
		<p>This is a custom newsletter with HTML content.</p>
		<p>You can include any HTML formatting you want!</p>
	`
	if err := newsletterService.SendNewsletter("Special Announcement", customContent, false, nil); err != nil {
		log.Fatalf("Failed to send custom newsletter: %v", err)
	}
	log.Println("✓ Custom newsletter sent")

	// Example 5: Using Cron Service for Scheduled Newsletters
	log.Println("\nSetting up scheduled newsletters...")
	cronService := serviceManager.Cron

	// Schedule weekly newsletter every Monday at 9 AM
	if err := cronService.ScheduleNewsletterWeekly(time.Monday, 9, 0); err != nil {
		log.Printf("Failed to schedule weekly newsletter: %v", err)
	}

	// Start cron service
	cronService.Start()
	log.Println("✓ Cron service started with scheduled newsletters")

	// Keep the program running to allow scheduled tasks to execute
	log.Println("\nPress Ctrl+C to stop...")

	// Block forever
	select {}
}
