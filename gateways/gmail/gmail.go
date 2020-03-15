package gmail

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Gmail interface {
	SendEmail(toEmail string, toName string, message string) error
}

type GmailClient struct {
	logger *zap.Logger
}

func NewGmailClient() *GmailClient {
	logger, _ := zap.NewProduction()
	return &GmailClient{
		logger: logger,
	}
}

// smtpServer data to smtp server
type smtpServer struct {
	host string
	port string
}

// Address URI to smtp server
func (s *smtpServer) Address() string {
	return s.host + ":" + s.port
}

const (
	gmailHost = "smtp.gmail.com"
	gmailPort = "587"
)

func (g *GmailClient) SendEmail(toEmail string, fromName string, message string) error {
	smtpServer := smtpServer{host: gmailHost, port: gmailPort}
	toEmailSlice := []string{toEmail}
	fromEmail := os.Getenv("GMAIL_ADDRESS")

	// Authentication.
	auth := smtp.PlainAuth(
		"",
		fromEmail,
		os.Getenv("GMAIL_PASSWORD"),
		smtpServer.host,
	)

	// Sending email.
	err := smtp.SendMail(
		smtpServer.Address(),
		auth,
		fromEmail,
		toEmailSlice,
		composeMessage(fromName, toEmail, message),
	)
	if err != nil {
		g.logger.Error(
			"Error sending email",
			zap.Error(err),
		)
		return err
	}

	return nil
}

func composeMessage(from string, to string, body string) []byte {
	subjectText := viper.GetString("messages.email_subject")
	subjectText = fmt.Sprintf(subjectText, from)

	message := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subjectText + "\n\n" +
		body

	return []byte(message)
}
