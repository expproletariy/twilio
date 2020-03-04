package tasks

import (
	"encoding/json"
	"github.com/expproletariy/twilio/autopilot/tasks/taskfields"
	"github.com/expproletariy/twilio/autopilot/tasks/tasksamples"
	"github.com/expproletariy/twilio/autopilot/tasks/types"
	"github.com/expproletariy/twilio/common/errors"
	commontypes "github.com/expproletariy/twilio/common/types"
	"net/http"
	"net/url"
	"strings"
)

type taskApiService struct {
	config commontypes.Config
}

func New(config commontypes.Config) Task {
	return newTaskApiService(config)
}

func newTaskApiService(config commontypes.Config) Task {
	config.BaseApiUrl += "/Tasks"
	return &taskApiService{config: config}
}

func (t taskApiService) TaskSamples(taskSid string) tasksamples.TaskSample {
	config := commontypes.Config{
		Client:             t.config.Client,
		BaseApiUrl:         t.config.BaseApiUrl + "/" + taskSid,
		ApiVersion:         t.config.ApiVersion,
		TwilioApiKeySid:    t.config.TwilioApiKeySid,
		TwilioApiKeySecret: t.config.TwilioApiKeySecret,
	}
	return tasksamples.New(config)
}

func (t taskApiService) TaskFields(taskSid string) taskfields.TaskField {
	config := commontypes.Config{
		Client:             t.config.Client,
		BaseApiUrl:         t.config.BaseApiUrl + "/" + taskSid,
		ApiVersion:         t.config.ApiVersion,
		TwilioApiKeySid:    t.config.TwilioApiKeySid,
		TwilioApiKeySecret: t.config.TwilioApiKeySecret,
	}
	return taskfields.New(config)
}

func (t taskApiService) Create(arguments types.TaskCreateArguments) (types.TaskResponse, errors.HttpError) {
	params := url.Values{}
	params.Set("UniqueName", arguments.UniqueName)
	params.Set("FriendlyName", arguments.FriendlyName)

	req, err := http.NewRequest("POST", t.config.BaseApiUrl, strings.NewReader(params.Encode()))
	if err != nil {
		return types.TaskResponse{}, errors.NewHttpError(err)
	}
	req.SetBasicAuth(t.config.TwilioApiKeySid, t.config.TwilioApiKeySecret)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := t.config.Client.Do(req)
	if err != nil {
		return types.TaskResponse{}, errors.NewHttpError(err)
	}
	if res.StatusCode != http.StatusCreated {
		return types.TaskResponse{}, errors.NewHttpErrorNotCreated()
	}
	defer res.Body.Close()

	body := types.TaskResponse{}
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return types.TaskResponse{}, errors.NewHttpError(err)
	}

	return body, nil
}

func (t taskApiService) GetBySid(taskSid string) (types.TaskResponse, errors.HttpError) {
	requestUrl := t.config.BaseApiUrl + "/" + taskSid
	req, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return types.TaskResponse{}, errors.NewHttpError(err)
	}
	req.SetBasicAuth(t.config.TwilioApiKeySid, t.config.TwilioApiKeySecret)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := t.config.Client.Do(req)
	if err != nil {
		return types.TaskResponse{}, errors.NewHttpError(err)
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusUnauthorized {
		return types.TaskResponse{}, errors.NewHttpErrorUnauthorized()
	}

	if res.StatusCode == http.StatusNotFound {
		return types.TaskResponse{}, errors.NewHttpErrorNotFound()
	}

	body := types.TaskResponse{}
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return types.TaskResponse{}, errors.NewHttpError(err)
	}
	return body, nil
}

func (t taskApiService) Get(meta types.Meta) (types.TaskPaginationResponse, errors.HttpError) {
	if len(meta.NextPageURL) == 0 {
		meta.NextPageURL = t.config.BaseApiUrl
	}
	req, err := http.NewRequest("GET", meta.NextPageURL, nil)
	if err != nil {
		return types.TaskPaginationResponse{}, errors.NewHttpError(err)
	}
	req.SetBasicAuth(t.config.TwilioApiKeySid, t.config.TwilioApiKeySecret)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := t.config.Client.Do(req)
	if err != nil {
		return types.TaskPaginationResponse{}, errors.NewHttpError(err)
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusUnauthorized {
		return types.TaskPaginationResponse{}, errors.NewHttpErrorUnauthorized()
	}

	if res.StatusCode == http.StatusNotFound {
		return types.TaskPaginationResponse{}, errors.NewHttpErrorNotFound()
	}

	body := types.TaskPaginationResponse{}
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return types.TaskPaginationResponse{}, errors.NewHttpError(err)
	}
	return body, nil
}

func (t taskApiService) Update(arguments types.TaskUpdateArguments) (types.TaskResponse, errors.HttpError) {
	panic("implement me")
}

func (t taskApiService) Delete(taskSid string) errors.HttpError {
	panic("implement me")
}
