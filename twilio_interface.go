package twilio

import (
	"github.com/expproletariy/twililo/autopilot"
)

type Twilio interface {
	Autopilot(botSid string) autopilot.Autopilot
}
