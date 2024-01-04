package reservation

import "fields/internal/field"

type ReservationService struct {
	repo          ReservationRepository
	field_service field.FieldService
}

func NewReservationService(repo ReservationRepository, field_service field.FieldService) ReservationService {
	return ReservationService{
		repo,
		field_service,
	}
}

func (s *ReservationService) CreateReservation(r Reservation) error {
	if _, err := s.field_service.CheckFieldAvailability(r.FieldId, r.StartTime, r.EndTime); err != nil {
		return err
	}
	r.Details.Status = "pending"

	return s.repo.CreateReservation(r)
}

func (s *ReservationService) GetReservation(id int) (Reservation, error) {
	return s.repo.GetReservation(id)
}

func (s *ReservationService) CancelReservation(reservationId int) error {
	//TODO: Add logic of who can cancell , and the rules, before canceling
	r, err := s.GetReservation(reservationId)
	if err != nil {
		return err
	}

	r.Details.Status = "cancelled"

	return s.repo.UpdateReservation(r)
}

func (s *ReservationService) ListReservation(page, pageSize int) ([]*Reservation, int, error) {
	return s.repo.ListReservation(page, pageSize)
}

func (s *ReservationService) ListReservationByBookerId(bookerId int, page, pageSize int) ([]*Reservation, int, error) {
	return s.repo.ListReservationByBookerId(bookerId, page, pageSize)
}

func (s *ReservationService) ListReservationByFieldId(fieldId int, page, pageSize int) ([]*Reservation, int, error) {
	return s.repo.ListReservationByFieldId(fieldId, page, pageSize)
}
