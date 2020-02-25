package types

type TaskUpdateArguments struct {
	Sid          string `json:"sid"`
	UniqueName   string `json:"unique_name"`
	FriendlyName string `json:"friendly_name"`
	Actions      string `json:"actions"` // The JSON string that specifies the actions
	ActionsUrl   string `json:"actions_url"`
}
