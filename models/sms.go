package models

import "time"

type SmsResponse struct {
	AccountSid          string             `json:"account_sid"`
	APIVersion          string             `json:"api_version"`
	Body                string             `json:"body"`
	DateCreated         string             `json:"date_created"`
	DateSent            string             `json:"date_sent"`
	DateUpdated         string             `json:"date_updated"`
	Direction           string             `json:"direction"`
	ErrorCode           string             `json:"error_code"`
	ErrorMessage        string             `json:"error_message"`
	From                PhoneNumber        `json:"from"`
	MessagingServiceSid string             `json:"messaging_service_sid"`
	NumMedia            string             `json:"num_media"`
	NumSegments         string             `json:"num_segments"`
	Price               string             `json:"price"`
	PriceUnit           string             `json:"price_unit"`
	Sid                 string             `json:"sid"`
	Status              string             `json:"status"`
	SubresourceUris     SMSSubresourceUris `json:"subresource_uris"`
	To                  PhoneNumber        `json:"to"`
	URI                 string             `json:"uri"`
}

type SMSSubresourceUris struct {
	Media string `json:"media"`
}

// DateCreatedAsTime returns SmsResponse.DateCreated as a time.Time object
// instead of a string.
func (sms *SmsResponse) DateCreatedAsTime() (time.Time, error) {
	return time.Parse(time.RFC1123Z, sms.DateCreated)
}

// DateUpdateAsTime returns SmsResponse.DateUpdated as a time.Time object
// instead of a string.
func (sms *SmsResponse) DateUpdatedAsTime() (time.Time, error) {
	return time.Parse(time.RFC1123Z, sms.DateUpdated)
}

// DateSentAsTime returns SmsResponse.DateSent as a time.Time object
// instead of a string.
func (sms *SmsResponse) DateSentAsTime() (time.Time, error) {
	return time.Parse(time.RFC1123Z, sms.DateSent)
}
