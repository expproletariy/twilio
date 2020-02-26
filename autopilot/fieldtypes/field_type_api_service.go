package fieldtypes

import (
	"github.com/expproletariy/twilio/autopilot/fieldtypes/fieldvalues"
	"github.com/expproletariy/twilio/autopilot/fieldtypes/types"
	commontypes "github.com/expproletariy/twilio/common/types"
)

type fieldTypeApiService struct {
	config commontypes.Config
}

func New(config commontypes.Config) FieldType {
	return newFieldTypeApiService(config)
}

func newFieldTypeApiService(config commontypes.Config) FieldType {
	config.BaseApiUrl += "/FieldTypes"
	return &fieldTypeApiService{config: config}
}

func (f fieldTypeApiService) FieldValues(fieldTypeSid string) fieldvalues.FieldValue {
	config := commontypes.Config{
		Client:             f.config.Client,
		BaseApiUrl:         f.config.BaseApiUrl + "/" + fieldTypeSid,
		ApiVersion:         f.config.ApiVersion,
		TwilioApiKeySid:    f.config.TwilioApiKeySid,
		TwilioApiKeySecret: f.config.TwilioApiKeySecret,
	}
	return fieldvalues.New(config)
}

func (f fieldTypeApiService) Create(arguments types.FiledTypeCreateArguments) (types.FieldTypeResponse, error) {
	panic("implement me")
}

func (f fieldTypeApiService) GetBySid(fieldTypeSid string) (types.FieldTypeResponse, error) {
	panic("implement me")
}

func (f fieldTypeApiService) Get(meta types.Meta) (types.FieldTypePaginationResponse, error) {
	panic("implement me")
}

func (f fieldTypeApiService) Update(arguments types.FieldTypeUpdateArguments) (types.FieldTypeResponse, error) {
	panic("implement me")
}

func (f fieldTypeApiService) Delete(fieldTypeSid string) error {
	panic("implement me")
}
