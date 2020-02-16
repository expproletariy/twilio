package autopilot

import "github.com/expproletariy/twililo/autopilot/types"

type Query interface {
	Create(arguments types.QueryCreateArguments) (types.QueryCreateResult, error)
	Get()
	Update()
	Delete()
}
