package services

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
	"os"
	"strings"
)

type EmailService struct {
	smtpHost     string
	smtpPort     string
	smtpUsername string
	smtpPassword string
	fromEmail    string
	fromName     string
}

type EmailData struct {
	ToEmail       string
	ToName        string
	Subject       string
	TemplateName  string
	TemplateData  map[string]interface{}
	PlainTextBody string
}

func NewEmailService() *EmailService {
	return &EmailService{
		smtpHost:     os.Getenv("SMTP_HOST"),
		smtpPort:     os.Getenv("SMTP_PORT"),
		smtpUsername: os.Getenv("SMTP_USERNAME"),
		smtpPassword: os.Getenv("SMTP_PASSWORD"),
		fromEmail:    os.Getenv("SMTP_FROM_EMAIL"),
		fromName:     os.Getenv("SMTP_FROM_NAME"),
	}
}

// SendEmail sends an email using SMTP
func (s *EmailService) SendEmail(emailData EmailData) error {
	// Validate SMTP configuration
	if err := s.validateConfig(); err != nil {
		return fmt.Errorf("invalid SMTP configuration: %w", err)
	}

	// Build email message
	message, err := s.buildEmailMessage(emailData)
	if err != nil {
		return fmt.Errorf("failed to build email message: %w", err)
	}

	// Setup authentication
	auth := smtp.PlainAuth("", s.smtpUsername, s.smtpPassword, s.smtpHost)

	// Setup TLS config
	tlsConfig := &tls.Config{
		ServerName: s.smtpHost,
	}

	// Connect to SMTP server
	addr := fmt.Sprintf("%s:%s", s.smtpHost, s.smtpPort)
	conn, err := tls.Dial("tcp", addr, tlsConfig)
	if err != nil {
		// Try without TLS
		return s.sendWithoutTLS(auth, emailData.ToEmail, message)
	}
	defer conn.Close()

	// Create SMTP client
	client, err := smtp.NewClient(conn, s.smtpHost)
	if err != nil {
		return fmt.Errorf("failed to create SMTP client: %w", err)
	}
	defer client.Close()

	// Authenticate
	if err := client.Auth(auth); err != nil {
		return fmt.Errorf("SMTP authentication failed: %w", err)
	}

	// Set sender
	if err := client.Mail(s.fromEmail); err != nil {
		return fmt.Errorf("failed to set sender: %w", err)
	}

	// Set recipient
	if err := client.Rcpt(emailData.ToEmail); err != nil {
		return fmt.Errorf("failed to set recipient: %w", err)
	}

	// Send email body
	w, err := client.Data()
	if err != nil {
		return fmt.Errorf("failed to send email data: %w", err)
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		return fmt.Errorf("failed to write email message: %w", err)
	}

	err = w.Close()
	if err != nil {
		return fmt.Errorf("failed to close email writer: %w", err)
	}

	log.Printf("Email sent successfully to %s", emailData.ToEmail)
	return nil
}

// sendWithoutTLS sends email without TLS (fallback method)
func (s *EmailService) sendWithoutTLS(auth smtp.Auth, to string, message string) error {
	addr := fmt.Sprintf("%s:%s", s.smtpHost, s.smtpPort)
	err := smtp.SendMail(addr, auth, s.fromEmail, []string{to}, []byte(message))
	if err != nil {
		return fmt.Errorf("failed to send email without TLS: %w", err)
	}
	log.Printf("Email sent successfully (without TLS) to %s", to)
	return nil
}

// SendBulkEmail sends email to multiple recipients
func (s *EmailService) SendBulkEmail(recipients []EmailData) []error {
	var errors []error
	successCount := 0

	for _, emailData := range recipients {
		err := s.SendEmail(emailData)
		if err != nil {
			log.Printf("Failed to send email to %s: %v", emailData.ToEmail, err)
			errors = append(errors, fmt.Errorf("failed to send to %s: %w", emailData.ToEmail, err))
		} else {
			successCount++
		}
	}

	log.Printf("Bulk email completed: %d successful, %d failed", successCount, len(errors))
	return errors
}

