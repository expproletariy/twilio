package types

import "time"

type FieldTypeResponse struct {
	AssistantSid string         `json:"assistant_sid"`
	UniqueName   string         `json:"unique_name"`
	DateUpdated  time.Time      `json:"date_updated"`
	FriendlyName string         `json:"friendly_name"`
	AccountSid   string         `json:"account_sid"`
	URL          string         `json:"url"`
	Sid          string         `json:"sid"`
	DateCreated  time.Time      `json:"date_created"`
	Links        FieldTypeLinks `json:"links"`
}
