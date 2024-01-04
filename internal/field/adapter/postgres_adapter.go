package adapter

import (
	"database/sql"
	"fields/internal/field"
	"fields/pkg/apperror"
	"fmt"
	"time"
)

type FieldRepositoryDB struct {
	db *sql.DB
}

func NewFieldRepositoryDB(conn *sql.DB) field.FieldRepository {
	return &FieldRepositoryDB{
		conn,
	}
}

func (repo *FieldRepositoryDB) CreateField(f field.Field) error {
	query := `INSERT INTO fields (owner_id, street, city, country, status, open_time, close_time) 
              VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := repo.db.Exec(query, f.OwnerId, f.Location.Street, f.Location.City,
		f.Location.Country, f.Status, f.OpenTime, f.CloseTime)
	return apperror.NewInternalError(fmt.Sprintf("Error creating field: %v", err))

}

func (repo *FieldRepositoryDB) GetField(id int) (field.Field, error) {
	query := `SELECT id, owner_id, street, city, country, status, open_time, close_time FROM fields WHERE id = $1`
	var f field.Field

	err := repo.db.QueryRow(query, id).Scan(&f.Id, &f.OwnerId, &f.Location.Street,
		&f.Location.City, &f.Location.Country, &f.Status,
		&f.OpenTime, &f.CloseTime)
	if err != nil {
		return field.Field{}, apperror.NewInternalError(fmt.Sprintf("error getting field: %v", err))
	}
	return f, nil
}

func (repo *FieldRepositoryDB) UpdateField(f field.Field) error {
	query := `UPDATE fields SET owner_id = $1, street = $2, city = $3, country = $4, status = $5, 
              open_time = $6, close_time = $7 WHERE id = $8`
	_, err := repo.db.Exec(query, f.OwnerId, f.Location.Street, f.Location.City,
		f.Location.Country, f.Status, f.OpenTime, f.CloseTime, f.Id)
	return apperror.NewInternalError(fmt.Sprintf("error updating field: %v", err))
}

func (repo *FieldRepositoryDB) DeleteField(id int) error {
	query := `DELETE FROM fields WHERE id = $1`
	_, err := repo.db.Exec(query, id)
	return apperror.NewInternalError(fmt.Sprintf("error deleting field: %v", err))
}

func (repo *FieldRepositoryDB) ListFields(page, pageSize int) ([]*field.Field, int, error) {
	query := `SELECT id, owner_id, street, city, country, status, open_time, close_time FROM fields
              LIMIT $1 OFFSET $2`
	rows, err := repo.db.Query(query, pageSize, (page-1)*pageSize)
	if err != nil {
		return nil, 0, apperror.NewInternalError(fmt.Sprintf("error listing fields: %v", err))
	}
	defer rows.Close()

	fields := make([]*field.Field, 0)
	for rows.Next() {
		var field field.Field
		if err := rows.Scan(&field.Id, &field.OwnerId, &field.Location.Street, &field.Location.City,
			&field.Location.Country, &field.Status, &field.OpenTime, &field.CloseTime); err != nil {
			return nil, 0, apperror.NewInternalError(fmt.Sprintf("error scanning field: %v", err))
		}
		fields = append(fields, &field)
	}

	var totalCount int
	countQuery := `SELECT COUNT(*) FROM fields`
	err = repo.db.QueryRow(countQuery).Scan(&totalCount)
	if err != nil {
		return nil, 0, apperror.NewInternalError(fmt.Sprintf("error counting fields: %v", err))
	}

	return fields, totalCount, nil
}

func (repo *FieldRepositoryDB) ListFieldsByUserId(id int) ([]*field.Field, error) {
	query := `SELECT id, owner_id, street, city, country, status, open_time, close_time FROM fields WHERE owner_id = $1`
	rows, err := repo.db.Query(query, id)
	if err != nil {
		return nil, apperror.NewInternalError(fmt.Sprintf("error listing fields: %v", err))
	}
	defer rows.Close()

	fields := make([]*field.Field, 0)
	for rows.Next() {
		var field field.Field
		if err := rows.Scan(&field.Id, &field.OwnerId, &field.Location.Street, &field.Location.City,
			&field.Location.Country, &field.Status, &field.OpenTime, &field.CloseTime); err != nil {
			return nil, apperror.NewInternalError(fmt.Sprintf("error scanning field: %v", err))
		}
		fields = append(fields, &field)
	}

	return fields, nil
}

func (repo *FieldRepositoryDB) ListAvailableFields(startTime, endTime time.Time) ([]*field.Field, error) {
	query := `
    SELECT f.id, f.owner_id, f.street, f.city, f.country, f.status, f.open_time, f.close_time 
    FROM fields f
    WHERE NOT EXISTS (
        SELECT 1
        FROM reservations r
        WHERE r.field_id = f.id
        AND r.start_time < $2
        AND r.end_time > $1
    )
    AND f.open_time <= $1
    AND f.close_time >= $2
    `

	rows, err := repo.db.Query(query, startTime, endTime)
	if err != nil {
		return nil, apperror.NewInternalError(fmt.Sprintf("error listing fields: %v", err))
	}
	defer rows.Close()

	var f []*field.Field
	for rows.Next() {
		var field field.Field
		if err := rows.Scan(&field.Id, &field.OwnerId, &field.Location.Street, &field.Location.City,
			&field.Location.Country, &field.Status, &field.OpenTime, &field.CloseTime); err != nil {
			return nil, apperror.NewInternalError(fmt.Sprintf("error scanning field: %v", err))
		}
		f = append(f, &field)
	}

	return f, nil
}
