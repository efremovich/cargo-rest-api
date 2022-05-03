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

// Trip represent schema of table trip.
type Trip struct {
	UUID string `json:"uuid,omitempty" gorm:"size:36;not null;uniqueIndex;primary_key;"`

	RouteUUID          string         `json:"route_uuid"`
	Route              Route          `json:"route"                gorm:"foreignKey:RouteUUID"`
	VehicleUUID        string         `json:"vehicle_uuid"`
	Vehicle            Vehicle        `json:"vehicle"              gorm:"foreignKey:VehicleUUID"`
	DepartureTime      time.Time      `json:"departure_time"                                            from:"departure_time"`
	ArravialTive       time.Time      `json:"arravial_tive"                                             from:"arravial_tive"`
	RegularityTypeUUID string         `json:"regularity_type_uuid"`
	RegularityType     RegularityType `json:"regularity_type"      gorm:"foreignKey:RegularityTypeUUID"`
	DriverUUID         string         `json:"driver_uuid"`
	Driver             Driver         `json:"driver"               gorm:"foreignKey:DriverUUID"`

	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt
}

// TripFaker represent content when generate fake data of trip.
type TripFaker struct {
	UUID               string    `faker:"uid_hyphenated"`
	RouteUUID          string    `faker:"route_uuid"`
	VehicleUUID        string    `faker:"vehicle_uuid"`
	DepartureTime      time.Time `faker:"departure_time"`
	ArravialTive       time.Time `faker:"arravial_tive"`
	RegularityTypeUUID string    `faker:"regularity_type_uuid"`
	DriverUUID         string    `faker:"driver_uuid"`
}

// Trips represent multiple Trip.
type Trips []*Trip

// DetailTrip represent format of detail Trip.
type DetailTrip struct {
	TripFieldsForDetail
}

// DetailTripList represent format of DetailTrip for Trip list.
type DetailTripList struct {
	TripFieldsForDetail
	TripFieldsForList
}

// TripFieldsForDetail represent fields of detail Trip.
type TripFieldsForDetail struct {
	UUID string `json:"uuid"`

	RouteUUID          string    `json:"route_uuid"`
	VehicleUUID        string    `json:"vehicle_uuid"`
	DepartureTime      time.Time `json:"departure_time"`
	ArravialTive       time.Time `json:"arravial_tive"`
	RegularityTypeUUID string    `json:"regularity_type_uuid"`
	DriverUUID         string    `json:"driver_uuid"`
}

// TripFieldsForList represent fields of detail Trip for Trip list.
type TripFieldsForList struct {
	CreatedAt time.Time `json:"created_at"`
}

// TableName return name of table.
func (u *Trip) TableName() string {
	return "trips"
}

// FilterableFields return fields.
func (u *Trip) FilterableFields() []interface{} {
	return []interface{}{
		"uuid",
		"route_uuid",
		"vehicle_uuid",
		"departure_time",
		"arravial_tive",
		"regularity_type_uuid",
		"driver_uuid",
	}
}

// Prepare will prepare submitted data of trip.
func (u *Trip) Prepare() {
	u.UUID = html.EscapeString(strings.TrimSpace(u.UUID))
	u.RouteUUID = html.EscapeString(strings.TrimSpace(u.RouteUUID))
	u.VehicleUUID = html.EscapeString(strings.TrimSpace(u.VehicleUUID))
	u.RegularityTypeUUID = html.EscapeString(strings.TrimSpace(u.RegularityTypeUUID))
	u.DriverUUID = html.EscapeString(strings.TrimSpace(u.DriverUUID))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

// BeforeCreate handle uuid generation and password hashing.
func (u *Trip) BeforeCreate(tx *gorm.DB) error {
	generateUUID := uuid.New()
	if u.UUID == "" {
		u.UUID = generateUUID.String()
	}
	return nil
}

// DetailTrips will return formatted trip detail of multiple trip.
func (trip Trips) DetailTrips() []interface{} {
	result := make([]interface{}, len(trip))
	for index, trip := range trip {
		result[index] = trip.DetailTripList()
	}
	return result
}

// DetailTrip will return formatted trip detail of trip.
func (u *Trip) DetailTrip() interface{} {
	return &DetailTrip{
		TripFieldsForDetail: TripFieldsForDetail{
			UUID:               u.UUID,
			DepartureTime:      u.DepartureTime,
			ArravialTive:       u.ArravialTive,
			RouteUUID:          u.RouteUUID,
			VehicleUUID:        u.VehicleUUID,
			RegularityTypeUUID: u.RegularityTypeUUID,
			DriverUUID:         u.DriverUUID,
		},
	}
}

// DetailTripList will return formatted trip detail of trip for trip list.
func (u *Trip) DetailTripList() interface{} {
	return &DetailTripList{
		TripFieldsForDetail: TripFieldsForDetail{
			UUID:               u.UUID,
			DepartureTime:      u.DepartureTime,
			ArravialTive:       u.ArravialTive,
			RouteUUID:          u.RouteUUID,
			VehicleUUID:        u.VehicleUUID,
			RegularityTypeUUID: u.RegularityTypeUUID,
			DriverUUID:         u.DriverUUID,
		},
		TripFieldsForList: TripFieldsForList{
			CreatedAt: u.CreatedAt,
		},
	}
}

// ValidateSaveTrip will validate create a new trip request.
func (u *Trip) ValidateSaveTrip() []response.ErrorForm {
	validation := validator.New()
	// validation.
	// 	Set(
	// 		"from",
	// 		u.FromUUID,
	// 		validation.AddRule().Required().IsAlphaNumericSpaceAndSpecialCharacter().Length(3, 64).Apply(),
	// 	)
	return validation.Validate()
}

// ValidateUpdateTrip will validate update a new trip request.
func (u *Trip) ValidateUpdateTrip() []response.ErrorForm {
	validation := validator.New()
	// validation.
	// 	Set(
	// 		"from",
	// 		u.FromUUID,
	// 		validation.AddRule().Required().IsAlphaNumericSpaceAndSpecialCharacter().Length(3, 64).Apply(),
	// 	)
	return validation.Validate()
}
