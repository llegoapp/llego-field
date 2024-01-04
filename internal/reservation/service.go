package reservation

import "fields/internal/field"

type Service struct {
	repo          Repository
	field_service field.FieldService
}

func NewReservationService(repo Repository, field_service field.FieldService) Service {
	return Service{
		repo,
		field_service,
	}
}

func (s *Service) CreateReservation(r Reservation) error {
	if _, err := s.field_service.CheckFieldAvailability(r.FieldId, r.StartTime, r.EndTime); err != nil {
		return err
	}

	r.SetDefaultDetails()

	return s.repo.CreateReservation(r)
}

func (s *Service) GetReservation(id int) (Reservation, error) {
	return s.repo.GetReservation(id)
}

func (s *Service) CancelReservation(reservationId int) error {
	//TODO: Add logic of who can cancell , and the rules, before canceling
	r, err := s.GetReservation(reservationId)
	if err != nil {
		return err
	}

	r.Details.Status = "cancelled"

	return s.repo.UpdateReservation(r)
}

func (s *Service) ListReservation(page, pageSize int) ([]*Reservation, int, error) {
	return s.repo.ListReservation(page, pageSize)
}

func (s *Service) ListReservationByBookerId(bookerId int, page, pageSize int) ([]*Reservation, int, error) {
	return s.repo.ListReservationByBookerId(bookerId, page, pageSize)
}

func (s *Service) ListReservationByFieldId(fieldId int, page, pageSize int) ([]*Reservation, int, error) {
	return s.repo.ListReservationByFieldId(fieldId, page, pageSize)
}
