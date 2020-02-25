package types

type TaskPaginationResponse struct {
	Meta  Meta           `json:"meta"`
	Tasks []TaskResponse `json:"tasks"`
}
