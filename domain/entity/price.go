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

// Price represent schema of table sities.
type Price struct {
	UUID              string         `gorm:"size:36;not null;uniqueIndex;primary_key;" json:"uuid,omitempty"`
	PassengerTypeUUID string         `gorm:"size:36;not null;" json:"passenger_type_uuid,omitempty" form:"passenger_type_uuid"`
	PassengerType     PassengerType  `gorm:"foreignKey:PassengerTypeUUID" json:"passenger_type,omitempty"`
	Price             float64        `json:"price,omitempty" from:"price"`
	CreatedAt         time.Time      `json:"created_at,omitempty"`
	UpdatedAt         time.Time      `json:"updated_at,omitempty"`
	DeletedAt         gorm.DeletedAt `json:"deleted_at,omitempty"`
}

// PriceFaker represent content when generate fake data of passenger_type.
type PriceFaker struct {
	UUID              string  `faker:"uuid_hyphenated"`
	PassengerTypeUUID string  `faker:"passenger_type_uuid"`
	Price             float64 `faker:"price"`
}

// Prices represent multiple Price.
type Prices []*Price

// DetailPrice represent format of detail Price.
type DetailPrice struct {
	PriceFieldsForDetail
}

// DetailPriceList represent format of DetailPrice for Price list.
type DetailPriceList struct {
	PriceFieldsForDetail
	PriceFieldsForList
}

// PriceFieldsForDetail represent fields of detail Price.
type PriceFieldsForDetail struct {
	UUID              string      `json:"uuid"`
	PassengerTypeUUID string      `json:"passenger_type_uuid"`
	PassengerType     interface{} `json:"passenger_type"`
	Price             float64     `json:"price"`
}

// PriceFieldsForList represent fields of detail Price for Price list.
type PriceFieldsForList struct {
	CreatedAt time.Time `json:"created_at"`
}

// TableName return name of table.
func (u *Price) TableName() string {
	return "prices"
}

// FilterableFields return fields.
func (u *Price) FilterableFields() []interface{} {
	return []interface{}{"uuid", "passenger_type_uuid", "price"}
}

// Prepare will prepare submitted data of passenger_type.
func (u *Price) Prepare() {
	u.PassengerTypeUUID = html.EscapeString(strings.TrimSpace(u.PassengerTypeUUID))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

// BeforeCreate handle uuid generation and password hashing.
func (u *Price) BeforeCreate(tx *gorm.DB) error {
	generateUUID := uuid.New()
	if u.UUID == "" {
		u.UUID = generateUUID.String()
	}
	return nil
}

// DetailPrices will return formatted passenger_type detail of multiple passenger_type.
func (prices Prices) DetailPrices() []interface{} {
	result := make([]interface{}, len(prices))
	for index, price := range prices {
		result[index] = price.DetailPriceList()
	}
	return result
}

// DetailPrice will return formatted passenger_type detail of passenger_type.
func (u *Price) DetailPrice() interface{} {
	return &DetailPrice{
		PriceFieldsForDetail: PriceFieldsForDetail{
			UUID:              u.UUID,
			PassengerTypeUUID: u.PassengerTypeUUID,
			PassengerType:     u.PassengerType.Type,
			Price:             u.Price,
		},
	}
}

// GetPriceTypeDetail will return .
func (u *Price) GetPricePassengerType() interface{} {
	pst := PassengerType{UUID: u.UUID}
	a := pst.DetailPassengerTypeList()
	return a
}

// DetailPriceList will return formatted passenger_type detail of passenger_type for passenger_type list.
func (u *Price) DetailPriceList() interface{} {
	return &DetailPriceList{
		PriceFieldsForDetail: PriceFieldsForDetail{
			UUID:              u.UUID,
			PassengerTypeUUID: u.PassengerTypeUUID,
			Price:             u.Price,
		},
		PriceFieldsForList: PriceFieldsForList{
			CreatedAt: u.CreatedAt,
		},
	}
}

// ValidateSavePrice will validate create a new passenger_type request.
func (u *Price) ValidateSavePrice() []response.ErrorForm {
	validation := validator.New()
	validation.
		Set("passenger_type_uuid", u.PassengerTypeUUID, validation.AddRule().Required().Length(3, 64).Apply()).
		Set("price", u.Price, validation.AddRule().Required().Apply())
	return validation.Validate()
}

// ValidateUpdatePrice will validate update a new passenger_type request.
func (u *Price) ValidateUpdatePrice() []response.ErrorForm {
	validation := validator.New()
	validation.
		Set("passenger_type_uuid", u.PassengerTypeUUID, validation.AddRule().Required().Length(3, 64).Apply()).
		Set("price", u.Price, validation.AddRule().Required().Apply())
	return validation.Validate()
}
