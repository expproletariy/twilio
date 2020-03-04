package fieldvalues

import (
	"encoding/json"
	"github.com/expproletariy/twilio/autopilot/fieldtypes/fieldvalues/types"
	"github.com/expproletariy/twilio/common/errors"
	commontypes "github.com/expproletariy/twilio/common/types"
	"net/http"
	"net/url"
	"strings"
)

type fieldValueApiService struct {
	config commontypes.Config
}

func New(config commontypes.Config) FieldValue {
	return newFieldValueApiService(config)
}

func newFieldValueApiService(config commontypes.Config) FieldValue {
	config.BaseApiUrl += "/FieldValues"
	return &fieldValueApiService{config: config}
}

func (f fieldValueApiService) Create(arguments types.FieldValueCreateArguments) (types.FieldValueResponse, errors.HttpError) {
	params := url.Values{}
	params.Set("Value", arguments.Value)
	params.Set("Language", arguments.Language.String())

	req, err := http.NewRequest("POST", f.config.BaseApiUrl, strings.NewReader(params.Encode()))
	if err != nil {
		return types.FieldValueResponse{}, errors.NewHttpError(err)
	}
	req.SetBasicAuth(f.config.TwilioApiKeySid, f.config.TwilioApiKeySecret)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := f.config.Client.Do(req)
	if err != nil {
		return types.FieldValueResponse{}, errors.NewHttpError(err)
	}
	if res.StatusCode != http.StatusCreated {
		return types.FieldValueResponse{}, errors.NewHttpErrorNotCreatedWithResource(f.config.BaseApiUrl + ": " + params.Encode())
	}
	defer res.Body.Close()

	body := types.FieldValueResponse{}
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return types.FieldValueResponse{}, errors.NewHttpError(err)
	}

	return body, nil
}

func (f fieldValueApiService) GetBySid(fieldValueSid string) (types.FieldValueResponse, errors.HttpError) {
	requestUrl := f.config.BaseApiUrl + "/" + fieldValueSid
	req, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return types.FieldValueResponse{}, errors.NewHttpError(err)
	}
	req.SetBasicAuth(f.config.TwilioApiKeySid, f.config.TwilioApiKeySecret)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := f.config.Client.Do(req)
	if err != nil {
		return types.FieldValueResponse{}, errors.NewHttpError(err)
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusUnauthorized {
		return types.FieldValueResponse{}, errors.NewHttpErrorUnauthorized()
	}

	if res.StatusCode == http.StatusNotFound {
		return types.FieldValueResponse{}, errors.NewHttpErrorNotFound()
	}

	body := types.FieldValueResponse{}
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return types.FieldValueResponse{}, errors.NewHttpError(err)
	}
	return body, nil
}

func (f fieldValueApiService) Get(meta types.Meta) (types.FieldValuePaginationResponse, errors.HttpError) {
	if len(meta.NextPageURL) == 0 {
		meta.NextPageURL = f.config.BaseApiUrl
	}
	req, err := http.NewRequest("GET", meta.NextPageURL, nil)
	if err != nil {
		return types.FieldValuePaginationResponse{}, errors.NewHttpError(err)
	}
	req.SetBasicAuth(f.config.TwilioApiKeySid, f.config.TwilioApiKeySecret)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := f.config.Client.Do(req)
	if err != nil {
		return types.FieldValuePaginationResponse{}, errors.NewHttpError(err)
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusUnauthorized {
		return types.FieldValuePaginationResponse{}, errors.NewHttpErrorUnauthorized()
	}

	if res.StatusCode == http.StatusNotFound {
		return types.FieldValuePaginationResponse{}, errors.NewHttpErrorNotFound()
	}

	body := types.FieldValuePaginationResponse{}
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return types.FieldValuePaginationResponse{}, errors.NewHttpError(err)
	}
	return body, nil
}

func (f fieldValueApiService) Delete(fieldValueSid string) errors.HttpError {
	panic("implement me")
}
