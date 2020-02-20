package types

type FieldTypePaginationResponse struct {
	FieldTypes FieldTypeResponse `json:"field_types"`
	Meta       Meta              `json:"meta"`
}
