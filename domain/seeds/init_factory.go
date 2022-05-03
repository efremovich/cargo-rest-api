package seeds

import (
	"cargo-rest-api/domain/entity"
	"time"

	"gorm.io/gorm"

	"github.com/google/uuid"
)

type InitFactory struct {
	seeders []Seed
}

var (
	user = &entity.User{
		UUID:     uuid.New().String(),
		Name:     "Aleksandr Efremov",
		Email:    "sasha.fima@gmail.com",
		Phone:    "+79384100025",
		Password: "020407",
	}

	role = &entity.Role{
		UUID: uuid.New().String(), Name: "Super Administrator",
	}
	otherRoles = []*entity.Role{
		{UUID: uuid.New().String(), Name: "Driver"},
		{UUID: uuid.New().String(), Name: "User"},
	}
	permissions = []*entity.Permission{
		{UUID: uuid.New().String(), ModuleKey: "user", PermissionKey: "read"},
		{UUID: uuid.New().String(), ModuleKey: "user", PermissionKey: "create"},
		{UUID: uuid.New().String(), ModuleKey: "user", PermissionKey: "update"},
		{UUID: uuid.New().String(), ModuleKey: "user", PermissionKey: "delete"},
		{UUID: uuid.New().String(), ModuleKey: "user", PermissionKey: "bulk_delete"},
		{UUID: uuid.New().String(), ModuleKey: "user", PermissionKey: "detail"},
		{UUID: uuid.New().String(), ModuleKey: "role", PermissionKey: "read"},
		{UUID: uuid.New().String(), ModuleKey: "role", PermissionKey: "create"},
		{UUID: uuid.New().String(), ModuleKey: "role", PermissionKey: "update"},
		{UUID: uuid.New().String(), ModuleKey: "role", PermissionKey: "delete"},
		{UUID: uuid.New().String(), ModuleKey: "role", PermissionKey: "bulk_delete"},
		{UUID: uuid.New().String(), ModuleKey: "role", PermissionKey: "detail"},
		{UUID: uuid.New().String(), ModuleKey: "tour", PermissionKey: "read"},
		{UUID: uuid.New().String(), ModuleKey: "tour", PermissionKey: "create"},
		{UUID: uuid.New().String(), ModuleKey: "tour", PermissionKey: "update"},
		{UUID: uuid.New().String(), ModuleKey: "tour", PermissionKey: "delete"},
		{UUID: uuid.New().String(), ModuleKey: "tour", PermissionKey: "bulk_delete"},
		{UUID: uuid.New().String(), ModuleKey: "tour", PermissionKey: "detail"},
		{UUID: uuid.New().String(), ModuleKey: "sity", PermissionKey: "read"},
		{UUID: uuid.New().String(), ModuleKey: "sity", PermissionKey: "create"},
		{UUID: uuid.New().String(), ModuleKey: "sity", PermissionKey: "update"},
		{UUID: uuid.New().String(), ModuleKey: "sity", PermissionKey: "delete"},
		{UUID: uuid.New().String(), ModuleKey: "sity", PermissionKey: "bulk_delete"},
		{UUID: uuid.New().String(), ModuleKey: "sity", PermissionKey: "detail"},
		{UUID: uuid.New().String(), ModuleKey: "vehicle", PermissionKey: "read"},
		{UUID: uuid.New().String(), ModuleKey: "vehicle", PermissionKey: "create"},
		{UUID: uuid.New().String(), ModuleKey: "vehicle", PermissionKey: "update"},
		{UUID: uuid.New().String(), ModuleKey: "vehicle", PermissionKey: "delete"},
		{UUID: uuid.New().String(), ModuleKey: "vehicle", PermissionKey: "bulk_delete"},
		{UUID: uuid.New().String(), ModuleKey: "vehicle", PermissionKey: "detail"},
		{UUID: uuid.New().String(), ModuleKey: "price", PermissionKey: "read"},
		{UUID: uuid.New().String(), ModuleKey: "price", PermissionKey: "create"},
		{UUID: uuid.New().String(), ModuleKey: "price", PermissionKey: "update"},
		{UUID: uuid.New().String(), ModuleKey: "price", PermissionKey: "delete"},
		{UUID: uuid.New().String(), ModuleKey: "price", PermissionKey: "bulk_delete"},
		{UUID: uuid.New().String(), ModuleKey: "price", PermissionKey: "detail"},
		{UUID: uuid.New().String(), ModuleKey: "passenger_type", PermissionKey: "read"},
		{UUID: uuid.New().String(), ModuleKey: "passenger_type", PermissionKey: "create"},
		{UUID: uuid.New().String(), ModuleKey: "passenger_type", PermissionKey: "update"},
		{UUID: uuid.New().String(), ModuleKey: "passenger_type", PermissionKey: "delete"},
		{UUID: uuid.New().String(), ModuleKey: "passenger_type", PermissionKey: "bulk_delete"},
		{UUID: uuid.New().String(), ModuleKey: "passenger_type", PermissionKey: "detail"},
		{UUID: uuid.New().String(), ModuleKey: "passenger", PermissionKey: "read"},
		{UUID: uuid.New().String(), ModuleKey: "passenger", PermissionKey: "create"},
		{UUID: uuid.New().String(), ModuleKey: "passenger", PermissionKey: "update"},
		{UUID: uuid.New().String(), ModuleKey: "passenger", PermissionKey: "delete"},
		{UUID: uuid.New().String(), ModuleKey: "passenger", PermissionKey: "bulk_delete"},
		{UUID: uuid.New().String(), ModuleKey: "passenger", PermissionKey: "detail"},
		{UUID: uuid.New().String(), ModuleKey: "document_type", PermissionKey: "read"},
		{UUID: uuid.New().String(), ModuleKey: "document_type", PermissionKey: "create"},
		{UUID: uuid.New().String(), ModuleKey: "document_type", PermissionKey: "update"},
		{UUID: uuid.New().String(), ModuleKey: "document_type", PermissionKey: "delete"},
		{UUID: uuid.New().String(), ModuleKey: "document_type", PermissionKey: "bulk_delete"},
		{UUID: uuid.New().String(), ModuleKey: "document_type", PermissionKey: "detail"},
		{UUID: uuid.New().String(), ModuleKey: "regularity_type", PermissionKey: "read"},
		{UUID: uuid.New().String(), ModuleKey: "regularity_type", PermissionKey: "create"},
		{UUID: uuid.New().String(), ModuleKey: "regularity_type", PermissionKey: "update"},
		{UUID: uuid.New().String(), ModuleKey: "regularity_type", PermissionKey: "delete"},
		{UUID: uuid.New().String(), ModuleKey: "regularity_type", PermissionKey: "bulk_delete"},
		{UUID: uuid.New().String(), ModuleKey: "regularity_type", PermissionKey: "detail"},
		{UUID: uuid.New().String(), ModuleKey: "driver", PermissionKey: "read"},
		{UUID: uuid.New().String(), ModuleKey: "driver", PermissionKey: "create"},
		{UUID: uuid.New().String(), ModuleKey: "driver", PermissionKey: "update"},
		{UUID: uuid.New().String(), ModuleKey: "driver", PermissionKey: "delete"},
		{UUID: uuid.New().String(), ModuleKey: "driver", PermissionKey: "bulk_delete"},
		{UUID: uuid.New().String(), ModuleKey: "driver", PermissionKey: "detail"},
		{UUID: uuid.New().String(), ModuleKey: "route", PermissionKey: "read"},
		{UUID: uuid.New().String(), ModuleKey: "route", PermissionKey: "create"},
		{UUID: uuid.New().String(), ModuleKey: "route", PermissionKey: "update"},
		{UUID: uuid.New().String(), ModuleKey: "route", PermissionKey: "delete"},
		{UUID: uuid.New().String(), ModuleKey: "route", PermissionKey: "bulk_delete"},
		{UUID: uuid.New().String(), ModuleKey: "route", PermissionKey: "detail"},
		{UUID: uuid.New().String(), ModuleKey: "trip", PermissionKey: "read"},
		{UUID: uuid.New().String(), ModuleKey: "trip", PermissionKey: "create"},
		{UUID: uuid.New().String(), ModuleKey: "trip", PermissionKey: "update"},
		{UUID: uuid.New().String(), ModuleKey: "trip", PermissionKey: "delete"},
		{UUID: uuid.New().String(), ModuleKey: "trip", PermissionKey: "bulk_delete"},
		{UUID: uuid.New().String(), ModuleKey: "trip", PermissionKey: "detail"},
		{UUID: uuid.New().String(), ModuleKey: "order", PermissionKey: "read"},
		{UUID: uuid.New().String(), ModuleKey: "order", PermissionKey: "create"},
		{UUID: uuid.New().String(), ModuleKey: "order", PermissionKey: "update"},
		{UUID: uuid.New().String(), ModuleKey: "order", PermissionKey: "delete"},
		{UUID: uuid.New().String(), ModuleKey: "order", PermissionKey: "bulk_delete"},
		{UUID: uuid.New().String(), ModuleKey: "order", PermissionKey: "detail"},
		{UUID: uuid.New().String(), ModuleKey: "payment", PermissionKey: "read"},
		{UUID: uuid.New().String(), ModuleKey: "payment", PermissionKey: "create"},
		{UUID: uuid.New().String(), ModuleKey: "payment", PermissionKey: "update"},
		{UUID: uuid.New().String(), ModuleKey: "payment", PermissionKey: "delete"},
		{UUID: uuid.New().String(), ModuleKey: "payment", PermissionKey: "bulk_delete"},
		{UUID: uuid.New().String(), ModuleKey: "payment", PermissionKey: "detail"},
	}
	userRole = &entity.UserRole{
		UUID:     uuid.New().String(),
		UserUUID: user.UUID,
		RoleUUID: role.UUID,
	}
	storageCategory = []*entity.StorageCategory{
		{
			UUID:      uuid.New().String(),
			Slug:      "avatar",
			Path:      "avatar",
			Name:      "Avatar",
			MimeTypes: "image/jpg,image/jpeg,image/png,image/bmp,image/gif",
		},
		{
			UUID:      uuid.New().String(),
			Slug:      "document",
			Path:      "document",
			Name:      "Document",
			MimeTypes: "application/pdf",
		},
		{
			UUID:      uuid.New().String(),
			Slug:      "file",
			Path:      "file",
			Name:      "File",
			MimeTypes: "application/pdf",
		},
		{
			UUID:      uuid.New().String(),
			Slug:      "thumbnail",
			Path:      "thumbnail",
			Name:      "Thumbnail",
			MimeTypes: "image/png",
		},
	}
	application = &entity.Application{
		UUID: uuid.New().String(),
		Name: "cargo-rest-api",
	}
	applicationApiKey = &entity.ApplicationApiKey{
		UUID:            uuid.New().String(),
		ApplicationUUID: application.UUID,
		Name:            "cargo-rest-api-api-key",
		ApiKey:          "9fde26c5-bc87-4081-a270-50b414d70fb6",
	}
	applicationOauth = &entity.ApplicationOauth{
		UUID:             uuid.New().String(),
		ApplicationUUID:  application.UUID,
		Name:             "cargo-rest-api-oauth",
		SupportEmails:    "support_one@example.com,support_two@example.com",
		DeveloperEmails:  "dev@example.com",
		Logo:             "",
		HomePageURL:      "https://github.com/efremovich/cargo-rest-api",
		TosURL:           "https://github.com/efremovich/cargo-rest-api",
		PrivacyPolicyURL: "https://github.com/efremovich/cargo-rest-api",
		Domains:          "http://localhost:8181",
		Scopes:           "read write",
	}
	applicationOauthClient = &entity.ApplicationOauthClient{
		UUID:            uuid.New().String(),
		ApplicationUUID: application.UUID,
		Name:            "cargo-rest-api-oauth-client",
		ClientID:        "10e207ab-79ec-42ed-85f2-3a10e3b3ddbb",
		ClientSecret:    "9e8c5bfe-a93e-4041-b404-0ae326a1e491",
		Referrers:       "http://localhost:8181",
		Callbacks:       "http://localhost:8181/oauth2/callback",
	}
	sities = []*entity.Sity{
		{
			UUID:      uuid.New().String(),
			Name:      "Волгоград",
			Region:    "Волгоградская область",
			Latitude:  "48.7194",
			Longitude: "44.5018",
		},
		{
			UUID:      uuid.New().String(),
			Name:      "Елань",
			Region:    "Волгоградская область",
			Latitude:  "50.5656",
			Longitude: "43.4416",
		},
		{
			UUID:      uuid.New().String(),
			Name:      "Сочи",
			Region:    "Краснодарский край",
			Latitude:  "43.3557",
			Longitude: "39.4332",
		},
		{
			UUID:      uuid.New().String(),
			Name:      "Краснодар",
			Region:    "Краснодарский край",
			Latitude:  "45.0241",
			Longitude: "38.5833",
		},
	}
	vehicles = []*entity.Vehicle{
		{UUID: uuid.New().String(), Model: "Ford transit", RegCode: "x245уы132", NumberOfSeats: 5, Class: "Комфорт +"},
		{UUID: uuid.New().String(), Model: "BMV Colt", RegCode: "й111шш132", NumberOfSeats: 7, Class: "Люкс"},
		{UUID: uuid.New().String(), Model: "Renaute Bibi", RegCode: "т777уз99", NumberOfSeats: 4, Class: "Эконом"},
		{UUID: uuid.New().String(), Model: "Газель", RegCode: "а666дд132", NumberOfSeats: 13, Class: "Бомж"},
	}
	passengerTypes = []*entity.PassengerType{
		{UUID: "04e9b29e-064b-4a13-8bab-074b14ae465d", Type: "Взрослый"},
		{UUID: "1c888dfd-78be-40ca-a85a-61cc3ab7fb1e", Type: "Детский"},
		{UUID: "7f3eb88e-98bd-4f5b-8a8c-34aaed1c7ffd", Type: "Пенсионный"},
	}
	prices = []*entity.Price{
		{
			UUID:              "c1dadc6c-76e0-4213-a669-140dd389bed2",
			PassengerTypeUUID: "7f3eb88e-98bd-4f5b-8a8c-34aaed1c7ffd",
			Price:             350.00,
		},
		{
			UUID:              "f2480365-2a2b-4e63-904f-03cb92ef06ec",
			PassengerTypeUUID: "1c888dfd-78be-40ca-a85a-61cc3ab7fb1e",
			Price:             250.00,
		},
		{
			UUID:              "17a77ff0-5dd4-42ed-8320-d9f204027dda",
			PassengerTypeUUID: "04e9b29e-064b-4a13-8bab-074b14ae465d",
			Price:             550.00,
		},
	}
	passengers = []*entity.Passenger{
		{
			UUID:              uuid.New().String(),
			FirstName:         "Владимир",
			LastName:          "Ульянов",
			Patronomic:        "Ильич",
			BirthDay:          time.Date(1870, time.April, 22, 0, 0, 0, 0, time.UTC),
			DocumentSeries:    "0401",
			DocumentNumber:    "564247",
			UserUUID:          user.UUID,
			PassengerTypeUUID: passengerTypes[2].UUID,
		},
		{
			UUID:              uuid.New().String(),
			FirstName:         "Николай",
			LastName:          "Чехидзе",
			Patronomic:        "Семёнович",
			BirthDay:          time.Date(1864, time.April, 9, 0, 0, 0, 0, time.UTC),
			DocumentSeries:    "5501",
			DocumentNumber:    "014247",
			UserUUID:          user.UUID,
			PassengerTypeUUID: passengerTypes[0].UUID,
		},
	}
	documentTypes = []*entity.DocumentType{
		{UUID: "04e9b29e-064b-4a13-8bab-074b14ae465d", Type: "Паспорт"},
		{UUID: "1c888dfd-78be-40ca-a85a-61cc3ab7fb1e", Type: "Свидетельство о рождении"},
		{UUID: "7f3eb88e-98bd-4f5b-8a8c-34aaed1c7ffd", Type: "Водительские парава"},
	}
	regularityTypes = []*entity.RegularityType{
		{UUID: "58e9b29e-064b-4a13-8bab-074b14ae465d", Type: "Каждый день"},
		{UUID: "1c888dfd-d8be-40ca-a85a-61cc3ab7fb1e", Type: "Каждый х день интервала (1 день недели или месяца)"},
		{UUID: "743eb88e-98bd-4f5b-8a8c-34aaed1c7ffd", Type: "В указанные даты"},
	}
	orderStatusTypes = []*entity.OrderStatusType{
		{UUID: "04e9be9e-064b-4a13-8bab-074b14ae465d", Type: "Оплачен"},
		{UUID: "1c888sfd-78ie-40ca-a85a-61cc3ab7fb1e", Type: "Не оплачен"},
		{UUID: "7f3ebl8e-98bd-4f5b-8a8c-34aaed1c7ffd", Type: "Отменен"},
	}
	drivers = []*entity.Driver{
		{
			UUID:     uuid.New().String(),
			Name:     "Мамука Тбилиский",
			UserUUID: user.UUID,
			Vehicles: []*entity.Vehicle{{UUID: vehicles[0].UUID}},
		},
		{UUID: uuid.New().String(), Name: "Гагик Анпский", UserUUID: user.UUID},
		{UUID: uuid.New().String(), Name: "Арам Хачитурян", UserUUID: user.UUID},
		{UUID: uuid.New().String(), Name: "Кирил Радионов", UserUUID: user.UUID},
	}
	routes = []*entity.Route{
		{
			UUID:         uuid.New().String(),
			FromUUID:     sities[0].UUID,
			ToUUID:       sities[1].UUID,
			Distance:     500000,
			DistanceTime: 500,
			Prices:       prices,
		},
		{
			UUID:         uuid.New().String(),
			FromUUID:     sities[1].UUID,
			ToUUID:       sities[1].UUID,
			Distance:     500000,
			DistanceTime: 500,
			Prices:       prices,
		},
		{
			UUID:         uuid.New().String(),
			FromUUID:     sities[2].UUID,
			ToUUID:       sities[3].UUID,
			Distance:     300000,
			DistanceTime: 360,
			Prices:       prices,
		},
		{
			UUID:         uuid.New().String(),
			FromUUID:     sities[3].UUID,
			ToUUID:       sities[2].UUID,
			Distance:     300000,
			DistanceTime: 360,
			Prices:       prices,
		},
	}
	trips = []*entity.Trip{
		{
			UUID:               uuid.New().String(),
			RouteUUID:          routes[0].UUID,
			VehicleUUID:        vehicles[0].UUID,
			DepartureTime:      time.Date(2022, time.May, 25, 11, 0, 0, 0, time.UTC),
			ArravialTive:       time.Date(2022, time.May, 25, 19, 30, 0, 0, time.UTC),
			RegularityTypeUUID: regularityTypes[0].UUID,
			DriverUUID:         drivers[0].UUID,
		},
		{
			UUID:               uuid.New().String(),
			RouteUUID:          routes[1].UUID,
			VehicleUUID:        vehicles[0].UUID,
			DepartureTime:      time.Date(2022, time.May, 26, 11, 0, 0, 0, time.UTC),
			ArravialTive:       time.Date(2022, time.May, 26, 19, 30, 0, 0, time.UTC),
			RegularityTypeUUID: regularityTypes[0].UUID,
			DriverUUID:         drivers[1].UUID,
		},
		{
			UUID:               uuid.New().String(),
			RouteUUID:          routes[2].UUID,
			VehicleUUID:        vehicles[1].UUID,
			DepartureTime:      time.Date(2022, time.May, 28, 11, 0, 0, 0, time.UTC),
			ArravialTive:       time.Date(2022, time.May, 28, 19, 30, 0, 0, time.UTC),
			RegularityTypeUUID: regularityTypes[0].UUID,
			DriverUUID:         drivers[2].UUID,
		},
		{
			UUID:               uuid.New().String(),
			RouteUUID:          routes[3].UUID,
			VehicleUUID:        vehicles[1].UUID,
			DepartureTime:      time.Date(2022, time.May, 28, 11, 0, 0, 0, time.UTC),
			ArravialTive:       time.Date(2022, time.May, 28, 19, 30, 0, 0, time.UTC),
			RegularityTypeUUID: regularityTypes[0].UUID,
			DriverUUID:         drivers[3].UUID,
		},
	}
	orders = []*entity.Order{
		{
			UUID:        uuid.New().String(),
			OrdrDate:    time.Date(2022, time.May, 9, 9, 0, 0, 0, time.UTC),
			PaymentDate: time.Date(2022, time.May, 9, 9, 0, 0, 0, time.UTC),
			TripUUID:    trips[0].UUID,
			Seat:        "Без места",
			StatusUUID:  orderStatusTypes[0].UUID, // оплачен
		},
		{
			UUID:       uuid.New().String(),
			OrdrDate:   time.Date(2022, time.May, 9, 9, 0, 0, 0, time.UTC),
			TripUUID:   trips[1].UUID,
			Seat:       "1 (У окна)",
			StatusUUID: orderStatusTypes[1].UUID, // не оплачен, бронирование
		},
	}
	payments = []*entity.Payment{
		{
			UUID:         uuid.New().String(),
			PaymentDate:  time.Now(),
			UserUUID:     user.UUID,
			TripUUID:     trips[0].UUID,
			Orders:       orders[:0],
			ExternalUUID: uuid.New().String(),
		},
		{
			UUID:         uuid.New().String(),
			PaymentDate:  time.Now(),
			UserUUID:     user.UUID,
			TripUUID:     trips[0].UUID,
			Orders:       orders[:1],
			ExternalUUID: uuid.New().String(),
		},
	}
)

