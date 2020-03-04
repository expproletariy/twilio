package tasksamples

import (
	"encoding/json"
	"github.com/expproletariy/twilio/autopilot/tasks/tasksamples/types"
	"github.com/expproletariy/twilio/common/errors"
	commontypes "github.com/expproletariy/twilio/common/types"
	"net/http"
	"net/url"
	"strings"
)

type taskSampleApiService struct {
	config commontypes.Config
}

func New(config commontypes.Config) TaskSample {
	return newTaskSampleApiService(config)
}

func newTaskSampleApiService(config commontypes.Config) TaskSample {
	config.BaseApiUrl += "/Samples"
	return &taskSampleApiService{config: config}
}

func (t taskSampleApiService) Create(arguments types.TaskSampleCreateArguments) (types.TaskSampleResponse, errors.HttpError) {
	params := url.Values{}
	params.Set("TaggedText", arguments.TaggedText)
	params.Set("Language", arguments.Language.String())

	req, err := http.NewRequest("POST", t.config.BaseApiUrl, strings.NewReader(params.Encode()))
	if err != nil {
		return types.TaskSampleResponse{}, errors.NewHttpError(err)
	}
	req.SetBasicAuth(t.config.TwilioApiKeySid, t.config.TwilioApiKeySecret)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := t.config.Client.Do(req)
	if err != nil {
		return types.TaskSampleResponse{}, errors.NewHttpError(err)
	}
	if res.StatusCode != http.StatusCreated {
		return types.TaskSampleResponse{}, errors.NewHttpErrorNotCreated()
	}
	defer res.Body.Close()

	body := types.TaskSampleResponse{}
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return types.TaskSampleResponse{}, errors.NewHttpError(err)
	}

	return body, nil
}

func (t taskSampleApiService) Get(meta types.Meta) (types.TaskSamplePaginationResponse, errors.HttpError) {
	if len(meta.NextPageURL) == 0 {
		meta.NextPageURL = t.config.BaseApiUrl
	}
	req, err := http.NewRequest("GET", meta.NextPageURL, nil)
	if err != nil {
		return types.TaskSamplePaginationResponse{}, errors.NewHttpError(err)
	}
	req.SetBasicAuth(t.config.TwilioApiKeySid, t.config.TwilioApiKeySecret)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := t.config.Client.Do(req)
	if err != nil {
		return types.TaskSamplePaginationResponse{}, errors.NewHttpError(err)
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusUnauthorized {
		return types.TaskSamplePaginationResponse{}, errors.NewHttpErrorUnauthorized()
	}

	if res.StatusCode == http.StatusNotFound {
		return types.TaskSamplePaginationResponse{}, errors.NewHttpErrorNotFound()
	}

	body := types.TaskSamplePaginationResponse{}
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return types.TaskSamplePaginationResponse{}, errors.NewHttpError(err)
	}
	return body, nil
}

func (t taskSampleApiService) GetBySid(taskSampleSid string) (types.TaskSampleResponse, errors.HttpError) {
	requestUrl := t.config.BaseApiUrl + "/" + taskSampleSid
	req, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return types.TaskSampleResponse{}, errors.NewHttpError(err)
	}
	req.SetBasicAuth(t.config.TwilioApiKeySid, t.config.TwilioApiKeySecret)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := t.config.Client.Do(req)
	if err != nil {
		return types.TaskSampleResponse{}, errors.NewHttpError(err)
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusUnauthorized {
		return types.TaskSampleResponse{}, errors.NewHttpErrorUnauthorized()
	}

	if res.StatusCode == http.StatusNotFound {
		return types.TaskSampleResponse{}, errors.NewHttpErrorNotFound()
	}

	body := types.TaskSampleResponse{}
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return types.TaskSampleResponse{}, errors.NewHttpError(err)
	}
	return body, nil
}

func (t taskSampleApiService) Update(arguments types.TaskSampleUpdateArguments) (types.TaskSampleResponse, errors.HttpError) {
	panic("implement me")
}

func (t taskSampleApiService) Delete(taskSampleSid string) errors.HttpError {
	panic("implement me")
}
