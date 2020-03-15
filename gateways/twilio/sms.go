package twilio

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/brandenc40/wuphf.com/models"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// SendSMS -
func (t *TwilioClient) SendSMS(toNumber string, fromName string, message string) (*http.Response, error) {

	// Set initial variables
	accountSid, authToken := getAuthCredentials()
	urlStr := baseURL + "/Accounts/" + accountSid + messagesEndpoint

	// Build out the data for our message
	values := url.Values{}
	values.Set("To", toNumber)
	values.Set("From", os.Getenv("TWILIO_PHONE_NUMBER"))
	values.Set("Body", buildSMSMessage(fromName, message))

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
		errStruct := models.TwilioError{}
		json.NewDecoder(resp.Body).Decode(&errStruct)
		t.logger.Error(
			"Twilio returned a non 200 status code when sending SMS",
			zap.String("response_status", resp.Status),
			zap.String("error_message", errStruct.Message),
			zap.Int("error_code", errStruct.Code),
		)
		return nil, errors.New(string(errStruct.Message))
	} else if err != nil {
		t.logger.Error(
			"Error sending sms",
			zap.Error(err),
		)
		return nil, err
	}

	return resp, nil
}

func (t *TwilioClient) UnmarshalSmsResponse(resp *http.Response) (*models.SmsResponse, error) {
	smsResponse := models.SmsResponse{}

	if err := json.NewDecoder(resp.Body).Decode(&smsResponse); err != nil {
		t.logger.Error(
			"Unable to unmarshal response body to struct",
			zap.Error(err),
		)
		return nil, err
	}
	return &smsResponse, nil
}

func buildSMSMessage(fromName string, message string) string {
	template := viper.GetString("messages.sms_template")
	return fmt.Sprintf(template, fromName, message)
}
