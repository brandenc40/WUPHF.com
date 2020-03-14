package controllers

import (
	"errors"

	"github.com/brandenc40/wuphf.com/models"
	"golang.org/x/sync/errgroup"
)

type WuphfParams struct {
	Message    string             `json:"message",form:"message"`
	FromName   string             `json:"from_name",form:"from_name"`
	SMSNumber  models.PhoneNumber `json:"sms_number",form:"sms_number"`
	CallNumber models.PhoneNumber `json:"call_number",form:"call_number"`
	ToEmail    string             `json:"to_email",form:"to_email"`
}

func (c *Controllers) SendWuphf(params *WuphfParams) error {

	if err := validateParams(params); err != nil {
		return err
	}

	var g errgroup.Group

	// Send SMS
	if params.SMSNumber != "" {
		g.Go(func() error {
			_, err := c.gateways.Twilio.SendSMS(params.SMSNumber.Friendly(), params.FromName, params.Message)
			return err
		})
	}

	// Place call
	if params.CallNumber != "" {
		g.Go(func() error {
			_, err := c.gateways.Twilio.PlaceCall(params.CallNumber.Friendly(), params.FromName, params.Message)
			return err
		})
	}

	// Send Email
	if params.ToEmail != "" {
		g.Go(func() error {
			err := c.gateways.Gmail.SendEmail(params.ToEmail, params.FromName, params.Message)
			return err
		})
	}

	if err := g.Wait(); err != nil {
		return err
	}

	return nil
}

func validateParams(params *WuphfParams) error {
	if params.ToEmail == "" && params.CallNumber == "" && params.SMSNumber == "" {
		return errors.New("Must include an email, call phone, or sms phone.")
	}
	return nil
}
