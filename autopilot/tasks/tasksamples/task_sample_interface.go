package tasksamples

import "github.com/expproletariy/twilio/autopilot/tasks/tasksamples/types"

type TaskSample interface {
	Create(arguments types.TaskSampleCreateArguments) (types.TaskSampleResponse, error)
	Get(meta types.Meta) (types.TaskSamplePaginationResponse, error)
	GetBySid(taskSampleSid string) (types.TaskSampleResponse, error)
	Update(arguments types.TaskSampleUpdateArguments) (types.TaskSampleResponse, error)
	Delete(taskSampleSid string) error
}
