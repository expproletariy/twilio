package types

import (
	"github.com/expproletariy/twililo/autopilot/queries/enums"
	"time"
)

type queryResponse struct {
	AssistantSid  string            `json:"assistant_sid"`
	SampleSid     string            `json:"sample_sid"`
	DialogueSid   string            `json:"dialogue_sid"`
	Language      string            `json:"language"`
	DateUpdated   time.Time         `json:"date_updated"`
	Sid           string            `json:"sid"`
	Results       QueryResult       `json:"results"`
	AccountSid    string            `json:"account_sid"`
	ModelBuildSid string            `json:"model_build_sid"`
	URL           string            `json:"url"`
	Status        enums.QueryStatus `json:"status"`
	Query         string            `json:"queries"`
	DateCreated   time.Time         `json:"date_created"`
	SourceChannel string            `json:"source_channel"`
}
