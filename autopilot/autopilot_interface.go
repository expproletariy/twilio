package autopilot

import (
	"github.com/expproletariy/twililo/autopilot/fieldtypes"
	"github.com/expproletariy/twililo/autopilot/queries"
)

type Autopilot interface {
	Queries() queries.Query
	FieldTypes() fieldtypes.FieldType
}
