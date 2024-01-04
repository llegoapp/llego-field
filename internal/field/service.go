package field

import "time"

type FieldService struct {
	repo Repository
}

func NewFieldService(repo Repository) FieldService {
	return FieldService{
		repo,
	}
}

func (s *FieldService) GetField(id int) (Field, error) {
	return s.repo.GetField(id)
}

func (s *FieldService) CreateField(field Field) error {
	return s.repo.CreateField(field)
}

func (s *FieldService) DeleteField(id int) error {
	//TODO: check if the owner id is the same as the one who create the field
	return s.repo.DeleteField(id)
}

func (s *FieldService) ListFields(page, pageSize int) ([]*Field, int, error) {
	return s.repo.ListFields(page, pageSize)
}

func (s *FieldService) ListFieldsByOwnerId(id int) ([]*Field, error) {
	return s.repo.ListFieldsByOwnerId(id)
}

func (s *FieldService) ListAvailableFields(startTime, endTime time.Time, page, pageSize int) ([]*Field, int, error) {
	return s.repo.ListAvailableFields(startTime, endTime, page, pageSize)
}

func (s *FieldService) CheckFieldAvailability(fieldId int, startTime, endTime time.Time) (bool, error) {
	return s.repo.CheckFieldAvailability(fieldId, startTime, endTime)
}
