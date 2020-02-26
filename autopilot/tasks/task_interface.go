package tasks

import (
	"github.com/expproletariy/twilio/autopilot/tasks/taskfields"
	"github.com/expproletariy/twilio/autopilot/tasks/tasksamples"
	"github.com/expproletariy/twilio/autopilot/tasks/types"
)

type Task interface {
	TaskSamples(taskSid string) tasksamples.TaskSample
	TaskFields(taskSid string) taskfields.TaskField
	Create(arguments types.TaskCreateArguments) (types.TaskResponse, error)
	GetBySid(taskSid string) (types.TaskResponse, error)
	Get(meta types.Meta) (types.TaskPaginationResponse, error)
	Update(arguments types.TaskUpdateArguments) (types.TaskResponse, error)
	Delete(taskSid string) error
}
