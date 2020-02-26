package twilio

import (
	"github.com/expproletariy/twilio/autopilot"
)

type Twilio interface {
	Autopilot(botSid string) autopilot.Autopilot
}
