package autopilot

import (
	"github.com/expproletariy/twilio/autopilot/fieldtypes"
	"github.com/expproletariy/twilio/autopilot/queries"
	"github.com/expproletariy/twilio/autopilot/tasks"
	commontypes "github.com/expproletariy/twilio/common/types"
)

type autopilotApiService struct {
	config commontypes.Config
}

func New(config commontypes.Config, botSid string) Autopilot {
	return newAutopilotApiService(config, botSid)
}

func newAutopilotApiService(config commontypes.Config, botSid string) Autopilot {
	config.BaseApiUrl += "/Assistants/" + botSid
	return &autopilotApiService{config: config}
}

func (a autopilotApiService) Queries() queries.Query {
	return queries.New(a.config)
}

func (a autopilotApiService) FieldTypes() fieldtypes.FieldType {
	return fieldtypes.New(a.config)
}

func (a autopilotApiService) Tasks() tasks.Task {
	return tasks.New(a.config)
}
