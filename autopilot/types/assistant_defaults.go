package types

import "encoding/json"

type AssistantDefaults struct {
	AccountSid   string      `json:"account_sid"`
	AssistantSid string      `json:"assistant_sid"`
	Data         interface{} `json:"data"`
	URL          string      `json:"url"`
}

func (ad *AssistantDefaults) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, ad)
}

func (ad AssistantDefaults) MarshalJSON() ([]byte, error) {
	return json.Marshal(ad)
}
