package tasksamples

import (
	"github.com/expproletariy/twililo/autopilot/tasks/tasksamples/types"
	commontypes "github.com/expproletariy/twililo/common/types"
)

type taskSampleApiService struct {
	config commontypes.Config
}

func New(config commontypes.Config) TaskSample {
	return newTaskSampleApiService(config)
}

func newTaskSampleApiService(config commontypes.Config) TaskSample {
	config.BaseApiUrl = "/Samples"
	return &taskSampleApiService{config: config}
}

func (t taskSampleApiService) Create(arguments types.TaskSampleCreateArguments) (types.TaskSampleResponse, error) {
	panic("implement me")
}

func (t taskSampleApiService) Get(meta types.Meta) (types.TaskSamplePaginationResponse, error) {
	panic("implement me")
}

func (t taskSampleApiService) GetBySid(taskSampleSid string) (types.TaskSampleResponse, error) {
	panic("implement me")
}

func (t taskSampleApiService) Update(arguments types.TaskSampleUpdateArguments) (types.TaskSampleResponse, error) {
	panic("implement me")
}

func (t taskSampleApiService) Delete(taskSampleSid string) error {
	panic("implement me")
}
