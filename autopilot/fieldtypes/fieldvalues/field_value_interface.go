package fieldvalues

import "github.com/expproletariy/twililo/autopilot/fieldtypes/fieldvalues/types"

type FieldValue interface {
	Create(arguments types.FieldValueCreateArguments) (types.FieldValueResponse, error)
	GetBySid(fieldValueSid string) (types.FieldValueResponse, error)
	Get(meta types.Meta) (types.FieldValuePaginationResponse, error)
	Delete(fieldValueSid string) error
}
