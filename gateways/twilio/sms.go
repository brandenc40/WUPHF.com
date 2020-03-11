package twilio

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"go.uber.org/zap"
)

// SmsResponse is returned after a text/sms message is posted to Twilio
type SmsResponse struct {
	Sid         string  `json:"sid"`
	DateCreated string  `json:"date_created"`
	DateUpdate  string  `json:"date_updated"`
	DateSent    string  `json:"date_sent"`
	AccountSid  string  `json:"account_sid"`
	To          string  `json:"to"`
	From        string  `json:"from"`
	NumMedia    string  `json:"num_media"`
	Body        string  `json:"body"`
	Status      string  `json:"status"`
	Direction   string  `json:"direction"`
	ApiVersion  string  `json:"api_version"`
	Price       *string `json:"price,omitempty"`
	Url         string  `json:"uri"`
}

// DateCreatedAsTime returns SmsResponse.DateCreated as a time.Time object
// instead of a string.
func (sms *SmsResponse) DateCreatedAsTime() (time.Time, error) {
	return time.Parse(time.RFC1123Z, sms.DateCreated)
}

// DateUpdateAsTime returns SmsResponse.DateUpdate as a time.Time object
// instead of a string.
func (sms *SmsResponse) DateUpdateAsTime() (time.Time, error) {
	return time.Parse(time.RFC1123Z, sms.DateUpdate)
}

// DateSentAsTime returns SmsResponse.DateSent as a time.Time object
// instead of a string.
func (sms *SmsResponse) DateSentAsTime() (time.Time, error) {
	return time.Parse(time.RFC1123Z, sms.DateSent)
}

// SendSMS -
func (t *TwilioClient) SendSMS(toNumber string, message string) (*SmsResponse, error) {

	// Set initial variables
	smsResponse := SmsResponse{}
	accountSid, authToken := getAuthCredentials()
	urlStr := baseURL + "/Accounts/" + accountSid + messagesEndpoint

	// Build out the data for our message
	values := url.Values{}
	values.Set("To", toNumber)
	values.Set("From", twilioNumber)
	values.Set("Body", message)

	// Make the request
	req, err := http.NewRequest("POST", urlStr, strings.NewReader(values.Encode()))
	if err != nil {
		t.logger.Error(
			"Error creating POST request",
			zap.Error(err),
		)
		return nil, err
	}
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", headerAccept)
	req.Header.Add("Content-Type", headerContentType)

	// Handle the response
	resp, err := t.httpClient.Do(req)
	if !successCodes[resp.StatusCode] {
		errorBody, _ := ioutil.ReadAll(resp.Body)
		t.logger.Error(
			"Twilio returned a non 200 status code when sending SMS",
			zap.String("response_status", resp.Status),
			zap.ByteString("error_body", errorBody),
		)
		return nil, errors.New(string(errorBody))
	} else if err != nil {
		t.logger.Error(
			"Error sending sms",
			zap.Error(err),
		)
		return nil, err
	}
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.logger.Error(
			"Unable to read the reponse body to a string",
			zap.Error(err),
		)
		return nil, err
	}
	err = json.Unmarshal(responseBody, &smsResponse)
	if err != nil {
		t.logger.Error(
			"Unable to unmarshal response body to struct",
			zap.Error(err),
		)
		return nil, err
	}
	return &smsResponse, nil
}
