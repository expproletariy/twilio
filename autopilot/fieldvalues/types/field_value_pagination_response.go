package types

type FieldValuePaginationResponse struct {
	FieldValues []FieldValueResponse `json:"field_values"`
	Meta        Meta                 `json:"meta"`
}
