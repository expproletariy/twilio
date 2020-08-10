package types

import (
	"encoding/json"
	"github.com/expproletariy/twilio/autopilot/enums"
	commonEnums "github.com/expproletariy/twilio/common/enums"
	"time"
)

type Sample struct {
	URL           string               `json:"url"`
	TaskSid       string               `json:"task_sid"`
	Sid           string               `json:"sid"`
	AssistantSid  string               `json:"assistant_sid"`
	AccountSid    string               `json:"account_sid"`
	DateCreated   time.Time            `json:"date_created"`
	Language      commonEnums.Language `json:"language"`
	TaggedText    string               `json:"tagged_text"`
	DateUpdated   time.Time            `json:"date_updated"`
	SourceChannel enums.SourceChannel  `json:"source_channel"`
}

func (s *Sample) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, s)
}

func (s Sample) MarshalJSON() ([]byte, error) {
	return json.Marshal(s)
}
