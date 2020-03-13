package gmail

import (
	"fmt"
	"net/smtp"

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

func (g *GmailClient) SendEmail(toEmail string, toName string, message string) error {
	smtpServer := smtpServer{host: gmailHost, port: gmailPort}
	toEmailSlice := []string{toEmail}
	fromEmail := viper.GetString("gmail.email")

	// Authentication.
	auth := smtp.PlainAuth(
		"",
		fromEmail,
		viper.GetString("gmail.password"),
		smtpServer.host,
	)

	// Sending email.
	err := smtp.SendMail(
		smtpServer.Address(),
		auth,
		fromEmail,
		toEmailSlice,
		buildEmailMessage(toName, message),
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

func buildEmailMessage(fromName string, message string) []byte {
	template := viper.GetString("messages.email_template")
	formatted := fmt.Sprintf(template, fromName, message)
	return []byte(formatted)
}
