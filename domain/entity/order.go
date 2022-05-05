package entity

import (
	"cargo-rest-api/pkg/response"
	"cargo-rest-api/pkg/validator"
	"html"
	"strings"
	"time"

	"gorm.io/gorm"

	"github.com/google/uuid"
)

// Order represent schema of table order.
type Order struct {
	UUID       string          `json:"uuid,omitempty" gorm:"size:36;not null;uniqueIndex;primary_key;"`
	OrderDate  time.Time       `json:"order_date"`
	Passengers []*Passenger    `json:"passengers"     gorm:"many2many:order_passengers;"`
	TripUUID   string          `json:"trip_uuid"`
	Trip       Trip            `json:"trip"           gorm:"foreignKey:TripUUID"`
	Seat       string          `json:"seat"`
	StatusUUID string          `json:"status_uuid"`
	Status     OrderStatusType `json:"status"         gorm:"foreignKey:StatusUUID"`

	ExternalUUID string `json:"external_uuid"`

	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt
}

// OrderFaker represent content when generate fake data of order.
type OrderFaker struct {
	UUID       string    `faker:"uid_hyphenated" json:"uuid"`
	OrderDate  time.Time `faker:"order_date"`
	TripUUID   string    `faker:"trip_uuid"`
	Seat       string    `faker:"seat"`
	StatusUUID string    `faker:"status_uuid"`

	ExternalUUID string `faker:"external_uuid"`
}

// Orders represent multiple Order.
type Orders []*Order

// DetailOrder represent format of detail Order.
type DetailOrder struct {
	OrderFieldsForDetail
	Passengers []interface{} `json:"passengers,omitempty"`
	Status     interface{}   `json:"status,omitempty"`
	Payment    interface{}   `json:"payment,omitempty"`
	Trip       interface{}   `json:"trip,omitempty"`
}

// DetailOrderList represent format of DetailOrder for Order list.
type DetailOrderList struct {
	OrderFieldsForDetail
	OrderFieldsForList
}

// OrderFieldsForDetail represent fields of detail Order.
type OrderFieldsForDetail struct {
	UUID string `json:"uuid"`

	OrderDate  time.Time `json:"order_date"`
	TripUUID   string    `json:"trip_uuid,omitempty"`
	Seat       string    `json:"seat"`
	StatusUUID string    `json:"status_uuid,omitempty"`

	ExternalUUID string `json:"external_uuid,omitempty"`
}

// OrderFieldsForList represent fields of detail Order for Order list.
type OrderFieldsForList struct {
	CreatedAt time.Time `json:"created_at"`
}

// TableName return name of table.
func (u *Order) TableName() string {
	return "orders"
}

// FilterableFields return fields.
func (u *Order) FilterableFields() []interface{} {
	return []interface{}{
		"uuid",
		"order_date",
		"trip_uuid",
		"seat",
		"status_uuid",
		"external_uuid",
	}
}

// Prepare will prepare submitted data of order.
func (u *Order) Prepare() {
	u.UUID = html.EscapeString(strings.TrimSpace(u.UUID))
	u.TripUUID = html.EscapeString(strings.TrimSpace(u.TripUUID))
	u.Seat = html.EscapeString(strings.TrimSpace(u.Seat))
	u.ExternalUUID = html.EscapeString(strings.TrimSpace(u.ExternalUUID))
	u.StatusUUID = html.EscapeString(strings.TrimSpace(u.StatusUUID))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

// BeforeCreate handle uuid generation and password hashing.
func (u *Order) BeforeCreate(tx *gorm.DB) error {
	generateUUID := uuid.New()
	if u.UUID == "" {
		u.UUID = generateUUID.String()
	}
	return nil
}

// DetailOrders will return formatted order detail of multiple order.
func (order Orders) DetailOrders() []interface{} {
	result := make([]interface{}, len(order))
	for index, order := range order {
		result[index] = order.DetailOrderList()
	}
	return result
}

// DetailOrder will return formatted order detail of order.
func (u *Order) DetailOrder() interface{} {
	return &DetailOrder{
		OrderFieldsForDetail: OrderFieldsForDetail{
			UUID:         u.UUID,
			OrderDate:    u.OrderDate,
			Seat:         u.Seat,
			ExternalUUID: u.ExternalUUID,
		},
		Passengers: Passengers.DetailPassengers(u.Passengers),
		Status:     u.Status.DetailOrderStatusType(),
		Trip:       u.Trip.DetailTrip(),
	}
}

// DetailOrderList will return formatted order detail of order for order list.
func (u *Order) DetailOrderList() interface{} {
	return &DetailOrderList{
		OrderFieldsForDetail: OrderFieldsForDetail{
			UUID:         u.UUID,
			OrderDate:    u.OrderDate,
			TripUUID:     u.TripUUID,
			Seat:         u.Seat,
			StatusUUID:   u.StatusUUID,
			ExternalUUID: u.ExternalUUID,
		},
		OrderFieldsForList: OrderFieldsForList{
			CreatedAt: u.CreatedAt,
		},
	}
}

// ValidateSaveOrder will validate create a new order request.
func (u *Order) ValidateSaveOrder() []response.ErrorForm {
	validation := validator.New()
	// validation.
	// 	Set(
	// 		"from",
	// 		u.FromUUID,
	// 		validation.AddRule().Required().IsAlphaNumericSpaceAndSpecialCharacter().Length(3, 64).Apply(),
	// 	)
	return validation.Validate()
}

// ValidateUpdateOrder will validate update a new order request.
func (u *Order) ValidateUpdateOrder() []response.ErrorForm {
	validation := validator.New()
	// validation.
	// 	Set(
	// 		"from",
	// 		u.FromUUID,
	// 		validation.AddRule().Required().IsAlphaNumericSpaceAndSpecialCharacter().Length(3, 64).Apply(),
	// 	)
	return validation.Validate()
}
