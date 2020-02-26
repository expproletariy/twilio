package autopilot

import (
	"github.com/expproletariy/twilio/autopilot/fieldtypes"
	"github.com/expproletariy/twilio/autopilot/queries"
	"github.com/expproletariy/twilio/autopilot/tasks"
)

type Autopilot interface {
	Queries() queries.Query
	FieldTypes() fieldtypes.FieldType
	Tasks() tasks.Task
}
