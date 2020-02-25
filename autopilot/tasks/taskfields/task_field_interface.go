package taskfields

import "github.com/expproletariy/twililo/autopilot/tasks/taskfields/types"

type TaskField interface {
	Create(arguments types.TaskFieldCreateArguments) (types.TaskFieldResponse, error)
	GetBySid(taskFieldSid string) (types.TaskFieldResponse, error)
	Get(meta types.Meta) (types.TaskFieldPaginationResponse, error)
	Delete(taskFieldSid string) error
}