// buildEmailMessage constructs the email message with headers
func (s *EmailService) buildEmailMessage(emailData EmailData) (string, error) {
	var body string
	var err error

	// Use template if provided, otherwise use plain text
	if emailData.TemplateName != "" {
		body, err = s.renderTemplate(emailData.TemplateName, emailData.TemplateData)
		if err != nil {
			return "", err
		}
	} else {
		body = emailData.PlainTextBody
	}

	// Build email headers
	headers := make(map[string]string)
	headers["From"] = fmt.Sprintf("%s <%s>", s.fromName, s.fromEmail)
	headers["To"] = emailData.ToEmail
	headers["Subject"] = emailData.Subject
	headers["MIME-Version"] = "1.0"
	headers["Content-Type"] = "text/html; charset=UTF-8"

	// Construct message
	var message strings.Builder
	for key, value := range headers {
		message.WriteString(fmt.Sprintf("%s: %s\r\n", key, value))
	}
	message.WriteString("\r\n")
	message.WriteString(body)

	return message.String(), nil
}

// renderTemplate renders an HTML template with data
func (s *EmailService) renderTemplate(templateName string, data map[string]interface{}) (string, error) {
	templatePath := fmt.Sprintf("templates/email/%s.html", templateName)
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}

	return buf.String(), nil
}

// validateConfig checks if all required SMTP configuration is present
func (s *EmailService) validateConfig() error {
	if s.smtpHost == "" {
		return fmt.Errorf("SMTP_HOST is not set")
	}
	if s.smtpPort == "" {
		return fmt.Errorf("SMTP_PORT is not set")
	}
	if s.smtpUsername == "" {
		return fmt.Errorf("SMTP_USERNAME is not set")
	}
	if s.smtpPassword == "" {
		return fmt.Errorf("SMTP_PASSWORD is not set")
	}
	if s.fromEmail == "" {
		return fmt.Errorf("SMTP_FROM_EMAIL is not set")
	}
	if s.fromName == "" {
		return fmt.Errorf("SMTP_FROM_NAME is not set")
	}
	return nil
}

// TestConnection tests the SMTP connection
func (s *EmailService) TestConnection() error {
	if err := s.validateConfig(); err != nil {
		return err
	}

	auth := smtp.PlainAuth("", s.smtpUsername, s.smtpPassword, s.smtpHost)
	addr := fmt.Sprintf("%s:%s", s.smtpHost, s.smtpPort)

	// Try connecting
	client, err := smtp.Dial(addr)
	if err != nil {
		return fmt.Errorf("failed to connect to SMTP server: %w", err)
	}
	defer client.Close()

	// Try authenticating
	if err := client.Auth(auth); err != nil {
		return fmt.Errorf("SMTP authentication failed: %w", err)
	}

	log.Println("SMTP connection test successful")
	return nil
}

// RenderTemplate renders an HTML template with data (public method)
func (s *EmailService) RenderTemplate(templateName string, data map[string]interface{}) (string, error) {
	return s.renderTemplate(templateName, data)
}

// GetTemplateContent reads and returns the template file content
func (s *EmailService) GetTemplateContent(templateName string) (string, error) {
	templatePath := fmt.Sprintf("templates/email/%s.html", templateName)
	content, err := os.ReadFile(templatePath)
	if err != nil {
		return "", fmt.Errorf("failed to read template: %w", err)
	}
	return string(content), nil
}

// SaveTemplateContent saves the template content to file
func (s *EmailService) SaveTemplateContent(templateName string, content string) error {
	templatePath := fmt.Sprintf("templates/email/%s.html", templateName)

	// Create backup of existing template
	backupPath := fmt.Sprintf("templates/email/%s.backup.html", templateName)
	existingContent, err := os.ReadFile(templatePath)
	if err == nil {
		if err := os.WriteFile(backupPath, existingContent, 0644); err != nil {
			log.Printf("Warning: failed to create backup: %v", err)
		}
	}

	// Write new template content
	if err := os.WriteFile(templatePath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to save template: %w", err)
	}

	log.Printf("Template %s updated successfully", templateName)
	return nil
}
