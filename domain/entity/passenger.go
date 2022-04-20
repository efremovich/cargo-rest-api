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

// Passenger represent schema of table sities.
type Passenger struct {
	UUID              string         `gorm:"size:36;not null;uniqueIndex;primary_key;" json:"uuid,omitempty"`
	FirstName         string         `gorm:"size:100;"                                 json:"first_name,omitempty"          from:"first_name"`
	LastName          string         `gorm:"size:100;"                                 json:"last_name,omitempty"           from:"last_name"`
	Patronomic        string         `gorm:"size:100;"                                 json:"patronomic,omitempty"          from:"patronomic"`
	BirthDay          time.Time      `gorm:"size:100;"                                 json:"birthday,omitempty"            from:"birthday"        time_format:"2006-01-02"`
	PassportSeries    string         `gorm:"size:4;"                                   json:"passport_series,omitempty"     from:"passport_series"`
	PassportNumber    string         `gorm:"size:6;"                                   json:"passport_number,omitempty"     from:"passport_number"`
	UserUUID          string         `gorm:"size:36"                                   json:"user_uuid,omitempty"`
	PassengerTypeUUID string         `gorm:"size:36"                                   json:"passenger_type_uuid,omitempty"`
	CreatedAt         time.Time      `                                                 json:"created_at,omitempty"`
	UpdatedAt         time.Time      `                                                 json:"updated_at,omitempty"`
	DeletedAt         gorm.DeletedAt `                                                 json:"deleted_at,omitempty"`
}

// PassengerFaker represent content when generate fake data of passenger_type.
type PassengerFaker struct {
	UUID              string    `faker:"uuid_hyphenated"`
	FirstName         string    `faker:"first_name"`
	LastName          string    `faker:"last_name"`
	Patronomic        string    `faker:"patronomic"`
	BirthDay          time.Time `faker:"birthday"`
	PassportSeries    string    `faker:"passport_series"`
	PassportNumber    string    `faker:"passport_number"`
	UserUUID          string    `faker:"user_uuid"`
	PassengerTypeUUID string    `faker:"passenger_type_uuid"`
}

// Passengers represent multiple Passenger.
type Passengers []*Passenger

// DetailPassenger represent format of detail Passenger.
type DetailPassenger struct {
	PassengerFieldsForDetail
}

// DetailPassengerList represent format of DetailPassenger for Passenger list.
type DetailPassengerList struct {
	PassengerFieldsForDetail
	PassengerFieldsForList
}

// PassengerFieldsForDetail represent fields of detail Passenger.
type PassengerFieldsForDetail struct {
	UUID              string    `json:"uuid"`
	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	Patronomic        string    `json:"patronomic"`
	BirthDay          time.Time `json:"birthday"            time_format:"2006-01-02"`
	PassportSeries    string    `json:"passport_series"`
	PassportNumber    string    `json:"passport_number"`
	UserUUID          string    `json:"user_uuid"`
	PassengerTypeUUID string    `json:"passenger_type_uuid"`
}

// PassengerFieldsForList represent fields of detail Passenger for Passenger list.
type PassengerFieldsForList struct {
	CreatedAt time.Time `json:"created_at"`
}

// TableName return name of table.
func (u *Passenger) TableName() string {
	return "passengers"
}

// FilterableFields return fields.
func (u *Passenger) FilterableFields() []interface{} {
	return []interface{}{
		"uuid",
		"passenger_type_uuid",
		"first_name",
		"last_name",
		"patronomic",
		"passport_series",
		"passport_number",
		"birthday",
	}
}

// Prepare will prepare submitted data of passenger_type.
func (u *Passenger) Prepare() {
	u.PassengerTypeUUID = html.EscapeString(strings.TrimSpace(u.PassengerTypeUUID))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

// BeforeCreate handle uuid generation and password hashing.
func (u *Passenger) BeforeCreate(tx *gorm.DB) error {
	generateUUID := uuid.New()
	if u.UUID == "" {
		u.UUID = generateUUID.String()
	}
	return nil
}

// DetailPassengers will return formatted passenger_type detail of multiple passenger_type.
func (sities Passengers) DetailPassengers() []interface{} {
	result := make([]interface{}, len(sities))
	for index, passenger_type := range sities {
		result[index] = passenger_type.DetailPassengerList()
	}
	return result
}

// DetailPassenger will return formatted passenger_type detail of passenger_type.
func (u *Passenger) DetailPassenger() interface{} {
	return &DetailPassenger{
		PassengerFieldsForDetail: PassengerFieldsForDetail{
			UUID:              u.UUID,
			FirstName:         u.FirstName,
			LastName:          u.LastName,
			Patronomic:        u.Patronomic,
			BirthDay:          u.BirthDay,
			PassportSeries:    u.PassportSeries,
			PassportNumber:    u.PassportNumber,
			UserUUID:          u.UserUUID,
			PassengerTypeUUID: u.PassengerTypeUUID,
		},
	}
}

// DetailPassengerList will return formatted passenger_type detail of passenger_type for passenger_type list.
func (u *Passenger) DetailPassengerList() interface{} {
	return &DetailPassengerList{
		PassengerFieldsForDetail: PassengerFieldsForDetail{
			UUID:              u.UUID,
			FirstName:         u.FirstName,
			LastName:          u.LastName,
			Patronomic:        u.Patronomic,
			BirthDay:          u.BirthDay,
			PassportSeries:    u.PassportSeries,
			PassportNumber:    u.PassportNumber,
			UserUUID:          u.UserUUID,
			PassengerTypeUUID: u.PassengerTypeUUID,
		},
		PassengerFieldsForList: PassengerFieldsForList{
			CreatedAt: u.CreatedAt,
		},
	}
}

// ValidateSavePassenger will validate create a new passenger_type request.
func (u *Passenger) ValidateSavePassenger() []response.ErrorForm {
	validation := validator.New()
	validation.
		Set("first_name", u.FirstName, validation.AddRule().IsAlphaUnicode().Required().Apply()).
		Set("last_name", u.LastName, validation.AddRule().IsAlphaUnicode().Required().Apply()).
		Set("patronomic", u.Patronomic, validation.AddRule().IsAlphaUnicode().Required().Apply()).
		Set("birthday", u.BirthDay, validation.AddRule().Required().Apply()).
		Set("passport_series", u.PassportSeries, validation.AddRule().IsDigit().Required().Apply()).
		Set("passport_number", u.PassportNumber, validation.AddRule().IsDigit().Required().Apply())
	return validation.Validate()
}

// ValidateUpdatePassenger will validate update a new passenger_type request.
func (u *Passenger) ValidateUpdatePassenger() []response.ErrorForm {
	validation := validator.New()
	validation.
		Set("first_name", u.FirstName, validation.AddRule().IsAlphaUnicode().Required().Apply()).
		Set("last_name", u.LastName, validation.AddRule().IsAlphaUnicode().Required().Apply()).
		Set("patronomic", u.Patronomic, validation.AddRule().IsAlphaUnicode().Required().Apply()).
		Set("birthday", u.BirthDay, validation.AddRule().Required().Apply()).
		Set("passport_series", u.PassportSeries, validation.AddRule().IsDigit().Required().Apply()).
		Set("passport_number", u.PassportNumber, validation.AddRule().IsDigit().Required().Apply())
	return validation.Validate()
}
