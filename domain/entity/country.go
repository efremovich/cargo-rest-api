package entity

import (
	"html"
	"strings"
	"time"

	"gorm.io/gorm"

	"github.com/google/uuid"
)

// Country represent schema of table countries.
type Country struct {
	UUID      string    `gorm:"size:36;not null;uniqueIndex;primary_key;" json:"uuid"`
	Name      string    `gorm:"size:100;not null;" json:"name" form:"name"`
	Latitude  string    `gorm:"size:100;" json:"latitude" form:"latitude"`
	Longitude string    `gorm:"size:100;" json:"longitude" form:"longitude"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt
}

// CountryFaker represent content when generate fake data of country.
type CountryFaker struct {
	UUID      string `faker:"uuid_hyphenated"`
	Name      string `faker:"name"`
	Latitude  string `faker:"latitude"`
	Longitude string `faker:"longitude"`
}

// Countries represent multiple Country.
type Countries []Country

// DetailCountry represent format of detail Country.
type DetailCountry struct {
	CountryFieldsForDetail
	Role []interface{} `json:"roles,omitempty"`
}

// DetailCountryList represent format of DetailCountry for Country list.
type DetailCountryList struct {
	CountryFieldsForDetail
	CountryFieldsForList
}

// CountryFieldsForDetail represent fields of detail Country.
type CountryFieldsForDetail struct {
	UUID      string      `json:"uuid"`
	Name      string      `json:"name"`
	Latitude  interface{} `json:"latitude"`
	Longitude interface{} `json:"longitude"`
}

// CountryFieldsForList represent fields of detail Country for Country list.
type CountryFieldsForList struct {
	CreatedAt time.Time `json:"created_at"`
}

// TableName return name of table.
func (u *Country) TableName() string {
	return "countries"
}

// FilterableFields return fields.
func (u *Country) FilterableFields() []interface{} {
	return []interface{}{"uuid", "name", "latitude", "longitude"}
}

// Prepare will prepare submitted data of country.
func (u *Country) Prepare() {
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
	u.Latitude = html.EscapeString(strings.TrimSpace(u.Latitude))
	u.Longitude = html.EscapeString(strings.TrimSpace(u.Longitude))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

// BeforeCreate handle uuid generation and password hashing.
func (u *Country) BeforeCreate(tx *gorm.DB) error {
	generateUUID := uuid.New()
	if u.UUID == "" {
		u.UUID = generateUUID.String()
	}
	return nil
}

// DetailCountries will return formatted country detail of multiple country.
func (countries Countries) DetailCountries() []interface{} {
	result := make([]interface{}, len(countries))
	for index, country := range countries {
		result[index] = country.DetailCountryList()
	}
	return result
}

// DetailCountry will return formatted country detail of country.
func (u *Country) DetailCountry() interface{} {
	return &DetailCountry{
		CountryFieldsForDetail: CountryFieldsForDetail{
			UUID:      u.UUID,
			Name:      u.Name,
			Latitude:  u.Latitude,
			Longitude: u.Longitude,
		},
	}
}

// DetailCountryList will return formatted country detail of country for country list.
func (u *Country) DetailCountryList() interface{} {
	return &DetailCountryList{
		CountryFieldsForDetail: CountryFieldsForDetail{
			UUID:      u.UUID,
			Name:      u.Name,
			Latitude:  u.Latitude,
			Longitude: u.Longitude,
		},
		CountryFieldsForList: CountryFieldsForList{
			CreatedAt: u.CreatedAt,
		},
	}
}
