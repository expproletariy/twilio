package fieldvalues

import (
	"github.com/expproletariy/twilio/autopilot/fieldtypes/fieldvalues/types"
	"github.com/expproletariy/twilio/common/errors"
	commontypes "github.com/expproletariy/twilio/common/types"
)

type fieldValueApiService struct {
	config commontypes.Config
}

func New(config commontypes.Config) FieldValue {
	return newFieldValueApiService(config)
}

func newFieldValueApiService(config commontypes.Config) FieldValue {
	config.BaseApiUrl = "/FieldValues"
	return &fieldValueApiService{config: config}
}

func (f fieldValueApiService) Create(arguments types.FieldValueCreateArguments) (types.FieldValueResponse, errors.HttpError) {
	panic("implement me")
}

func (f fieldValueApiService) GetBySid(fieldValueSid string) (types.FieldValueResponse, errors.HttpError) {
	panic("implement me")
}

func (f fieldValueApiService) Get(meta types.Meta) (types.FieldValuePaginationResponse, errors.HttpError) {
	panic("implement me")
}

func (f fieldValueApiService) Delete(fieldValueSid string) errors.HttpError {
	panic("implement me")
}
