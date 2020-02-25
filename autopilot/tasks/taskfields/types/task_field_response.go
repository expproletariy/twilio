package types

import "time"

type TaskFieldResponse struct {
	AssistantSid string    `json:"assistant_sid"`
	FieldType    string    `json:"field_type"`
	UniqueName   string    `json:"unique_name"`
	TaskSid      string    `json:"task_sid"`
	DateUpdated  time.Time `json:"date_updated"`
	AccountSid   string    `json:"account_sid"`
	URL          string    `json:"url"`
	Sid          string    `json:"sid"`
	DateCreated  time.Time `json:"date_created"`
}
