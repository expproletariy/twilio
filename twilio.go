package twilio

import (
	"github.com/expproletariy/twililo/autopilot"
	commontypes "github.com/expproletariy/twililo/common/types"
)

type twilioApiService struct {
	config commontypes.Config
}

func New(config commontypes.Config) Twilio {
	return newTwilioApiService(config)
}

func newTwilioApiService(config commontypes.Config) Twilio {
	config.BaseApiUrl += "/" + config.ApiVersion
	return &twilioApiService{config: config}
}

func (t twilioApiService) Autopilot(botSid string) autopilot.Autopilot {
	return autopilot.New(t.config, botSid)
}
