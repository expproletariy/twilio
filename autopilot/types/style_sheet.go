package types

import "encoding/json"

type StyleSheet struct {
	AccountSid   string      `json:"account_sid"`
	AssistantSid string      `json:"assistant_sid"`
	Data         interface{} `json:"data"`
	URL          string      `json:"url"`
}

func (s *StyleSheet) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, s)
}

func (s StyleSheet) MarshalJSON() ([]byte, error) {
	return json.Marshal(s)
}