func newInitFactory() *InitFactory {
	return &InitFactory{seeders: make([]Seed, 0)}
}

func (is *InitFactory) generateUserSeeder() *InitFactory {
	is.seeders = append(is.seeders, Seed{
		Name: "Create initial user",
		Run: func(db *gorm.DB) error {
			_, errDB := createUser(db, user)
			return errDB
		},
	})

	return is
}

func (is *InitFactory) generateRoleSeeder() *InitFactory {
	is.seeders = append(is.seeders, Seed{
		Name: "Create initial role",
		Run: func(db *gorm.DB) error {
			_, errDB := createRole(db, role)
			return errDB
		},
	})

	return is
}

func (is *InitFactory) generateOtherRoleSeeder() *InitFactory {
	for _, role := range otherRoles {
		or := role
		is.seeders = append(is.seeders, Seed{
			Name: "Create initial other roles",
			Run: func(db *gorm.DB) error {
				_, errDB := createRole(db, or)
				return errDB
			},
		})
	}
	return is
}

func (is *InitFactory) generatePermissionsSeeder() *InitFactory {
	for _, p := range permissions {
		cp := p
		is.seeders = append(is.seeders, Seed{
			Name: "Create initial permission",
			Run: func(db *gorm.DB) error {
				_, errDB := createPermission(db, cp)
				return errDB
			},
		})
	}

	return is
}

