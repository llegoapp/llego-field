package field

import "time"

/* Field represents a soccer field with reservation details.*/
type Field struct {
	Id        int
	OwnerId   int
	Location  Address
	Status    FieldStatus
	OpenTime  time.Time
	CloseTime time.Time
}

func New(location Address, isReserved bool, ownerId int, openTime time.Time, closeTime time.Time) *Field {
	return &Field{
		OwnerId:   ownerId,
		Location:  location,
		Status:    StatusAvailable,
		OpenTime:  openTime,
		CloseTime: closeTime,
	}
}

/** Address represents a geographical address.*/
type Address struct {
	Street  Street
	City    City
	Country Country
}

func NewAddress(street Street, city City, country Country) *Address {
	return &Address{
		Street:  street,
		City:    city,
		Country: country,
	}
}

/*Street represents the street part of an address.*/
type Street string

/* City represents the city part of an address.*/
type City string

/* Country represents the country part of an address.*/
type Country string

/* FieldStatus represents the possible states of a field.*/
type FieldStatus string

const (
	StatusAvailable   FieldStatus = "available"
	StatusClosed      FieldStatus = "closed"
	StatusMaintenance FieldStatus = "maintenance"
)

func (f FieldStatus) IsAvalible() bool {
	return f == StatusAvailable
}
