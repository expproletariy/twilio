package types

import autopilotEnums "github.com/expproletariy/twililo/common/enums"

type TaskSampleUpdateArguments struct {
	Language   autopilotEnums.Language `json:"language"`
	TaggedText string                  `json:"tagged_text"`
}
