package types

type Meta struct {
	Key             string `json:"key"`
	PageSize        int    `json:"page_size"`
	NextPageURL     string `json:"next_page_url"`
	URL             string `json:"url"`
	Page            int    `json:"page"`
	FirstPageURL    string `json:"first_page_url"`
	PreviousPageURL string `json:"previous_page_url"`
}
