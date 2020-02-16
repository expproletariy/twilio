package types

import "time"

type QueryCreateResult struct {
	Task           string  `json:"task"`
	TaskConfidence float32 `json:"task_confidence"`
	NextBestTask   string  `json:"next_best_task"`
	Fields         []Field `json:"fields"`
}

type queryCreateResponse struct {
	AssistantSid  string            `json:"assistant_sid"`
	SampleSid     interface{}       `json:"sample_sid"`
	DialogueSid   interface{}       `json:"dialogue_sid"`
	Language      string            `json:"language"`
	DateUpdated   time.Time         `json:"date_updated"`
	Sid           string            `json:"sid"`
	Results       QueryCreateResult `json:"results"`
	AccountSid    string            `json:"account_sid"`
	ModelBuildSid string            `json:"model_build_sid"`
	URL           string            `json:"url"`
	Status        string            `json:"status"`
	Query         string            `json:"query"`
	DateCreated   time.Time         `json:"date_created"`
	SourceChannel string            `json:"source_channel"`
}
