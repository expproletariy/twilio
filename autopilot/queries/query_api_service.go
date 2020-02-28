package queries

import (
	"github.com/expproletariy/twilio/autopilot/queries/enums"
	"github.com/expproletariy/twilio/autopilot/queries/types"
	"github.com/expproletariy/twilio/common/errors"
	commontypes "github.com/expproletariy/twilio/common/types"
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
	panic("implement me")
}

func (q queryApiService) GetBySid(querySid string) (types.QueryResponse, errors.HttpError) {
	panic("implement me")
}

func (q queryApiService) Get(meta types.Meta) (types.QueryPaginationResponse, errors.HttpError) {
	panic("implement me")
}

func (q queryApiService) Update(querySid string, status enums.QueryStatus) (types.QueryResponse, errors.HttpError) {
	panic("implement me")
}

func (q queryApiService) Delete(querySid string) errors.HttpError {
	panic("implement me")
}
