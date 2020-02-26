package types

import (
	commonenums "github.com/expproletariy/twilio/common/enums"
	"time"
)

type FieldValueResponse struct {
	AssistantSid string               `json:"assistant_sid"`
	Language     commonenums.Language `json:"language"`
	DateUpdated  time.Time            `json:"date_updated"`
	SynonymOf    string               `json:"synonym_of"`
	Value        string               `json:"value"`
	AccountSid   string               `json:"account_sid"`
	URL          string               `json:"url"`
	Sid          string               `json:"sid"`
	DateCreated  time.Time            `json:"date_created"`
	FieldTypeSid string               `json:"field_type_sid"`
}
