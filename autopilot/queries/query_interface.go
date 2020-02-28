package queries

import (
	"github.com/expproletariy/twilio/autopilot/queries/enums"
	"github.com/expproletariy/twilio/autopilot/queries/types"
	"github.com/expproletariy/twilio/common/errors"
)

type Query interface {
	Create(arguments types.QueryCreateArguments) (types.QueryResponse, errors.HttpError)
	GetBySid(querySid string) (types.QueryResponse, errors.HttpError)
	Get(meta types.Meta) (types.QueryPaginationResponse, errors.HttpError)
	Update(querySid string, status enums.QueryStatus) (types.QueryResponse, errors.HttpError)
	Delete(querySid string) errors.HttpError
}