func (is *InitFactory) generateRolePermissionsSeeder() *InitFactory {
	r := role
	for _, p := range permissions {
		csp := p
		crp := &entity.RolePermission{
			UUID: uuid.New().String(),
		}

		is.seeders = append(is.seeders, Seed{
			Name: "Create initial permission",
			Run: func(db *gorm.DB) error {
				_, errDB := createRolePermission(db, r, csp, crp)
				return errDB
			},
		})
	}

	return is
}

func (is *InitFactory) generateUserRoleSeeder() *InitFactory {
	is.seeders = append(is.seeders, Seed{
		Name: "Assign initial role to user",
		Run: func(db *gorm.DB) error {
			_, errDB := createUserRole(db, user, role, userRole)
			return errDB
		},
	})

	return is
}

func (is *InitFactory) generateStorageCategorySeeder() *InitFactory {
	for _, sc := range storageCategory {
		csc := sc
		is.seeders = append(is.seeders, Seed{
			Name: "Create initial storage category",
			Run: func(db *gorm.DB) error {
				_, errDB := createStorageCategory(db, csc)
				return errDB
			},
		})
	}

	return is
}

func (is *InitFactory) generateApplicationSeeder() *InitFactory {
	is.seeders = append(is.seeders, Seed{
		Name: "Create initial application",
		Run: func(db *gorm.DB) error {
			_, errDB := createApplication(db, application)
			return errDB
		},
	})

	return is
}

