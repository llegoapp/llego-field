package owner

import "github.com/google/uuid"

type Repository interface {
	GetOwner(id uuid.UUID) (Owner, error)
	UpdateOwner(owner Owner) error
}
