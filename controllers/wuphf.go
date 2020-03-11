package controllers

import (
	"fmt"

	"golang.org/x/sync/errgroup"
)

type WuphfParams struct {
	Message   string
	SMSNumber string
}

func (c *Controllers) SendWuphf(params *WuphfParams) error {

	var err error
	var g errgroup.Group

	if params.SMSNumber != "" {
		g.Go(func() error {
			_, err = c.Twilio.SendSMS(params.SMSNumber, params.Message)
			return err
		})
	}

	if err = g.Wait(); err == nil {
		fmt.Println("finished clean") // TODO use zap
	} else {
		fmt.Printf("received error: %v", err) // TODO use zap
	}

	return err
}
