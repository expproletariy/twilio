package taskfields

import (
	"github.com/expproletariy/twililo/autopilot/tasks/taskfields/types"
	commontypes "github.com/expproletariy/twililo/common/types"
)

type taskFieldApiService struct {
	config commontypes.Config
}

func New(config commontypes.Config) TaskField {
	return newTaskFieldApiService(config)
}

func newTaskFieldApiService(config commontypes.Config) *taskFieldApiService {
	config.BaseApiUrl += "/Fields"
	return &taskFieldApiService{config: config}
}

func (t taskFieldApiService) Create(arguments types.TaskFieldCreateArguments) (types.TaskFieldResponse, error) {
	panic("implement me")
}

func (t taskFieldApiService) GetBySid(taskFieldSid string) (types.TaskFieldResponse, error) {
	panic("implement me")
}

func (t taskFieldApiService) Get(meta types.Meta) (types.TaskFieldPaginationResponse, error) {
	panic("implement me")
}

func (t taskFieldApiService) Delete(taskFieldSid string) error {
	panic("implement me")
}
