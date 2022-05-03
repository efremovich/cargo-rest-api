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

// Payment represent schema of table payment.
type Payment struct {
	UUID string `json:"uuid,omitempty" gorm:"size:36;not null;uniqueIndex;primary_key;"`

	PaymentDate time.Time `json:"payment_date"`

	Amount   float64  `json:"amount"`
	UserUUID string   `json:"user_uuid"`
	User     User     `json:"user"      foreignKey:"UserUUID"`
	Orders   []*Order `json:"orders"                          gorm:"many2many:payment_orders;"`

	TripUUID string `json:"trip_uuid"`
	Trip     Trip   `json:"trip"      gorm:"foreignKey:TripUUID"`

	ExternalUUID string `json:"external_uuid"`

	CreatedAt time.Time      `json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
type PaymentFaker struct {
	UUID         string    `faker:"uid_hyphenated" json:"uuid"`
	PaymentDate  time.Time `faker:"payment_date"`
	Amount       float64   `faker:"amount"`
	TripUUID     string    `faker:"trip_uuid"`
	ExternalUUID string    `faker:"external_uuid"  json:"external_uuid"`
	UserUUID     string    `faker:"user_uuid"`
}

// Payments represent multiple Payment.
type Payments []*Payment

// DetailPayment represent format of detail Payment.
type DetailPayment struct {
	PaymentFieldsForDetail
}

// DetailPaymentList represent format of DetailPayment for Payment list.
type DetailPaymentList struct {
	PaymentFieldsForDetail
	PaymentFieldsForList
}

// PaymentFieldsForDetail represent fields of detail Payment.
type PaymentFieldsForDetail struct {
	UUID        string    `json:"uuid"`
	PaymentDate time.Time `json:"payment_date"`
	Amount      float64   `json:"amount"`
	UserUUID    string    `json:"user_uuid"`
	TripUUID    string    `json:"trip_uuid"`

	ExternalUUID string `json:"external_uuid"`
}

// PaymentFieldsForList represent fields of detail Payment for Payment list.
type PaymentFieldsForList struct {
	CreatedAt time.Time `json:"created_at"`
}

// TableName return name of table.
func (u *Payment) TableName() string {
	return "payments"
}

// FilterableFields return fields.
func (u *Payment) FilterableFields() []interface{} {
	return []interface{}{
		"uuid",
		"payment_date",
		"amount",
		"user_uuid",
		"trip_uuid",
		"external_uuid",
	}
}

// Prepare will prepare submitted data of payment.
func (u *Payment) Prepare() {
	u.UUID = html.EscapeString(strings.TrimSpace(u.UUID))
	u.TripUUID = html.EscapeString(strings.TrimSpace(u.TripUUID))
	u.ExternalUUID = html.EscapeString(strings.TrimSpace(u.ExternalUUID))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

// BeforeCreate handle uuid generation and password hashing.
func (u *Payment) BeforeCreate(tx *gorm.DB) error {
	generateUUID := uuid.New()
	if u.UUID == "" {
		u.UUID = generateUUID.String()
	}
	return nil
}

// DetailPayments will return formatted payment detail of multiple payment.
func (payment Payments) DetailPayments() []interface{} {
	result := make([]interface{}, len(payment))
	for index, payment := range payment {
		result[index] = payment.DetailPaymentList()
	}
	return result
}

// DetailPayment will return formatted payment detail of payment.
func (u *Payment) DetailPayment() interface{} {
	return &DetailPayment{
		PaymentFieldsForDetail: PaymentFieldsForDetail{
			UUID:         u.UUID,
			PaymentDate:  u.PaymentDate,
			Amount:       u.Amount,
			UserUUID:     u.UserUUID,
			TripUUID:     u.TripUUID,
			ExternalUUID: u.ExternalUUID,
		},
	}
}

// DetailPaymentList will return formatted payment detail of payment for payment list.
func (u *Payment) DetailPaymentList() interface{} {
	return &DetailPaymentList{
		PaymentFieldsForDetail: PaymentFieldsForDetail{
			UUID:         u.UUID,
			PaymentDate:  u.PaymentDate,
			Amount:       u.Amount,
			UserUUID:     u.UserUUID,
			TripUUID:     u.TripUUID,
			ExternalUUID: u.ExternalUUID,
		},
		PaymentFieldsForList: PaymentFieldsForList{
			CreatedAt: u.CreatedAt,
		},
	}
}

// ValidateSavePayment will validate create a new payment request.
func (u *Payment) ValidateSavePayment() []response.ErrorForm {
	validation := validator.New()
	// validation.
	// 	Set(
	// 		"from",
	// 		u.FromUUID,
	// 		validation.AddRule().Required().IsAlphaNumericSpaceAndSpecialCharacter().Length(3, 64).Apply(),
	// 	)
	return validation.Validate()
}

// ValidateUpdatePayment will validate update a new payment request.
func (u *Payment) ValidateUpdatePayment() []response.ErrorForm {
	validation := validator.New()
	// validation.
	// 	Set(
	// 		"from",
	// 		u.FromUUID,
	// 		validation.AddRule().Required().IsAlphaNumericSpaceAndSpecialCharacter().Length(3, 64).Apply(),
	// 	)
	return validation.Validate()
}
