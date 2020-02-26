package types

import commonenums "github.com/expproletariy/twilio/common/enums"

type FieldValueCreateArguments struct {
	Language commonenums.Language `json:"language"`
	Value    string               `json:"value"`
}