func (is *InitFactory) generateApplicationApiKeySeeder() *InitFactory {
	is.seeders = append(is.seeders, Seed{
		Name: "Create initial api key for initial application",
		Run: func(db *gorm.DB) error {
			_, errDB := createApplicationApiKey(db, application, applicationApiKey)
			return errDB
		},
	})

	return is
}

func (is *InitFactory) generateApplicationOauthSeeder() *InitFactory {
	is.seeders = append(is.seeders, Seed{
		Name: "Create initial oauth for initial application",
		Run: func(db *gorm.DB) error {
			_, errDB := createApplicationOauth(db, application, applicationOauth)
			return errDB
		},
	})

	return is
}

func (is *InitFactory) generateApplicationOauthClientSeeder() *InitFactory {
	is.seeders = append(is.seeders, Seed{
		Name: "Create initial oauth clients for initial application",
		Run: func(db *gorm.DB) error {
			_, errDB := createApplicationOauthClient(db, application, applicationOauthClient)
			return errDB
		},
	})

	return is
}

func (is *InitFactory) generateSities() *InitFactory {
	for _, st := range sities {
		sity := st
		is.seeders = append(is.seeders, Seed{
			Name: "Create initial sities",
			Run: func(db *gorm.DB) error {
				_, errDB := createSity(db, sity)
				return errDB
			},
		})
	}
	return is
}

