package types

type Meta struct {
	FirstPageURL    string `json:"first_page_url"`
	PreviousPageURL string `json:"previous_page_url"`
	Key             string `json:"key"`
	NextPageURL     string `json:"next_page_url"`
	URL             string `json:"url"`
	Page            int    `json:"page"`
	PageSize        int    `json:"page_size"`
}
