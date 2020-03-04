package twilio

import (
	"github.com/expproletariy/twilio/autopilot"
	commontypes "github.com/expproletariy/twilio/common/types"
)

type twilioApiService struct {
	config commontypes.Config
}

func New(config commontypes.Config) Twilio {
	return newTwilioApiService(config)
}

func newTwilioApiService(config commontypes.Config) Twilio {
	config.BaseApiUrl = "https://autopilot.twilio.com"
	config.ApiVersion = "v1"
	config.BaseApiUrl += "/" + config.ApiVersion
	return &twilioApiService{config: config}
}

func (t twilioApiService) Autopilot(botSid string) autopilot.Autopilot {
	return autopilot.New(t.config, botSid)
}
