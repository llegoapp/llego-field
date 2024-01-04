package field

import "time"

type FieldService struct {
	repo FieldRepository
}

func NewFieldService(repo FieldRepository) FieldService {
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

func (s *FieldService) ListFieldsByUserId(id int) ([]*Field, error) {
	return s.repo.ListFieldsByUserId(id)
}

func (s *FieldService) ListAvailableFields(startTime, endTime time.Time) ([]*Field, error) {
	return s.repo.ListAvailableFields(startTime, endTime)
}
