package types

type QueryPaginationResponse struct {
	Meta    Meta            `json:"meta"`
	Queries []QueryResponse `json:"queries"`
}
