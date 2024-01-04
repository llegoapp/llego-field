package adapter

import (
	"database/sql"
	"fields/internal/reservation"
	"fields/pkg/apperror"
	"fmt"
	"time"
)

type ReservationRepositoryDB struct {
	db *sql.DB
}

func NewReservationRepositoryDB(db *sql.DB) reservation.ReservationRepository {
	return &ReservationRepositoryDB{
		db: db,
	}
}

func (repo *ReservationRepositoryDB) CreateReservation(r reservation.Reservation) error {
	// Check if the field is available for reservation
	if err := repo.checkFieldAvailability(r.FieldId, r.StartTime, r.EndTime); err != nil {
		return err
	}

	query := `INSERT INTO reservations (field_id, booker_id, start_time, end_time, status,payment_status,payment_id) VALUES ($1, $2, $3, $4, $5, $6 ,$7)`
	_, err := repo.db.Exec(query, r.FieldId, r.BookerId, r.StartTime, r.EndTime, r.Details.Status, r.Details.PaymentStatus, r.Details.PaymentID)
	if err != nil {
		return apperror.NewInternalError(fmt.Sprintf("error creating reservation: %v", err))
	}
	return nil
}

// checkFieldAvailability checks if the field is available for the given time period
// TODO: this should be in the reservation service with the fild service for be buisiness logic
func (repo *ReservationRepositoryDB) checkFieldAvailability(fieldId int, startTime, endTime time.Time) error {
	// Query to check if the field is already reserved during the given time
	reservationQuery := `
    SELECT EXISTS (
        SELECT 1 FROM reservations 
        WHERE field_id = $1 
        AND start_time < $3 
        AND end_time > $2
    )`
	var isReserved bool
	err := repo.db.QueryRow(reservationQuery, fieldId, startTime, endTime).Scan(&isReserved)
	if err != nil {
		return apperror.NewInternalError(fmt.Sprintf("error checking reservation availability: %v", err))
	}
	if isReserved {
		return apperror.NewBadRequestError("field is not available for the specified time")
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
		return apperror.NewInternalError(fmt.Sprintf("error checking field open times: %v", err))
	}
	if !isOpenForReservation {
		return apperror.NewBadRequestError("field is not open for reservation during the specified time")
	}

	return nil
}

func (repo *ReservationRepositoryDB) GetReservation(id int) (reservation.Reservation, error) {
	query := `SELECT id, field_id, booker_id, start_time, end_time, status, payment_status, payment_id FROM reservations WHERE id = $1`
	var r reservation.Reservation
	err := repo.db.QueryRow(query, id).Scan(&r.Id, &r.FieldId, &r.BookerId, &r.StartTime, &r.EndTime, &r.Details.Status, &r.Details.PaymentStatus, &r.Details.PaymentID)
	if err != nil {
		if err == sql.ErrNoRows {
			return reservation.Reservation{}, apperror.NewNotFoundError(fmt.Sprintf("Reservation with ID %d not found", id))
		}
		return reservation.Reservation{}, apperror.NewInternalError(fmt.Sprintf("Error getting reservation: %v", err))
	}
	return r, nil
}

func (repo *ReservationRepositoryDB) UpdateReservation(r reservation.Reservation) error {
	query := `UPDATE reservations SET field_id = $1, booker_id = $2, start_time = $3, end_time = $4, status = $5, payment_status = $6, payment_id = $7 WHERE id = $8`
	_, err := repo.db.Exec(query, r.FieldId, r.BookerId, r.StartTime, r.EndTime, r.Details.Status, r.Details.PaymentStatus, r.Details.PaymentID, r.Id)
	if err != nil {
		return apperror.NewInternalError(fmt.Sprintf("Error updating reservation: %v", err))
	}
	return nil
}

