package types

import "github.com/expproletariy/twililo/autopilot/enums"

type FieldValueCreateArguments struct {
	Language enums.Language `json:"language"`
	Value    string         `json:"value"`
}
