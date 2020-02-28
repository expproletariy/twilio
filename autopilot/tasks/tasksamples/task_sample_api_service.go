package tasksamples

import (
	"github.com/expproletariy/twilio/autopilot/tasks/tasksamples/types"
	"github.com/expproletariy/twilio/common/errors"
	commontypes "github.com/expproletariy/twilio/common/types"
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

func (t taskSampleApiService) Create(arguments types.TaskSampleCreateArguments) (types.TaskSampleResponse, errors.HttpError) {
	panic("implement me")
}

func (t taskSampleApiService) Get(meta types.Meta) (types.TaskSamplePaginationResponse, errors.HttpError) {
	panic("implement me")
}

func (t taskSampleApiService) GetBySid(taskSampleSid string) (types.TaskSampleResponse, errors.HttpError) {
	panic("implement me")
}

func (t taskSampleApiService) Update(arguments types.TaskSampleUpdateArguments) (types.TaskSampleResponse, errors.HttpError) {
	panic("implement me")
}

func (t taskSampleApiService) Delete(taskSampleSid string) errors.HttpError {
	panic("implement me")
}
