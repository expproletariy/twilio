package types

import (
	"github.com/expproletariy/twililo/common/enums"
	"time"
)

type TaskSampleResponse struct {
	URL           string         `json:"url"`
	TaskSid       string         `json:"task_sid"`
	Sid           string         `json:"sid"`
	AssistantSid  string         `json:"assistant_sid"`
	AccountSid    string         `json:"account_sid"`
	DateCreated   time.Time      `json:"date_created"`
	Language      enums.Language `json:"language"`
	TaggedText    string         `json:"tagged_text"`
	DateUpdated   time.Time      `json:"date_updated"`
	SourceChannel string         `json:"source_channel"`
}
