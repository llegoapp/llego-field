package reservation

type ReservationRepository interface {
	CreateReservation(r Reservation) error
	GetReservation(id int) (Reservation, error)
	UpdateReservation(id Reservation) error
	ListReservation(page, pageSize int) ([]*Reservation, int, error)
	ListReservationByBookerId(id int) ([]*Reservation, int, error)
	ListReservationByFieldId(id int) ([]*Reservation, int, error)
}
