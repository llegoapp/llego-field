package field

/* FieldRepository defines the interface for CRUD operations on Field entities.*/
type FieldRepository interface {
	CreateField(field Field) error

	GetField(id int) (Field, error)

	UpdateField(field Field) error

	DeleteField(id int) error

	ListFields(page, pageSize int) ([]*Field, int, error)

	ListFieldsByUserId(id int) ([]*Field, error)
}
