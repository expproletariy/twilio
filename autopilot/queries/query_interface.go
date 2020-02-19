package queries

import (
	"github.com/expproletariy/twililo/autopilot/queries/enums"
	"github.com/expproletariy/twililo/autopilot/queries/types"
)

type Query interface {
	Create(arguments types.QueryCreateArguments) (types.QueryResult, error)
	GetBySid(querySid string) (types.QueryResult, error)
	Get(meta types.Meta) ([]types.QueryResult, types.Meta, error)
	Update(querySid string, status enums.QueryStatus) (types.QueryResult, error)
	Delete(querySid string) error
}
