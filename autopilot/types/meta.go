package types

import "encoding/json"

type Meta struct {
	FirstPageURL    string `json:"first_page_url"`
	Key             string `json:"key"`
	NextPageURL     string `json:"next_page_url"`
	Page            int64  `json:"page"`
	PageSize        int64  `json:"page_size"`
	PreviousPageURL string `json:"previous_page_url"`
	URL             string `json:"url"`
}

func (m *Meta) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, m)
}

func (m Meta) MarshalJSON() ([]byte, error) {
	return json.Marshal(m)
}
