package types

import (
	enums "github.com/expproletariy/twilio/common/enums"
)

type QueryCreateArguments struct {
	Query    string         `json:"queries"`
	Language enums.Language `json:"language"`
}
