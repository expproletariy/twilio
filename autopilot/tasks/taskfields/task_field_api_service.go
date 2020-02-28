package taskfields

import (
	"github.com/expproletariy/twilio/autopilot/tasks/taskfields/types"
	"github.com/expproletariy/twilio/common/errors"
	commontypes "github.com/expproletariy/twilio/common/types"
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

func (t taskFieldApiService) Create(arguments types.TaskFieldCreateArguments) (types.TaskFieldResponse, errors.HttpError) {
	panic("implement me")
}

func (t taskFieldApiService) GetBySid(taskFieldSid string) (types.TaskFieldResponse, errors.HttpError) {
	panic("implement me")
}

func (t taskFieldApiService) Get(meta types.Meta) (types.TaskFieldPaginationResponse, errors.HttpError) {
	panic("implement me")
}

func (t taskFieldApiService) Delete(taskFieldSid string) errors.HttpError {
	panic("implement me")
}
