package twilio

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/brandenc40/wuphf.com/models"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// SendSMS -
func (t *TwilioClient) PlaceCall(toNumber string, fromName string, message string) (*models.CallResponse, error) {

	// Set initial variables
	callResponse := models.CallResponse{}
	accountSid, authToken := getAuthCredentials()
	urlStr := baseURL + "/Accounts/" + accountSid + callsEndpoint

	// Build out the data for our message
	values := url.Values{}
	values.Set("To", toNumber)
	values.Set("From", viper.GetString("twilio.phone_number"))
	values.Set("Url", MakeCallUrl(fromName, message))

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
			"Twilio returned a non 200 status code when placing Call",
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
	err = json.Unmarshal(responseBody, &callResponse)
	if err != nil {
		t.logger.Error(
			"Unable to unmarshal response body to struct",
			zap.Error(err),
		)
		return nil, err
	}
	return &callResponse, nil
}

func MakeCallUrl(fromName string, message string) string {
	baseUrl := viper.GetString("twilio.call_url")
	u, _ := url.Parse(baseUrl)
	q, _ := url.ParseQuery(u.RawQuery)
	q.Add("from_name", fromName)
	q.Add("message", message)
	u.RawQuery = q.Encode()
	return u.String()
}
