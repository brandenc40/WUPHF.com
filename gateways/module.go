package gateways

import (
	"github.com/brandenc40/wuphf.com/gateways/gmail"
	"github.com/brandenc40/wuphf.com/gateways/twilio"
)

type Gateway struct {
	Twilio *twilio.TwilioClient
	Gmail  *gmail.GmailClient
}

func New() *Gateway {
	return &Gateway{
		Twilio: twilio.NewTwilioClient(),
		Gmail:  gmail.NewGmailClient(),
	}
}
