package field

import (
	"time"

	"github.com/google/uuid"
)

type Service struct {
	repo Repository
}

func NewFieldService(repo Repository) Service {
	return Service{
		repo,
	}
}

func (s *Service) GetField(id uuid.UUID) (Field, error) {
	return s.repo.GetField(id)
}

func (s *Service) CreateField(field Field) error {
	return s.repo.CreateField(field)
}

func (s *Service) DeleteField(id uuid.UUID) error {
	//TODO: check if the owner id is the same as the one who create the field
	return s.repo.DeleteField(id)
}

func (s *Service) ListFields(page, pageSize int) ([]*Field, int, error) {
	return s.repo.ListFields(page, pageSize)
}

func (s *Service) ListFieldsByOwnerId(id uuid.UUID) ([]*Field, error) {
	return s.repo.ListFieldsByOwnerId(id)
}

func (s *Service) ListAvailableFields(startTime, endTime time.Time, page, pageSize int) ([]*Field, int, error) {
	return s.repo.ListAvailableFields(startTime, endTime, page, pageSize)
}

func (s *Service) CheckFieldAvailability(fieldId uuid.UUID, startTime, endTime time.Time) (bool, error) {
	return s.repo.CheckFieldAvailability(fieldId, startTime, endTime)
}