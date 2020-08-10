package types

import (
	"encoding/json"
	"time"
)

type Task struct {
	UniqueName   string    `json:"unique_name"`
	Links        TaskLinks `json:"links"`
	AccountSid   string    `json:"account_sid"`
	FriendlyName string    `json:"friendly_name"`
	URL          string    `json:"url"`
	Sid          string    `json:"sid"`
	DateUpdated  time.Time `json:"date_updated"`
	AssistantSid string    `json:"assistant_sid"`
	DateCreated  time.Time `json:"date_created"`
	ActionsURL   string    `json:"actions_url"`
}

func (t *Task) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, t)
}

func (t Task) MarshalJSON() ([]byte, error) {
	return json.Marshal(t)
}
