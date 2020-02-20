package queries

import (
	"github.com/expproletariy/twililo/autopilot/queries/enums"
	"github.com/expproletariy/twililo/autopilot/queries/types"
)

type Query interface {
	Create(arguments types.QueryCreateArguments) (types.QueryResponse, error)
	GetBySid(querySid string) (types.QueryResponse, error)
	Get(meta types.Meta) (types.QueryPaginationResponse, error)
	Update(querySid string, status enums.QueryStatus) (types.QueryResponse, error)
	Delete(querySid string) error
}
