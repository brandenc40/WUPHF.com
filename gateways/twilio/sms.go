package twilio

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/brandenc40/wuphf.com/models"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// SendSMS -
func (t *TwilioClient) SendSMS(toNumber string, fromName string, message string) (*models.SmsResponse, error) {

	// Set initial variables
	smsResponse := models.SmsResponse{}
	accountSid, authToken := getAuthCredentials()
	urlStr := baseURL + "/Accounts/" + accountSid + messagesEndpoint

	// Build out the data for our message
	values := url.Values{}
	values.Set("To", toNumber)
	values.Set("From", viper.GetString("twilio.phone_number"))
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

func buildSMSMessage(fromName string, message string) string {
	template := viper.GetString("messages.sms_template")
	return fmt.Sprintf(template, fromName, message)
}
