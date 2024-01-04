package owner

type OwnerRepository interface {
	GetOwner(id int) (Owner, error)
	UpdateOwner(owner Owner) error
}
