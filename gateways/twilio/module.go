package twilio

import (
	"net/http"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

const (
	baseURL           = "https://api.twilio.com/2010-04-01"
	messagesEndpoint  = "/Messages.json"
	headerAccept      = "application/json"
	headerContentType = "application/x-www-form-urlencoded"
	twilioNumber      = "+1 205-719-8997"
)

var successCodes = map[int]bool{
	200: true,
	201: true,
}

type Twilio interface {
	SendSMS(toNumber string, message string) (*SmsResponse, error)
}

type TwilioClient struct {
	httpClient *http.Client
	logger     *zap.Logger
}

func NewTwilioClient() *TwilioClient {
	logger, _ := zap.NewProduction()
	return &TwilioClient{
		httpClient: &http.Client{},
		logger:     logger,
	}
}

func getAuthCredentials() (string, string) {
	accountSid := viper.GetString("twilio.account_sid")
	authToken := viper.GetString("twilio.auth_token")
	return accountSid, authToken
}