func (repo *ReservationRepositoryDB) ListReservation(page, pageSize int) ([]*reservation.Reservation, int, error) {
	if page <= 0 {
		page = 1
	}
	offset := (page - 1) * pageSize

	query := `
    SELECT id, field_id, booker_id, start_time, end_time, status, payment_status, payment_id 
    FROM reservations 
    LIMIT $1 OFFSET $2`
	rows, err := repo.db.Query(query, pageSize, offset)
	if err != nil {
		return nil, 0, apperror.NewInternalError(fmt.Sprintf("Error listing reservations: %v", err))
	}
	defer rows.Close()

	var reservations []*reservation.Reservation
	for rows.Next() {
		var r reservation.Reservation
		details := &reservation.ReservationDetails{}
		if err := rows.Scan(&r.Id, &r.FieldId, &r.BookerId, &r.StartTime, &r.EndTime, &details.Status, &details.PaymentStatus, &details.PaymentID); err != nil {
			return nil, 0, apperror.NewInternalError(fmt.Sprintf("Error scanning reservation: %v", err))
		}
		r.Details = details
		reservations = append(reservations, &r)
	}

	countQuery := `SELECT COUNT(*) FROM reservations`
	var totalCount int
	err = repo.db.QueryRow(countQuery).Scan(&totalCount)
	if err != nil {
		return nil, 0, apperror.NewInternalError(fmt.Sprintf("Error counting reservations: %v", err))
	}

	return reservations, totalCount, nil
}

func (repo *ReservationRepositoryDB) ListReservationByBookerId(bookerId int, page, pageSize int) ([]*reservation.Reservation, int, error) {
	if page <= 0 {
		page = 1
	}
	offset := (page - 1) * pageSize

	query := `
    SELECT id, field_id, booker_id, start_time, end_time, status, payment_status, payment_id 
    FROM reservations 
    WHERE booker_id = $1
    LIMIT $2 OFFSET $3`
	rows, err := repo.db.Query(query, bookerId, pageSize, offset)
	if err != nil {
		return nil, 0, apperror.NewInternalError(fmt.Sprintf("Error listing reservations by booker ID: %v", err))
	}
	defer rows.Close()

	var reservations []*reservation.Reservation
	for rows.Next() {
		var r reservation.Reservation
		details := &reservation.ReservationDetails{}
		if err := rows.Scan(&r.Id, &r.FieldId, &r.BookerId, &r.StartTime, &r.EndTime, &details.Status, &details.PaymentStatus, &details.PaymentID); err != nil {
			return nil, 0, apperror.NewInternalError(fmt.Sprintf("Error scanning reservation: %v", err))
		}
		r.Details = details
		reservations = append(reservations, &r)
	}

	countQuery := `SELECT COUNT(*) FROM reservations WHERE booker_id = $1`
	var totalCount int
	err = repo.db.QueryRow(countQuery, bookerId).Scan(&totalCount)
	if err != nil {
		return nil, 0, apperror.NewInternalError(fmt.Sprintf("Error counting reservations by booker ID: %v", err))
	}

	return reservations, totalCount, nil
}

func (repo *ReservationRepositoryDB) ListReservationByFieldId(fieldId int, page, pageSize int) ([]*reservation.Reservation, int, error) {
	if page <= 0 {
		page = 1
	}
	offset := (page - 1) * pageSize

	query := `
    SELECT id, field_id, booker_id, start_time, end_time, status, payment_status, payment_id 
    FROM reservations 
    WHERE field_id = $1
    LIMIT $2 OFFSET $3`
	rows, err := repo.db.Query(query, fieldId, pageSize, offset)
	if err != nil {
		return nil, 0, apperror.NewInternalError(fmt.Sprintf("Error listing reservations by booker ID: %v", err))
	}
	defer rows.Close()

	var reservations []*reservation.Reservation
	for rows.Next() {
		var r reservation.Reservation
		details := &reservation.ReservationDetails{}
		if err := rows.Scan(&r.Id, &r.FieldId, &r.BookerId, &r.StartTime, &r.EndTime, &details.Status, &details.PaymentStatus, &details.PaymentID); err != nil {
			return nil, 0, apperror.NewInternalError(fmt.Sprintf("Error scanning reservation: %v", err))
		}
		r.Details = details
		reservations = append(reservations, &r)
	}

	countQuery := `SELECT COUNT(*) FROM reservations WHERE booker_id = $1`
	var totalCount int
	err = repo.db.QueryRow(countQuery, fieldId).Scan(&totalCount)
	if err != nil {
		return nil, 0, apperror.NewInternalError(fmt.Sprintf("Error counting reservations by booker ID: %v", err))
	}

	return reservations, totalCount, nil
}
