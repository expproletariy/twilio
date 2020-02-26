package types

import autopilotEnums "github.com/expproletariy/twilio/common/enums"

type TaskSampleUpdateArguments struct {
	Language   autopilotEnums.Language `json:"language"`
	TaggedText string                  `json:"tagged_text"`
}
