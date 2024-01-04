package owner

type Repository interface {
	GetOwner(id int) (Owner, error)
	UpdateOwner(owner Owner) error
}
