package queries

import (
	"encoding/json"
	"github.com/expproletariy/twilio/autopilot/queries/types"
	"github.com/expproletariy/twilio/common/errors"
	commontypes "github.com/expproletariy/twilio/common/types"
	"net/http"
	"net/url"
	"strings"
)

type queryApiService struct {
	config commontypes.Config
}

func New(config commontypes.Config) Query {
	return newQueryApiService(config)
}

func newQueryApiService(config commontypes.Config) Query {
	config.BaseApiUrl += "/Queries"
	return &queryApiService{config: config}
}

func (q queryApiService) Create(arguments types.QueryCreateArguments) (types.QueryResponse, errors.HttpError) {
	params := url.Values{}
	params.Set("Query", arguments.Query)
	params.Set("Language", arguments.Language.String())

	req, err := http.NewRequest("POST", q.config.BaseApiUrl, strings.NewReader(params.Encode()))
	if err != nil {
		return types.QueryResponse{}, errors.NewHttpError(err)
	}
	req.SetBasicAuth(q.config.TwilioApiKeySid, q.config.TwilioApiKeySecret)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := q.config.Client.Do(req)
	if err != nil {
		return types.QueryResponse{}, errors.NewHttpError(err)
	}
	if res.StatusCode != http.StatusCreated {
		return types.QueryResponse{}, errors.NewHttpErrorNotCreatedWithResource(q.config.BaseApiUrl + ": " + params.Encode())
	}
	defer res.Body.Close()

	body := types.QueryResponse{}
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return types.QueryResponse{}, errors.NewHttpError(err)
	}

	return body, nil
}

func (q queryApiService) GetBySid(querySid string) (types.QueryResponse, errors.HttpError) {
	requestUrl := q.config.BaseApiUrl + "/" + querySid
	req, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return types.QueryResponse{}, errors.NewHttpError(err)
	}
	req.SetBasicAuth(q.config.TwilioApiKeySid, q.config.TwilioApiKeySecret)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := q.config.Client.Do(req)
	if err != nil {
		return types.QueryResponse{}, errors.NewHttpError(err)
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusUnauthorized {
		return types.QueryResponse{}, errors.NewHttpErrorUnauthorized()
	}

	if res.StatusCode == http.StatusNotFound {
		return types.QueryResponse{}, errors.NewHttpErrorNotFound()
	}

	body := types.QueryResponse{}
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return types.QueryResponse{}, errors.NewHttpError(err)
	}
	return body, nil
}

func (q queryApiService) Get(meta types.Meta) (types.QueryPaginationResponse, errors.HttpError) {
	if len(meta.NextPageURL) == 0 {
		meta.NextPageURL = q.config.BaseApiUrl
	}
	req, err := http.NewRequest("GET", meta.NextPageURL, nil)
	if err != nil {
		return types.QueryPaginationResponse{}, errors.NewHttpError(err)
	}
	req.SetBasicAuth(q.config.TwilioApiKeySid, q.config.TwilioApiKeySecret)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := q.config.Client.Do(req)
	if err != nil {
		return types.QueryPaginationResponse{}, errors.NewHttpError(err)
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusUnauthorized {
		return types.QueryPaginationResponse{}, errors.NewHttpErrorUnauthorized()
	}

	if res.StatusCode == http.StatusNotFound {
		return types.QueryPaginationResponse{}, errors.NewHttpErrorNotFound()
	}

	body := types.QueryPaginationResponse{}
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return types.QueryPaginationResponse{}, errors.NewHttpError(err)
	}
	return body, nil
}

func (q queryApiService) Update(arguments types.QueryUpdateArguments) (types.QueryResponse, errors.HttpError) {
	params := url.Values{}
	resource := q.config.BaseApiUrl + "/" + arguments.Sid

	if len(arguments.SampleSid) != 0 {
		params.Set("SampleSid", arguments.SampleSid)
	}

	if len(arguments.Status.String()) != 0 {
		params.Set("Status", arguments.Status.String())
	}

	req, err := http.NewRequest("POST", resource, strings.NewReader(params.Encode()))
	if err != nil {
		return types.QueryResponse{}, errors.NewHttpError(err)
	}
	req.SetBasicAuth(q.config.TwilioApiKeySid, q.config.TwilioApiKeySecret)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := q.config.Client.Do(req)
	if err != nil {
		return types.QueryResponse{}, errors.NewHttpError(err)
	}
	if res.StatusCode != http.StatusOK {
		return types.QueryResponse{}, errors.NewHttpErrorNotUpdatedWithResource(q.config.BaseApiUrl + ": " + params.Encode())
	}
	defer res.Body.Close()

	body := types.QueryResponse{}
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return types.QueryResponse{}, errors.NewHttpError(err)
	}

	return body, nil
}

func (q queryApiService) Delete(querySid string) errors.HttpError {
	requestUrl := q.config.BaseApiUrl + "/" + querySid
	req, err := http.NewRequest("DELETE", requestUrl, nil)
	if err != nil {
		return errors.NewHttpError(err)
	}
	req.SetBasicAuth(q.config.TwilioApiKeySid, q.config.TwilioApiKeySecret)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := q.config.Client.Do(req)
	if err != nil {
		return errors.NewHttpError(err)
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusUnauthorized {
		return errors.NewHttpErrorUnauthorized()
	}

	if res.StatusCode == http.StatusNotFound {
		return errors.NewHttpErrorNotFound()
	}

	return nil
}