func (is *InitFactory) generateVehicle() *InitFactory {
	for _, st := range vehicles {
		vehicle := st
		is.seeders = append(is.seeders, Seed{
			Name: "Create initial vehicles",
			Run: func(db *gorm.DB) error {
				_, errDB := createVehicle(db, vehicle)
				return errDB
			},
		})
	}
	return is
}

func (is *InitFactory) generatePassengerType() *InitFactory {
	for _, st := range passengerTypes {
		passengerType := st
		is.seeders = append(is.seeders, Seed{
			Name: "Create initial passenger_types",
			Run: func(db *gorm.DB) error {
				_, errDB := createPassengerType(db, passengerType)
				return errDB
			},
		})
	}
	return is
}

func (is *InitFactory) generatePrice() *InitFactory {
	for _, st := range prices {
		price := st
		is.seeders = append(is.seeders, Seed{
			Name: "Create initial prices",
			Run: func(db *gorm.DB) error {
				_, errDB := createPrice(db, price)
				return errDB
			},
		})
	}
	return is
}

func (is *InitFactory) generatePassenger() *InitFactory {
	for _, st := range passengers {
		passenger := st
		is.seeders = append(is.seeders, Seed{
			Name: "Create initial passengers",
			Run: func(db *gorm.DB) error {
				_, errDB := createPassenger(db, passenger)
				return errDB
			},
		})
	}
	return is
}

