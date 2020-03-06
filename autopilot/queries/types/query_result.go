package types

type QueryResult struct {
	Task           string  `json:"task"`
	TaskConfidence string  `json:"task_confidence"`
	NextBestTask   string  `json:"next_best_task"`
	Fields         []Field `json:"fields"`
}
