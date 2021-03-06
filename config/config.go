package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// DBConfig represent db config keys.
type DBConfig struct {
	DBDriver   string
	DBHost     string
	DBPort     string
	DBUser     string
	DBName     string
	DBPassword string
	DBTimeZone string
	DBLog      bool
}

// DBTestConfig represent db test config keys.
type DBTestConfig struct {
	DBDriver   string
	DBHost     string
	DBPort     string
	DBUser     string
	DBName     string
	DBPassword string
	DBTimeZone string
	DBLog      bool
}

// RedisConfig represent redis config keys.
type RedisConfig struct {
	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisDB       int
}

// RedisTestConfig represent redis config keys.
type RedisTestConfig struct {
	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisDB       int
}

// MinioConfig represent minio config keys.
type MinioConfig struct {
	Endpoint  string
	AccessKey string
	SecretKey string
	Bucket    string
}

// SMTPConfig represent SMTP config keys.
type SMTPConfig struct {
	SMTPHost     string
	SMTPPort     int
	SMTPUsername string
	SMTPPassword string
}

// RollbarConfig represent rollbar config keys.
type RollbarConfig struct {
	Token       string
	Environment string
	CodeVersion string
	ServerHost  string
	ServerRoot  string
}

// Oauth2Config represent key config keys.
type Oauth2Config struct {
	OauthID     string
	OauthSecret string
	OauthDomain string
	OauthToken  string
}

// KeyConfig represent key config keys.
type KeyConfig struct {
	AppPrivateKey string
	AppPublicKey  string
}

// Config represent config keys.
type Config struct {
	DBConfig
	DBTestConfig
	RedisConfig
	RedisTestConfig
	MinioConfig
	SMTPConfig
	RollbarConfig
	Oauth2Config
	KeyConfig
	AppEnvironment  string
	AppLanguage     string
	AppTimezone     string
	EnableCors      bool
	EnableLogger    bool
	EnableRequestID bool
	DebugMode       bool
}

// New returns a new Config struct.
func New() *Config {
	return &Config{
		DBConfig: DBConfig{
			DBDriver:   getEnv("DB_DRIVER", "mysql"),
			DBHost:     getEnv("DB_HOST", "localhost"),
			DBPort:     getEnv("DB_PORT", "3306"),
			DBUser:     getEnv("DB_USER", "root"),
			DBName:     getEnv("DB_NAME", "go_rest_skeleton"),
			DBPassword: getEnv("DB_PASSWORD", ""),
			DBTimeZone: getEnv("APP_TIMEZONE", "Europe/Moscow"),
			DBLog:      getEnvAsBool("ENABLE_LOGGER", true),
		},
		DBTestConfig: DBTestConfig{
			DBDriver:   getEnv("TEST_DB_DRIVER", "mysql"),
			DBHost:     getEnv("TEST_DB_HOST", "localhost"),
			DBPort:     getEnv("TEST_DB_PORT", "3306"),
			DBUser:     getEnv("TEST_DB_USER", "root"),
			DBName:     getEnv("TEST_DB_NAME", "go_rest_skeleton_test"),
			DBPassword: getEnv("TEST_DB_PASSWORD", ""),
			DBTimeZone: getEnv("APP_TIMEZONE", "Europe/Moscow"),
			DBLog:      getEnvAsBool("ENABLE_LOGGER", true),
		},
		RedisConfig: RedisConfig{
			RedisHost:     getEnv("REDIS_HOST", "127.0.0.1"),
			RedisPort:     getEnv("REDIS_PORT", "6379"),
			RedisPassword: getEnv("REDIS_PASSWORD", ""),
			RedisDB:       getEnvAsInt("REDIS_DB", 0),
		},
		RedisTestConfig: RedisTestConfig{
			RedisHost:     getEnv("TEST_REDIS_HOST", "127.0.0.1"),
			RedisPort:     getEnv("TEST_REDIS_PORT", "6379"),
			RedisPassword: getEnv("TEST_REDIS_PASSWORD", ""),
			RedisDB:       getEnvAsInt("TEST_REDIS_DB", 10),
		},
		MinioConfig: MinioConfig{
			Endpoint:  getEnv("MINIO_HOST", "127.0.0.1:9000"),
			AccessKey: getEnv("MINIO_ROOT_USER", "minio"),
			SecretKey: getEnv("MINIO_ROOT_PASSWORD", "minioPass"),
			Bucket:    getEnv("MINIO_BUCKET", "cargo-rest-api"),
		},
		SMTPConfig: SMTPConfig{
			SMTPHost:     getEnv("SMTP_HOST", ""),
			SMTPPort:     getEnvAsInt("SMTP_PORT", 587),
			SMTPUsername: getEnv("SMTP_USERNAME", ""),
			SMTPPassword: getEnv("SMTP_PASSWORD", ""),
		},
		RollbarConfig: RollbarConfig{
			Token:       getEnv("ROLLBAR_TOKEN", ""),
			Environment: getEnv("APP_ENV", "local"),
			CodeVersion: getEnv("APP_VERSION", "1.0.0"),
			ServerHost:  getEnv("APP_SERVER_HOST", ""),
			ServerRoot:  getEnv("APP_SERVER_ROOT", ""),
		},
		Oauth2Config: Oauth2Config{
			OauthID:     getEnv("OAUTH_ID", "cargo-rest-api"),
			OauthSecret: getEnv("OAUTH_SECRET", "cargo-rest-api"),
			OauthDomain: getEnv("OAUTH_DOMAIN",
				fmt.Sprintf("http://localhost:%s", os.Getenv("APP_PORT"))),
			OauthToken: getEnv("OAUTH_TOKEN", "Bearer"),
		},
		KeyConfig: KeyConfig{
			AppPrivateKey: getEnv("APP_PRIVATE_KEY", "default-private-key"),
			AppPublicKey:  getEnv("APP_PUBLIC_KEY", "default-public-key"),
		},
		AppEnvironment:  getEnv("APP_ENV", "local"),
		AppLanguage:     getEnv("APP_LANG", "en"),
		AppTimezone:     getEnv("APP_TIMEZONE", "Europe/Moscow"),
		EnableCors:      getEnvAsBool("ENABLE_CORS", true),
		EnableLogger:    getEnvAsBool("ENABLE_LOGGER", true),
		EnableRequestID: getEnvAsBool("ENABLE_REQUEST_ID", true),
		DebugMode:       getEnv("APP_ENV", "local") != "production",
	}
}

// Simple helper function to read an environment or return a default value.
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	if nextValue := os.Getenv(key); nextValue != "" {
		return nextValue
	}

	return defaultVal
}

// Simple helper function to read an environment variable into integer or return a default value.
func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

// Helper to read an environment variable into a bool or return default value.
func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}

// Helper to read an environment variable into a string slice or return default value.
func getEnvAsSlice(name string, defaultVal []string, sep string) []string {
	valStr := getEnv(name, "")

	if valStr == "" {
		return defaultVal
	}

	val := strings.Split(valStr, sep)

	return val
}