func (is *InitFactory) generateDocumentType() *InitFactory {
	for _, st := range documentTypes {
		documentType := st
		is.seeders = append(is.seeders, Seed{
			Name: "Create initial document_types",
			Run: func(db *gorm.DB) error {
				_, errDB := createDocumentType(db, documentType)
				return errDB
			},
		})
	}
	return is
}

func (is *InitFactory) generateRegularityType() *InitFactory {
	for _, st := range regularityTypes {
		regularityType := st
		is.seeders = append(is.seeders, Seed{
			Name: "Create initial regularity_types",
			Run: func(db *gorm.DB) error {
				_, errDB := createRegularityType(db, regularityType)
				return errDB
			},
		})
	}
	return is
}

func (is *InitFactory) generateOrderStatusType() *InitFactory {
	for _, st := range orderStatusTypes {
		orderStatusType := st
		is.seeders = append(is.seeders, Seed{
			Name: "Create initial order_status_types",
			Run: func(db *gorm.DB) error {
				_, errDB := createOrderStatusType(db, orderStatusType)
				return errDB
			},
		})
	}
	return is
}

func (is *InitFactory) generateDriver() *InitFactory {
	for _, st := range drivers {
		driver := st
		is.seeders = append(is.seeders, Seed{
			Name: "Create initial driver",
			Run: func(db *gorm.DB) error {
				_, errDB := createDriver(db, driver)
				return errDB
			},
		})
	}
	return is
}

