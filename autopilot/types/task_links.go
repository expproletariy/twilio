package types

import "encoding/json"

type TaskLinks struct {
	Fields      string `json:"fields"`
	Samples     string `json:"samples"`
	TaskActions string `json:"task_actions"`
	Statistics  string `json:"statistics"`
}

func (l *TaskLinks) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, l)
}

func (l TaskLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(l)
}
