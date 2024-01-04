package adapter

import (
	"database/sql"
	"fields/internal/owner"
	"fields/pkg/apperror"
	"fmt"
)

type OwnerRepositoryDB struct {
	db *sql.DB
}

func NewOwnerRepositoryDB(db *sql.DB) owner.Repository {
	return &OwnerRepositoryDB{
		db: db,
	}
}

func (repo *OwnerRepositoryDB) GetOwner(id int) (owner.Owner, error) {
	query := `SELECT id, name, country_code, phone_number, password FROM owners WHERE id = $1`
	var o owner.Owner
	var phoneNumber owner.PhoneNumber
	err := repo.db.QueryRow(query, id).Scan(&o.Id, &o.Name, &phoneNumber.CountryCode, &phoneNumber.Number, &o.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return owner.Owner{}, apperror.NewNotFoundError(fmt.Sprintf("Owner with ID %d not found", id))
		}
		return owner.Owner{}, apperror.NewInternalError(fmt.Sprintf("Error getting owner: %v", err))
	}
	o.PhoneNumber = phoneNumber
	return o, nil
}

func (repo *OwnerRepositoryDB) UpdateOwner(o owner.Owner) error {
	query := `UPDATE owners SET name = $1, country_code = $2, phone_number = $3, password = $4 WHERE id = $5`
	_, err := repo.db.Exec(query, o.Name, o.PhoneNumber.CountryCode, o.PhoneNumber.Number, o.Password, o.Id)
	if err != nil {
		return apperror.NewInternalError(fmt.Sprintf("Error updating owner: %v", err))
	}
	return nil
}