func (is *InitFactory) generateRoute() *InitFactory {
	for _, st := range routes {
		route := st
		is.seeders = append(is.seeders, Seed{
			Name: "Create initial route",
			Run: func(db *gorm.DB) error {
				_, errDB := createRoute(db, route)
				return errDB
			},
		})
	}
	return is
}

func (is *InitFactory) generateTrip() *InitFactory {
	for _, st := range trips {
		trip := st
		is.seeders = append(is.seeders, Seed{
			Name: "Create initial trip",
			Run: func(db *gorm.DB) error {
				_, errDB := createTrip(db, trip)
				return errDB
			},
		})
	}
	return is
}

func (is *InitFactory) generateOrder() *InitFactory {
	for _, st := range orders {
		order := st
		is.seeders = append(is.seeders, Seed{
			Name: "Create initial order",
			Run: func(db *gorm.DB) error {
				_, errDB := createOrder(db, order)
				return errDB
			},
		})
	}
	return is
}

func (is *InitFactory) generatePayment() *InitFactory {
	for _, st := range payments {
		payment := st
		is.seeders = append(is.seeders, Seed{
			Name: "Create initial payment",
			Run: func(db *gorm.DB) error {
				_, errDB := createPayment(db, payment)
				return errDB
			},
		})
	}
	return is
}

func initFactory() []Seed {
	initialSeeds := newInitFactory()
	initialSeeds.generateUserSeeder()
	initialSeeds.generateRoleSeeder()
	initialSeeds.generateOtherRoleSeeder()
	initialSeeds.generatePermissionsSeeder()
	initialSeeds.generateRolePermissionsSeeder()
	initialSeeds.generateUserRoleSeeder()
	initialSeeds.generateStorageCategorySeeder()
	initialSeeds.generateApplicationSeeder()
	initialSeeds.generateApplicationApiKeySeeder()
	initialSeeds.generateApplicationOauthSeeder()
	initialSeeds.generateApplicationOauthClientSeeder()
	initialSeeds.generateSities()
	initialSeeds.generateVehicle()
	initialSeeds.generatePassengerType()
	initialSeeds.generatePrice()
	initialSeeds.generatePassenger()
	initialSeeds.generateDocumentType()
	initialSeeds.generateRegularityType()
	initialSeeds.generateOrderStatusType()
	initialSeeds.generateDriver()
	initialSeeds.generateRoute()
	initialSeeds.generateTrip()
	initialSeeds.generateOrder()
	initialSeeds.generatePayment()
	return initialSeeds.seeders
}
