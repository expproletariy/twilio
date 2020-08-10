package types

import (
	"encoding/json"
	"time"
)

type Assistant struct {
	AccountSid          string         `json:"account_sid"`
	DateCreated         time.Time      `json:"date_created"`
	DateUpdated         time.Time      `json:"date_updated"`
	FriendlyName        string         `json:"friendly_name"`
	LatestModelBuildSid string         `json:"latest_model_build_sid"`
	LogQueries          bool           `json:"log_queries"`
	DevelopmentStage    string         `json:"development_stage"`
	NeedsModelBuild     bool           `json:"needs_model_build"`
	Sid                 string         `json:"sid"`
	UniqueName          string         `json:"unique_name"`
	Links               AssistantLinks `json:"links"`
	URL                 string         `json:"url"`
	CallbackURL         string         `json:"callback_url"`
	CallbackEvents      string         `json:"callback_events"`
}

func (a *Assistant) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, a)
}

func (a Assistant) MarshalJSON() ([]byte, error) {
	return json.Marshal(a)
}
