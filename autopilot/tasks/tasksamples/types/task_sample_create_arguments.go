package types

import "github.com/expproletariy/twililo/common/enums"

type TaskSampleCreateArguments struct {
	Language   enums.Language `json:"language"`
	TaggedText string         `json:"tagged_text"`
}
