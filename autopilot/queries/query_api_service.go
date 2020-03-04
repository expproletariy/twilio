package queries

import (
	"encoding/json"
	"github.com/expproletariy/twilio/autopilot/queries/enums"
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
		return types.QueryResponse{}, errors.NewHttpErrorNotCreated()
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
	requestUrl := t.config.BaseApiUrl + "/" + querySid
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

func (q queryApiService) Update(querySid string, status enums.QueryStatus) (types.QueryResponse, errors.HttpError) {
	panic("implement me")
}

func (q queryApiService) Delete(querySid string) errors.HttpError {
	panic("implement me")
}
