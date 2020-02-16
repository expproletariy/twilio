package types

import "github.com/expproletariy/twililo/autopilot/enums"

type QueryCreateArguments struct {
	Query    string              `json:"query"`
	Language enums.QueryLanguage `json:"language"`
}
