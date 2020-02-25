package autopilot

import (
	"github.com/expproletariy/twililo/autopilot/fieldtypes"
	"github.com/expproletariy/twililo/autopilot/queries"
	"github.com/expproletariy/twililo/autopilot/tasks"
)

type Autopilot interface {
	Queries() queries.Query
	FieldTypes() fieldtypes.FieldType
	Tasks() tasks.Task
}
