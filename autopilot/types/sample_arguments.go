package types

import "github.com/expproletariy/twilio/common/enums"

type SampleArguments struct {
	Language      enums.Language `json:"language"`
	TaggedText    string         `json:"tagged_text"`
	SourceChannel string         `json:"source_channel"`
}
