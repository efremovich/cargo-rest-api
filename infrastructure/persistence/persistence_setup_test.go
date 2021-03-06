package persistence_test

import (
	"cargo-rest-api/config"
	"cargo-rest-api/domain/entity"
	"cargo-rest-api/domain/registry"
	"cargo-rest-api/infrastructure/authorization"
	"cargo-rest-api/infrastructure/persistence"
	"cargo-rest-api/pkg/util"
	"fmt"
	"log"
	"testing"

	"github.com/bxcodec/faker"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"

	"github.com/google/uuid"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

const (
	driverMysql    = "mysql"
	driverPostgres = "postgres"
)

// SkipThis is a function.
func SkipThis(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test")
	}
}

func InitConfig() *config.Config {
	if err := godotenv.Load(fmt.Sprintf("%s/.env", util.RootDir())); err != nil {
		log.Println("no .env file provided")
	}

	return config.New()
}

// DBConn is a function.
func DBConn() (*gorm.DB, error) {
	conf := InitConfig()
	return DBConnSetup(conf.DBTestConfig)
}

// DBService is a function.
func DBService() (*persistence.Repositories, error) {
	conf := InitConfig()
	return DBServiceSetup(conf.DBTestConfig)
}

// RedisConn is a function.
func RedisConn() (*gorm.DB, error) {
	conf := InitConfig()
	return DBConnSetup(conf.DBTestConfig)
}

// RedisService is a function.
func RedisService() (*persistence.RedisService, error) {
	conf := InitConfig()
	return RedisServiceSetup(conf.RedisTestConfig)
}

// DBConnSetup is a function.
func DBConnSetup(config config.DBTestConfig) (*gorm.DB, error) {
	dbURL := ""
	switch config.DBDriver {
	case driverPostgres:
		dbURL = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
			config.DBHost,
			config.DBPort,
			config.DBUser,
			config.DBName,
			config.DBPassword,
		)
	case driverMysql:
		dbURL = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			config.DBUser,
			config.DBPassword,
			config.DBHost,
			config.DBPort,
			config.DBName,
		)
	}

	gormConfig := &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	}

	db, err := gorm.Open(mysql.Open(dbURL), gormConfig)

	if err != nil {
		return nil, err
	}

	tables := registry.CollectTableNames()
	for _, table := range tables {
		err := db.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", table.Name)).Error
		if err != nil {
			log.Fatal(err)
		}
	}

	entities := registry.CollectEntities()
	for _, model := range entities {
		err := db.AutoMigrate(model.Entity)
		if err != nil {
			log.Fatal(err)
		}
	}

	return db, nil
}

// DBServiceSetup will initialize db connection and return repositories.
func DBServiceSetup(config config.DBTestConfig) (*persistence.Repositories, error) {
	db, err := DBConnSetup(config)
	if err != nil {
		return nil, err
	}

	return &persistence.Repositories{
		Permission:      persistence.NewPermissionRepository(db),
		Role:            persistence.NewRoleRepository(db),
		StorageCategory: persistence.NewStorageCategoryRepository(db),
		StorageFile:     persistence.NewStorageFileRepository(db),
		User:            persistence.NewUserRepository(db),
		UserPreference:  persistence.NewUserPreferenceRepository(db),
		DB:              db,
	}, nil
}

// RedisConnSetup will initialize connection to redis server.
func RedisConnSetup(config config.RedisTestConfig) (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.RedisHost + ":" + config.RedisPort,
		Password: config.RedisPassword,
		DB:       config.RedisDB,
	})

	return redisClient, nil
}

// RedisServiceSetup will initialize connection to redis server.
func RedisServiceSetup(config config.RedisTestConfig) (*persistence.RedisService, error) {
	redisClient, err := RedisConnSetup(config)
	if err != nil {
		return nil, err
	}

	return &persistence.RedisService{
		Auth:   authorization.NewAuth(redisClient),
		Client: redisClient,
	}, nil
}

func seedUser(db *gorm.DB) (*entity.User, *entity.UserFaker, error) {
	userFaker := entity.UserFaker{}
	_ = faker.FakeData(&userFaker)
	user := entity.User{
		Name:     userFaker.Name,
		Email:    userFaker.Email,
		Phone:    userFaker.Phone,
		Password: userFaker.Password,
	}
	err := db.Create(&user).Error
	if err != nil {
		return nil, nil, err
	}

	return &user, &userFaker, nil
}

func seedRole(db *gorm.DB) (*entity.Role, error) {
	role := entity.Role{
		UUID: uuid.New().String(),
		Name: "Example Role",
	}
	err := db.Create(&role).Error
	if err != nil {
		return nil, err
	}

	return &role, nil
}

func seedRoles(db *gorm.DB) ([]entity.Role, error) {
	roles := []entity.Role{
		{UUID: uuid.New().String(), Name: "Super Administrator"},
		{UUID: uuid.New().String(), Name: "Administrator"},
		{UUID: uuid.New().String(), Name: "User"},
	}

	for _, v := range roles {
		role := v
		err := db.Create(&role).Error
		if err != nil {
			return nil, err
		}
	}

	return roles, nil
}

func seedUserPreference(db *gorm.DB) (*entity.UserPreference, error) {
	var userPreference entity.UserPreference
	userPreference.UUID = uuid.New().String()
	userPreference.UserUUID = uuid.New().String()
	userPreference.Preference = userPreference.BuildDefaultPreference()

	err := db.Create(&userPreference).Error
	if err != nil {
		return nil, err
	}

	return &userPreference, nil
}
