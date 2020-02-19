package types

type queryPaginationResponse struct {
	Meta    Meta          `json:"meta"`
	Queries []QueryResult `json:"queries"`
}
