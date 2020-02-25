package types

type TaskCreateArguments struct {
	UniqueName   string `json:"unique_name"` // required
	FriendlyName string `json:"friendly_name"`
	Actions      string `json:"actions"` // The JSON string that specifies the actions
	ActionsUrl   string `json:"actions_url"`
}
