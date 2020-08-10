package types

import "github.com/expproletariy/twilio/autopilot/queries/enums"

type QueryUpdateArguments struct {
	SampleSid string            `json:"sample_sid"`
	Status    enums.QueryStatus `json:"status"`
	Sid       string            `json:"sid"`
}
