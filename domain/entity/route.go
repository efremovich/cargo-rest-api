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

// Route represent schema of table route.
type Route struct {
	UUID string `json:"uuid,omitempty" gorm:"size:36;not null;uniqueIndex;primary_key;"`

	SityFrom Sity   `json:"sity_from"      gorm:"foreignKey:FromUUID"`
	SityTo   Sity   `json:"sity_to"        gorm:"foreignKey:ToUUID"`
	FromUUID string `json:"from,omitempty"                            form:"from"`
	ToUUID   string `json:"to,omitempty"                              form:"to"`

	Distance     int       `json:"distance,omitempty"      form:"distance"`
	DistanceTime time.Time `json:"distance_time,omitempty" form:"distance_time"`

	Prices []*Price `json:"prices" gorm:"many2many:route_prices;"`

	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt
}

// RouteFaker represent content when generate fake data of route.
type RouteFaker struct {
	UUID         string    `faker:"uid_hyphenated"`
	FromUUID     string    `faker:"from_uuid"`
	ToUUID       string    `faker:"to_uuid"`
	Distance     int       `faker:"distance"`
	DistanceTime time.Time `faker:"distance_time"`
	Prices       []*Price  `faker:"prices"`
}

// Routes represent multiple Route.
type Routes []*Route

// DetailRoute represent format of detail Route.
type DetailRoute struct {
	RouteFieldsForDetail
	Prices []interface{} `json:"prices,omitempty`
}

// DetailRouteList represent format of DetailRoute for Route list.
type DetailRouteList struct {
	RouteFieldsForDetail
	RouteFieldsForList
}

// RouteFieldsForDetail represent fields of detail Route.
type RouteFieldsForDetail struct {
	UUID         string    `json:"uuid"`
	FromUUID     string    `json:"from_uuid"`
	ToUUID       string    `json:"to_uuid"`
	Distance     int       `json:"distance"`
	DistanceTime time.Time `json:"distance_time"`
}

// RouteFieldsForList represent fields of detail Route for Route list.
type RouteFieldsForList struct {
	CreatedAt time.Time `json:"created_at"`
}

// TableName return name of table.
func (u *Route) TableName() string {
	return "routes"
}

// FilterableFields return fields.
func (u *Route) FilterableFields() []interface{} {
	return []interface{}{"uuid", "from_uuid", "to_uuid", "distance", "distance_time"}
}

// Prepare will prepare submitted data of route.
func (u *Route) Prepare() {
	u.UUID = html.EscapeString(strings.TrimSpace(u.UUID))
	u.FromUUID = html.EscapeString(strings.TrimSpace(u.FromUUID))
	u.ToUUID = html.EscapeString(strings.TrimSpace(u.ToUUID))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

// BeforeCreate handle uuid generation and password hashing.
func (u *Route) BeforeCreate(tx *gorm.DB) error {
	generateUUID := uuid.New()
	if u.UUID == "" {
		u.UUID = generateUUID.String()
	}
	return nil
}

// DetailRoutes will return formatted route detail of multiple route.
func (route Routes) DetailRoutes() []interface{} {
	result := make([]interface{}, len(route))
	for index, route := range route {
		result[index] = route.DetailRouteList()
	}
	return result
}

// DetailRoute will return formatted route detail of route.
func (u *Route) DetailRoute() interface{} {
	return &DetailRoute{
		RouteFieldsForDetail: RouteFieldsForDetail{
			UUID:         u.UUID,
			FromUUID:     u.FromUUID,
			ToUUID:       u.ToUUID,
			Distance:     u.Distance,
			DistanceTime: u.DistanceTime,
		},
	}
}

// DetailRouteList will return formatted route detail of route for route list.
func (u *Route) DetailRouteList() interface{} {
	return &DetailRouteList{
		RouteFieldsForDetail: RouteFieldsForDetail{
			UUID:         u.UUID,
			FromUUID:     u.FromUUID,
			ToUUID:       u.ToUUID,
			Distance:     u.Distance,
			DistanceTime: u.DistanceTime,
		},
		RouteFieldsForList: RouteFieldsForList{
			CreatedAt: u.CreatedAt,
		},
	}
}

// ValidateSaveRoute will validate create a new route request.
func (u *Route) ValidateSaveRoute() []response.ErrorForm {
	validation := validator.New()
	validation.
		Set(
			"from_uuid",
			u.FromUUID,
			validation.AddRule().Required().IsAlphaNumericSpaceAndSpecialCharacter().Length(3, 64).Apply(),
		).
		Set(
			"to_uuid",
			u.ToUUID,
			validation.AddRule().Required().IsAlphaNumericSpaceAndSpecialCharacter().Length(3, 64).Apply(),
		)
	return validation.Validate()
}

// ValidateUpdateRoute will validate update a new route request.
func (u *Route) ValidateUpdateRoute() []response.ErrorForm {
	validation := validator.New()
	validation.
		Set(
			"from_uuid",
			u.FromUUID,
			validation.AddRule().Required().IsAlphaNumericSpaceAndSpecialCharacter().Length(3, 64).Apply(),
		).
		Set(
			"to_uuid",
			u.ToUUID,
			validation.AddRule().Required().IsAlphaNumericSpaceAndSpecialCharacter().Length(3, 64).Apply(),
		)
	return validation.Validate()
}
