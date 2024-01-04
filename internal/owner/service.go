package owner

type Service struct {
	repo Repository
}

func NewOwnerService(repo Repository) Service {
	return Service{
		repo,
	}
}

func (s *Service) UpdateOwner(o Owner) error {
	//TODO: check if the owner is the ower, use a bearer token to check the id
	owner, err := s.repo.GetOwner(o.Id)
	if err != nil {
		return err
	}

	if err = o.Validate(); err != nil {
		return err
	}
	err = owner.Update(o)
	if err != nil {
		return err
	}
	err = s.repo.UpdateOwner(owner)
	if err != nil {
		return err
	}

	return nil

}

func (s *Service) GetOwner(id int) (Owner, error) {
	return s.repo.GetOwner(id)
}
