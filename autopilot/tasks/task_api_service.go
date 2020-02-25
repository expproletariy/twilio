package tasks

import (
	"github.com/expproletariy/twililo/autopilot/tasks/taskfields"
	"github.com/expproletariy/twililo/autopilot/tasks/tasksamples"
	"github.com/expproletariy/twililo/autopilot/tasks/types"
	commontypes "github.com/expproletariy/twililo/common/types"
)

type taskApiService struct {
	config commontypes.Config
}

func New(config commontypes.Config) Task {
	return newTaskApiService(config)
}

func newTaskApiService(config commontypes.Config) Task {
	config.BaseApiUrl += "/Tasks"
	return &taskApiService{config: config}
}

func (t taskApiService) TaskSamples(taskSid string) tasksamples.TaskSample {
	config := commontypes.Config{
		Client:             t.config.Client,
		BaseApiUrl:         t.config.BaseApiUrl + "/" + taskSid,
		ApiVersion:         t.config.ApiVersion,
		TwilioApiKeySid:    t.config.TwilioApiKeySid,
		TwilioApiKeySecret: t.config.TwilioApiKeySecret,
	}
	return tasksamples.New(config)
}

func (t taskApiService) TaskFields(taskSid string) taskfields.TaskField {
	config := commontypes.Config{
		Client:             t.config.Client,
		BaseApiUrl:         t.config.BaseApiUrl + "/" + taskSid,
		ApiVersion:         t.config.ApiVersion,
		TwilioApiKeySid:    t.config.TwilioApiKeySid,
		TwilioApiKeySecret: t.config.TwilioApiKeySecret,
	}
	return taskfields.New(config)
}

func (t taskApiService) Create(arguments types.TaskCreateArguments) (types.TaskResponse, error) {
	panic("implement me")
}

func (t taskApiService) GetBySid(taskSid string) (types.TaskResponse, error) {
	panic("implement me")
}

func (t taskApiService) Get(meta types.Meta) (types.TaskPaginationResponse, error) {
	panic("implement me")
}

func (t taskApiService) Update(arguments types.TaskUpdateArguments) (types.TaskResponse, error) {
	panic("implement me")
}

func (t taskApiService) Delete(taskSid string) error {
	panic("implement me")
}
