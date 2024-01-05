package adapter

import (
	"database/sql"
	"fields/internal/field"
	"fields/pkg/apperror"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type FieldRepositoryDB struct {
	db *sql.DB
}

func NewFieldRepositoryDB(conn *sql.DB) field.Repository {
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

func (repo *FieldRepositoryDB) GetField(id uuid.UUID) (field.Field, error) {
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

func (repo *FieldRepositoryDB) DeleteField(id uuid.UUID) error {
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

func (repo *FieldRepositoryDB) ListFieldsByOwnerId(id uuid.UUID) ([]*field.Field, error) {
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

func (repo *FieldRepositoryDB) ListAvailableFields(startTime, endTime time.Time, page, pageSize int) ([]*field.Field, int, error) {
	// Adjust page number and set offset for SQL query
	if page <= 0 {
		page = 1
	}
	offset := (page - 1) * pageSize

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
    LIMIT $3 OFFSET $4
    `

	rows, err := repo.db.Query(query, startTime, endTime, pageSize, offset)
	if err != nil {
		return nil, 0, apperror.NewInternalError(fmt.Sprintf("error listing available fields: %v", err))
	}
	defer rows.Close()

	var fields []*field.Field
	for rows.Next() {
		var field field.Field
		if err := rows.Scan(&field.Id, &field.OwnerId, &field.Location.Street, &field.Location.City,
			&field.Location.Country, &field.Status, &field.OpenTime, &field.CloseTime); err != nil {
			return nil, 0, apperror.NewInternalError(fmt.Sprintf("error scanning field: %v", err))
		}
		fields = append(fields, &field)
	}

	// Count total available fields for pagination
	countQuery := `
    SELECT COUNT(*) 
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

	var totalCount int
	err = repo.db.QueryRow(countQuery, startTime, endTime).Scan(&totalCount)
	if err != nil {
		return nil, 0, apperror.NewInternalError(fmt.Sprintf("error counting available fields: %v", err))
	}

	return fields, totalCount, nil
}

func (repo *FieldRepositoryDB) CheckFieldAvailability(fieldId uuid.UUID, startTime, endTime time.Time) (bool, error) {

	// Check if the field is in status 'available'
	statusQuery := `
    SELECT status FROM fields WHERE id = $1`
	var currentStatus field.FieldStatus
	err := repo.db.QueryRow(statusQuery, fieldId).Scan(&currentStatus)
	if err != nil {
		return false, apperror.NewInternalError(fmt.Sprintf("error checking field status: %v", err))
	}
	if currentStatus.IsAvalible() {
		return false, apperror.NewBadRequestError("field is not available")
	}

	reservationQuery := `
    SELECT EXISTS (
        SELECT 1 FROM reservations 
        WHERE field_id = $1 
        AND start_time < $3 
        AND end_time > $2
    )`
	var isReserved bool
	err = repo.db.QueryRow(reservationQuery, fieldId, startTime, endTime).Scan(&isReserved)
	if err != nil {
		return false, apperror.NewInternalError(fmt.Sprintf("error checking reservation availability: %v", err))
	}
	if isReserved {
		return false, apperror.NewBadRequestError("field is not available for the specified time")
	}
	// Query to check if the reservation time is within the field's open and close times
	fieldTimeQuery := `
    SELECT EXISTS (
        SELECT 1 FROM fields 
        WHERE id = $1 
        AND open_time <= $2 
        AND close_time >= $3
    )`
	var isOpenForReservation bool
	err = repo.db.QueryRow(fieldTimeQuery, fieldId, startTime, endTime).Scan(&isOpenForReservation)
	if err != nil {
		return false, apperror.NewInternalError(fmt.Sprintf("error checking field open times: %v", err))
	}
	if !isOpenForReservation {
		return false, apperror.NewBadRequestError("field is not open for reservation during the specified time")
	}

	return true, nil
}