package field

import "time"

/* FieldRepository defines the interface for CRUD operations on Field entities.*/
type FieldRepository interface {
	CreateField(field Field) error

	GetField(id int) (Field, error)

	UpdateField(field Field) error

	DeleteField(id int) error

	ListFields(page, pageSize int) ([]*Field, int, error)

	ListFieldsByOwnerId(id int) ([]*Field, error)

	ListAvailableFields(startTime, endTime time.Time, page, pageSize int) ([]*Field, int, error)

	CheckFieldAvailability(fieldId int, startTime, endTime time.Time) (bool, error)
}
