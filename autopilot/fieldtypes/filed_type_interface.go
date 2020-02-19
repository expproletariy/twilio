package fieldtypes

type FieldType interface {
	Create()
	GetBySid()
	Get()
	Update()
	Delete()
}
