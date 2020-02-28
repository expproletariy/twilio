package fieldvalues

import (
	"github.com/expproletariy/twilio/autopilot/fieldtypes/fieldvalues/types"
	"github.com/expproletariy/twilio/common/errors"
)

type FieldValue interface {
	Create(arguments types.FieldValueCreateArguments) (types.FieldValueResponse, errors.HttpError)
	GetBySid(fieldValueSid string) (types.FieldValueResponse, errors.HttpError)
	Get(meta types.Meta) (types.FieldValuePaginationResponse, errors.HttpError)
	Delete(fieldValueSid string) errors.HttpError
}
