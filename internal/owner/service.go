package owner

type OwnerService struct {
	repo OwnerRepository
}

func NewOwnerService(repo OwnerRepository) OwnerService {
	return OwnerService{
		repo,
	}
}

func (s *OwnerService) UpdateOwner(o Owner) error {
	//TODO: check if the owner is the ower, use a bearer token to check the id
	owner, err := s.repo.GetOwner(o.Id)
	if err != nil {
		return err
	}

	if err = o.Validate(); err != nil {
		return err
	}
	owner.Update(o)
	s.repo.UpdateOwner(owner)

	return nil

}

func (s *OwnerService) GetOwner(id int) (Owner, error) {
	return s.repo.GetOwner(id)
}
