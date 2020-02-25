package types

import commonenums "github.com/expproletariy/twililo/common/enums"

type FieldValueCreateArguments struct {
	Language commonenums.Language `json:"language"`
	Value    string               `json:"value"`
}
