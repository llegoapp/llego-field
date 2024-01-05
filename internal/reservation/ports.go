package reservation

import "github.com/google/uuid"

type Repository interface {
	CreateReservation(r Reservation) error
	GetReservation(id uuid.UUID) (Reservation, error)
	UpdateReservation(r Reservation) error
	ListReservation(page, pageSize int) ([]*Reservation, int, error)
	ListReservationByBookerId(bookerId uuid.UUID, page, pageSize int) ([]*Reservation, int, error)
	ListReservationByFieldId(fieldId int, page, pageSize int) ([]*Reservation, int, error)
}
