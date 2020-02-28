package fieldtypes

import (
	"github.com/expproletariy/twilio/autopilot/fieldtypes/fieldvalues"
	"github.com/expproletariy/twilio/autopilot/fieldtypes/types"
	"github.com/expproletariy/twilio/common/errors"
)

type FieldType interface {
	FieldValues(fieldTypeSid string) fieldvalues.FieldValue
	Create(arguments types.FiledTypeCreateArguments) (types.FieldTypeResponse, errors.HttpError)
	GetBySid(fieldTypeSid string) (types.FieldTypeResponse, errors.HttpError)
	Get(meta types.Meta) (types.FieldTypePaginationResponse, errors.HttpError)
	Update(arguments types.FieldTypeUpdateArguments) (types.FieldTypeResponse, errors.HttpError)
	Delete(fieldTypeSid string) errors.HttpError
}
