package taskfields

import (
	"encoding/json"
	"github.com/expproletariy/twilio/autopilot/tasks/taskfields/types"
	"github.com/expproletariy/twilio/common/errors"
	commontypes "github.com/expproletariy/twilio/common/types"
	"net/http"
	"net/url"
	"strings"
)

type taskFieldApiService struct {
	config commontypes.Config
}

func New(config commontypes.Config) TaskField {
	return newTaskFieldApiService(config)
}

func newTaskFieldApiService(config commontypes.Config) *taskFieldApiService {
	config.BaseApiUrl += "/Fields"
	return &taskFieldApiService{config: config}
}

func (t taskFieldApiService) Create(arguments types.TaskFieldCreateArguments) (types.TaskFieldResponse, errors.HttpError) {
	params := url.Values{}
	params.Set("UniqueName", arguments.UniqueName)
	params.Set("FieldType", arguments.FieldType)

	req, err := http.NewRequest("POST", t.config.BaseApiUrl, strings.NewReader(params.Encode()))
	if err != nil {
		return types.TaskFieldResponse{}, errors.NewHttpError(err)
	}
	req.SetBasicAuth(t.config.TwilioApiKeySid, t.config.TwilioApiKeySecret)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := t.config.Client.Do(req)
	if err != nil {
		return types.TaskFieldResponse{}, errors.NewHttpError(err)
	}
	if res.StatusCode != http.StatusCreated {
		return types.TaskFieldResponse{}, errors.NewHttpErrorNotCreatedWithResource(t.config.BaseApiUrl + ": " + params.Encode())
	}
	defer res.Body.Close()

	body := types.TaskFieldResponse{}
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return types.TaskFieldResponse{}, errors.NewHttpError(err)
	}

	return body, nil
}

func (t taskFieldApiService) GetBySid(taskFieldSid string) (types.TaskFieldResponse, errors.HttpError) {
	requestUrl := t.config.BaseApiUrl + "/" + taskFieldSid
	req, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return types.TaskFieldResponse{}, errors.NewHttpError(err)
	}
	req.SetBasicAuth(t.config.TwilioApiKeySid, t.config.TwilioApiKeySecret)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := t.config.Client.Do(req)
	if err != nil {
		return types.TaskFieldResponse{}, errors.NewHttpError(err)
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusUnauthorized {
		return types.TaskFieldResponse{}, errors.NewHttpErrorUnauthorized()
	}

	if res.StatusCode == http.StatusNotFound {
		return types.TaskFieldResponse{}, errors.NewHttpErrorNotFound()
	}

	body := types.TaskFieldResponse{}
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return types.TaskFieldResponse{}, errors.NewHttpError(err)
	}
	return body, nil
}

func (t taskFieldApiService) Get(meta types.Meta) (types.TaskFieldPaginationResponse, errors.HttpError) {
	if len(meta.NextPageURL) == 0 {
		meta.NextPageURL = t.config.BaseApiUrl
	}
	req, err := http.NewRequest("GET", meta.NextPageURL, nil)
	if err != nil {
		return types.TaskFieldPaginationResponse{}, errors.NewHttpError(err)
	}
	req.SetBasicAuth(t.config.TwilioApiKeySid, t.config.TwilioApiKeySecret)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := t.config.Client.Do(req)
	if err != nil {
		return types.TaskFieldPaginationResponse{}, errors.NewHttpError(err)
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusUnauthorized {
		return types.TaskFieldPaginationResponse{}, errors.NewHttpErrorUnauthorized()
	}

	if res.StatusCode == http.StatusNotFound {
		return types.TaskFieldPaginationResponse{}, errors.NewHttpErrorNotFound()
	}

	body := types.TaskFieldPaginationResponse{}
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return types.TaskFieldPaginationResponse{}, errors.NewHttpError(err)
	}
	return body, nil
}

func (t taskFieldApiService) Delete(taskFieldSid string) errors.HttpError {
	panic("implement me")
}
