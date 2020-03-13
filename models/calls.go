package models

import "time"

type CallResponse struct {
	AccountSid      string              `json:"account_sid"`
	Annotation      string              `json:"annotation"`
	AnsweredBy      string              `json:"answered_by"`
	APIVersion      string              `json:"api_version"`
	CallerName      string              `json:"caller_name"`
	DateCreated     string              `json:"date_created"`
	DateUpdated     string              `json:"date_updated"`
	Direction       string              `json:"direction"`
	Duration        string              `json:"duration"`
	EndTime         string              `json:"end_time"`
	ForwardedFrom   PhoneNumber         `json:"forwarded_from"`
	From            PhoneNumber         `json:"from"`
	FromFormatted   string              `json:"from_formatted"`
	GroupSid        string              `json:"group_sid"`
	ParentCallSid   string              `json:"parent_call_sid"`
	PhoneNumberSid  string              `json:"phone_number_sid"`
	Price           string              `json:"price"`
	PriceUnit       string              `json:"price_unit"`
	Sid             string              `json:"sid"`
	StartTime       string              `json:"start_time"`
	Status          string              `json:"status"`
	SubresourceUris CallSubresourceUris `json:"subresource_uris"`
	To              PhoneNumber         `json:"to"`
	ToFormatted     string              `json:"to_formatted"`
	TrunkSid        string              `json:"trunk_sid"`
	URI             string              `json:"uri"`
	QueueTime       string              `json:"queue_time"`
}

type CallSubresourceUris struct {
	Notifications     string `json:"notifications"`
	Recordings        string `json:"recordings"`
	Feedback          string `json:"feedback"`
	FeedbackSummaries string `json:"feedback_summaries"`
	Payments          string `json:"payments"`
}

// DateCreatedAsTime returns CallResponse.DateCreated as a time.Time object
// instead of a string.
func (call *CallResponse) DateCreatedAsTime() (time.Time, error) {
	return time.Parse(time.RFC1123Z, call.DateCreated)
}

// DateUpdateAsTime returns CallResponse.DateUpdated as a time.Time object
// instead of a string.
func (call *CallResponse) DateUpdatedAsTime() (time.Time, error) {
	return time.Parse(time.RFC1123Z, call.DateUpdated)
}
