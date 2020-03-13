package controllers

import (
	"fmt"

	"github.com/brandenc40/wuphf.com/models"
	"golang.org/x/sync/errgroup"
)

type WuphfParams struct {
	Message    string
	FromName   string
	SMSNumber  models.PhoneNumber
	CallNumber models.PhoneNumber
}

func (c *Controllers) SendWuphf(params *WuphfParams) error {

	var err error
	var g errgroup.Group

	// Send SMS
	if params.SMSNumber != "" {
		g.Go(func() error {
			_, err = c.Twilio.SendSMS(params.SMSNumber.Friendly(), params.FromName, params.Message)
			return err
		})
	}

	// Place call
	if params.CallNumber != "" {
		g.Go(func() error {
			_, err = c.Twilio.PlaceCall(params.CallNumber.Friendly(), params.FromName, params.Message)
			return err
		})
	}

	if err = g.Wait(); err != nil {
		fmt.Printf("received error: %v", err) // TODO use zap
	} else {
		fmt.Println("finished clean") // TODO remove
	}

	return err
}
