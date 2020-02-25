package types

type TaskFieldPaginationResponse struct {
	Fields []TaskFieldResponse `json:"fields"`
	Meta   Meta                `json:"meta"`
}
