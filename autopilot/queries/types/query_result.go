package types

import "github.com/expproletariy/twililo/autopilot/queries/enums"

type QueryResult struct {
	Task           string            `json:"task"`
	TaskConfidence float32           `json:"task_confidence"`
	NextBestTask   string            `json:"next_best_task"`
	Fields         []Field           `json:"fieldtypes"`
	Status         enums.QueryStatus `json:"status"`
}
