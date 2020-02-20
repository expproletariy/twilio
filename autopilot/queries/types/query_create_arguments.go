package types

import (
	enums "github.com/expproletariy/twililo/autopilot/enums"
)

type QueryCreateArguments struct {
	Query    string         `json:"queries"`
	Language enums.Language `json:"language"`
}
