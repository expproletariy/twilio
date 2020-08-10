package types

import "encoding/json"

type AssistantLinks struct {
	FieldTypes  string `json:"field_types"`
	Tasks       string `json:"tasks"`
	ModelBuilds string `json:"model_builds"`
	Queries     string `json:"queries"`
	StyleSheet  string `json:"style_sheet"`
	Defaults    string `json:"defaults"`
	Dialogues   string `json:"dialogues"`
	Webhooks    string `json:"webhooks"`
}

func (l *AssistantLinks) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, l)
}

func (l AssistantLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(l)
}
