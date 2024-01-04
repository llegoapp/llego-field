package reservation

type ReservationRepository interface {
	CreateReservation(r Reservation) error
	GetReservation(id int) (Reservation, error)
	UpdateReservation(r Reservation) error
	ListReservation(page, pageSize int) ([]*Reservation, int, error)
	ListReservationByBookerId(bookerId int, page, pageSize int) ([]*Reservation, int, error)
	ListReservationByFieldId(fieldId int, page, pageSize int) ([]*Reservation, int, error)
}
