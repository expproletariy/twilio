package twilio

import (
	"github.com/expproletariy/twililo/autopilot"
)

type Twilio interface {
	Autopilot(assistantSid string) autopilot.Autopilot
}
