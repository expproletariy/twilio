package types

import "github.com/expproletariy/twililo/autopilot/queries/enums"

type QueryCreateArguments struct {
	Query    string              `json:"queries"`
	Language enums.QueryLanguage `json:"language"`
}
