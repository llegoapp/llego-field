package field

import (
	"time"

	"github.com/google/uuid"
)

/* Repository defines the interface for CRUD operations on Field entities.*/
type Repository interface {
	CreateField(field Field) error

	GetField(id uuid.UUID) (Field, error)

	UpdateField(field Field) error

	DeleteField(id uuid.UUID) error

	ListFields(page, pageSize int) ([]*Field, int, error)

	ListFieldsByOwnerId(id uuid.UUID) ([]*Field, error)

	ListAvailableFields(startTime, endTime time.Time, page, pageSize int) ([]*Field, int, error)

	CheckFieldAvailability(fieldId uuid.UUID, startTime, endTime time.Time) (bool, error)
}