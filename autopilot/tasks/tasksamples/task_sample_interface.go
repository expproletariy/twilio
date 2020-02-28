package tasksamples

import (
	"github.com/expproletariy/twilio/autopilot/tasks/tasksamples/types"
	"github.com/expproletariy/twilio/common/errors"
)

type TaskSample interface {
	Create(arguments types.TaskSampleCreateArguments) (types.TaskSampleResponse, errors.HttpError)
	Get(meta types.Meta) (types.TaskSamplePaginationResponse, errors.HttpError)
	GetBySid(taskSampleSid string) (types.TaskSampleResponse, errors.HttpError)
	Update(arguments types.TaskSampleUpdateArguments) (types.TaskSampleResponse, errors.HttpError)
	Delete(taskSampleSid string) errors.HttpError
}
