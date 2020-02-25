package queries

import (
	"github.com/expproletariy/twililo/autopilot/queries/enums"
	"github.com/expproletariy/twililo/autopilot/queries/types"
	commontypes "github.com/expproletariy/twililo/common/types"
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

func (q queryApiService) Create(arguments types.QueryCreateArguments) (types.QueryResponse, error) {
	panic("implement me")
}

func (q queryApiService) GetBySid(querySid string) (types.QueryResponse, error) {
	panic("implement me")
}

func (q queryApiService) Get(meta types.Meta) (types.QueryPaginationResponse, error) {
	panic("implement me")
}

func (q queryApiService) Update(querySid string, status enums.QueryStatus) (types.QueryResponse, error) {
	panic("implement me")
}

func (q queryApiService) Delete(querySid string) error {
	panic("implement me")
}
