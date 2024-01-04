package reservation

import "time"

type Reservation struct {
	Id        int
	FieldId   int
	BookerId  int       // ID of the user who booked the reservation
	StartTime time.Time // Start time of the reservation
	EndTime   time.Time // End time of the reservation
	// Duration can be calculated as EndTime.Sub(StartTime)
	Reservation *ReservationDetails
}

func New(fieldId, bookerId int, startTime, endTime time.Time) Reservation {
	return Reservation{
		FieldId:   fieldId,
		BookerId:  bookerId,
		StartTime: startTime,
		EndTime:   endTime,
	}
}

/* ReservationDetails holds details about a field reservation.*/
type ReservationDetails struct {
	Status        ReservationStatus
	PaymentStatus *PaymentStatus
	PaymentID     string
}

/* ReservationStatus represents the possible states of a field reservation.*/
type ReservationStatus string

const (
	StatusPending   ReservationStatus = "pending"
	StatusReserved  ReservationStatus = "reserved"
	StatusCancelled ReservationStatus = "cancelled"
)

/* PaymentStatus represents the possible states of payment for a reservation.*/
type PaymentStatus string

const (
	PaymentUnpaid  PaymentStatus = "unpaid"
	PaymentPending PaymentStatus = "pending"
	PaymentPaid    PaymentStatus = "paid"
	PaymentFailed  PaymentStatus = "failed"
)
