package gateways

import (
	"github.com/brandenc40/wuphf.com/gateways/twilio"
)

type Gateway struct {
	Twilio *twilio.TwilioClient
}

func New() *Gateway {
	return &Gateway{
		Twilio: twilio.NewTwilioClient(),
	}
}
