package taskfields

import (
	"github.com/expproletariy/twilio/autopilot/tasks/taskfields/types"
	"github.com/expproletariy/twilio/common/errors"
)

type TaskField interface {
	Create(arguments types.TaskFieldCreateArguments) (types.TaskFieldResponse, errors.HttpError)
	GetBySid(taskFieldSid string) (types.TaskFieldResponse, errors.HttpError)
	Get(meta types.Meta) (types.TaskFieldPaginationResponse, errors.HttpError)
	Delete(taskFieldSid string) errors.HttpError
}
