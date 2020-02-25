package types

type TaskSamplePaginationResponse struct {
	Samples []TaskSampleResponse `json:"samples"`
	Meta    Meta                 `json:"meta"`
}
