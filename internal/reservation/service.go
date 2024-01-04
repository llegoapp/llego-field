package reservation

type ReservationService struct {
	repo ReservationRepository
}

func NewReservationService(repo ReservationRepository) ReservationService {
	return ReservationService{
		repo,
	}
}

func (s *ReservationService) CreateReservation(r Reservation) error {
	return s.repo.CreateReservation(r)
}

func (s *ReservationService) GetReservation(id int) (Reservation, error) {
	return s.repo.GetReservation(id)
}

func (s *ReservationService) UpdateReservation(r Reservation) error {
	return s.repo.UpdateReservation(r)
}

func (s *ReservationService) ListReservation(page, pageSize int) ([]*Reservation, int, error) {
	return s.repo.ListReservation(page, pageSize)
}

func (s *ReservationService) ListReservationByBookerId(id int) ([]*Reservation, int, error) {
	return s.repo.ListReservationByBookerId(id)
}

func (s *ReservationService) ListReservationByFieldId(id int) ([]*Reservation, int, error) {
	return s.repo.ListReservationByFieldId(id)
}
