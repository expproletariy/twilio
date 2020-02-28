package tasks

import (
	"github.com/expproletariy/twilio/autopilot/tasks/taskfields"
	"github.com/expproletariy/twilio/autopilot/tasks/tasksamples"
	"github.com/expproletariy/twilio/autopilot/tasks/types"
	"github.com/expproletariy/twilio/common/errors"
)

type Task interface {
	TaskSamples(taskSid string) tasksamples.TaskSample
	TaskFields(taskSid string) taskfields.TaskField
	Create(arguments types.TaskCreateArguments) (types.TaskResponse, errors.HttpError)
	GetBySid(taskSid string) (types.TaskResponse, errors.HttpError)
	Get(meta types.Meta) (types.TaskPaginationResponse, errors.HttpError)
	Update(arguments types.TaskUpdateArguments) (types.TaskResponse, errors.HttpError)
	Delete(taskSid string) errors.HttpError
}
