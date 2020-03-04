package fieldtypes

import (
	"encoding/json"
	"github.com/expproletariy/twilio/autopilot/fieldtypes/enums"
	"github.com/expproletariy/twilio/autopilot/fieldtypes/fieldvalues"
	"github.com/expproletariy/twilio/autopilot/fieldtypes/types"
	"github.com/expproletariy/twilio/common/errors"
	commontypes "github.com/expproletariy/twilio/common/types"
	"net/http"
	"net/url"
	"strings"
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

func (f fieldTypeApiService) Create(arguments types.FiledTypeCreateArguments) (types.FieldTypeResponse, errors.HttpError) {

	if enums.IsBuildInFieldType(arguments.UniqueName) {
		return types.FieldTypeResponse{}, errors.NewHttpErrorBuildInType()
	}

	params := url.Values{}
	params.Set("UniqueName", arguments.UniqueName)

	req, err := http.NewRequest("POST", f.config.BaseApiUrl, strings.NewReader(params.Encode()))
	if err != nil {
		return types.FieldTypeResponse{}, errors.NewHttpError(err)
	}
	req.SetBasicAuth(f.config.TwilioApiKeySid, f.config.TwilioApiKeySecret)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := f.config.Client.Do(req)
	if err != nil {
		return types.FieldTypeResponse{}, errors.NewHttpError(err)
	}
	if res.StatusCode != http.StatusCreated {
		return types.FieldTypeResponse{}, errors.NewHttpErrorNotCreatedWithResource(f.config.BaseApiUrl + ": " + params.Encode())
	}
	defer res.Body.Close()

	body := types.FieldTypeResponse{}
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return types.FieldTypeResponse{}, errors.NewHttpError(err)
	}

	return body, nil
}

func (f fieldTypeApiService) GetBySid(fieldTypeSid string) (types.FieldTypeResponse, errors.HttpError) {
	requestUrl := f.config.BaseApiUrl + "/" + fieldTypeSid
	req, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return types.FieldTypeResponse{}, errors.NewHttpError(err)
	}
	req.SetBasicAuth(f.config.TwilioApiKeySid, f.config.TwilioApiKeySecret)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := f.config.Client.Do(req)
	if err != nil {
		return types.FieldTypeResponse{}, errors.NewHttpError(err)
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusUnauthorized {
		return types.FieldTypeResponse{}, errors.NewHttpErrorUnauthorized()
	}

	if res.StatusCode == http.StatusNotFound {
		return types.FieldTypeResponse{}, errors.NewHttpErrorNotFound()
	}

	body := types.FieldTypeResponse{}
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return types.FieldTypeResponse{}, errors.NewHttpError(err)
	}
	return body, nil
}

func (f fieldTypeApiService) Get(meta types.Meta) (types.FieldTypePaginationResponse, errors.HttpError) {
	if len(meta.NextPageURL) == 0 {
		meta.NextPageURL = f.config.BaseApiUrl
	}
	req, err := http.NewRequest("GET", meta.NextPageURL, nil)
	if err != nil {
		return types.FieldTypePaginationResponse{}, errors.NewHttpError(err)
	}
	req.SetBasicAuth(f.config.TwilioApiKeySid, f.config.TwilioApiKeySecret)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := f.config.Client.Do(req)
	if err != nil {
		return types.FieldTypePaginationResponse{}, errors.NewHttpError(err)
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusUnauthorized {
		return types.FieldTypePaginationResponse{}, errors.NewHttpErrorUnauthorized()
	}

	if res.StatusCode == http.StatusNotFound {
		return types.FieldTypePaginationResponse{}, errors.NewHttpErrorNotFound()
	}

	body := types.FieldTypePaginationResponse{}
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return types.FieldTypePaginationResponse{}, errors.NewHttpError(err)
	}
	return body, nil
}

func (f fieldTypeApiService) Update(arguments types.FieldTypeUpdateArguments) (types.FieldTypeResponse, errors.HttpError) {
	panic("implement me")
}

func (f fieldTypeApiService) Delete(fieldTypeSid string) errors.HttpError {
	panic("implement me")
}
