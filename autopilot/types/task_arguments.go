package types

import "encoding/json"

type TaskArguments struct {
	UniqueName   string      `json:"unique_name"` // required
	FriendlyName string      `json:"friendly_name"`
	Actions      interface{} `json:"actions"` // The JSON string that specifies the actions
	ActionsUrl   string      `json:"actions_url"`
}

func (t *TaskArguments) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, t)
}

func (t TaskArguments) MarshalJSON() ([]byte, error) {
	return json.Marshal(t)
}
