package fieldtypes

import (
	"github.com/expproletariy/twilio/autopilot/fieldtypes/fieldvalues"
	"github.com/expproletariy/twilio/autopilot/fieldtypes/types"
)

type FieldType interface {
	FieldValues(fieldTypeSid string) fieldvalues.FieldValue
	Create(arguments types.FiledTypeCreateArguments) (types.FieldTypeResponse, error)
	GetBySid(fieldTypeSid string) (types.FieldTypeResponse, error)
	Get(meta types.Meta) (types.FieldTypePaginationResponse, error)
	Update(arguments types.FieldTypeUpdateArguments) (types.FieldTypeResponse, error)
	Delete(fieldTypeSid string) error
}
