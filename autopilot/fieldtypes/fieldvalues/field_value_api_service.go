package fieldvalues

import (
	"github.com/expproletariy/twilio/autopilot/fieldtypes/fieldvalues/types"
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

func (f fieldValueApiService) Create(arguments types.FieldValueCreateArguments) (types.FieldValueResponse, error) {
	panic("implement me")
}

func (f fieldValueApiService) GetBySid(fieldValueSid string) (types.FieldValueResponse, error) {
	panic("implement me")
}

func (f fieldValueApiService) Get(meta types.Meta) (types.FieldValuePaginationResponse, error) {
	panic("implement me")
}

func (f fieldValueApiService) Delete(fieldValueSid string) error {
	panic("implement me")
}
