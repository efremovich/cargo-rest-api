package main

import (
	"cargo-rest-api/config"
	_ "cargo-rest-api/docs"
	"cargo-rest-api/infrastructure/persistence"
	"cargo-rest-api/interfaces/cmd"
	"cargo-rest-api/interfaces/routers"
	"log"
	"os"
	"time"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/rollbar/rollbar-go"

	"github.com/urfave/cli/v2"

	"github.com/joho/godotenv"
)

// @title cargo-rest-api
// @version 1.0
// @description This is a sample of RESTful api made by cargo-rest-api
// @contact.name Aleksand Efremov
// @contact.url https://github.com/efremovich/cargo-rest-api
// @contact.email sasha.fima@gmail.com
// @license.name MIT

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey JWTAuth
// @in header
// @name Authorization
// @description Enter your bearer token in the format **Bearer &lt;token>**

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name api_key

// @securityDefinitions.oauth2.password Oauth2Password
// @tokenUrl http://api.cargo.re-base.ru:8888/oauth/token
// @scope.all Grants all access

// @securityDefinitions.oauth2.accessCode Oauth2AccessCode
// @tokenUrl http://api.cargo.re-base.ru:8888/oauth/token
// @redirectUrl http://api.cargo.re-base.ru:8888/oauth2/callback
// @authorizationUrl http://api.cargo.re-base.ru:8888/oauth/authorize
// @scope.all Grants all access

// @host localhost:8888
// @schemes http
// main init the cargo-rest-api.
func main() {
	// Check .env file
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file provided")
	}

	// Init Config
	conf := config.New()
	timeLoc, _ := time.LoadLocation(conf.AppTimezone)
	time.Local = timeLoc

	// Connect to DB
	dbService, errDBService := persistence.NewDBService(conf.DBConfig)
	if errDBService != nil {
		panic(errDBService)
	}

	// Init DB Migrate
	errAutoMigrate := dbService.AutoMigrate()
	if errAutoMigrate != nil {
		panic(errAutoMigrate)
	}

	// Connect to redis
	redisService, errRedis := persistence.NewRedisService(conf.RedisConfig)
	if errRedis != nil {
		panic(errRedis)
	}

	// Connect to storage services
	storageService, _ := persistence.NewStorageService(conf.MinioConfig, dbService.DB)

	// Init notification services
	notificationService, _ := persistence.NewNotificationService(conf)

	// Init rollbar services
	rollbar.SetToken(conf.RollbarConfig.Token)
	rollbar.SetEnvironment(conf.RollbarConfig.Environment)

	// Init App
	app := cmd.NewCli()
	app.Action = func(c *cli.Context) error {
		// Init Router
		router := routers.NewRouter(conf, dbService, redisService, storageService, notificationService).Init()

		// Inject swagger handler on dev environment
		if conf.AppEnvironment != "production" {
			router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}

		// Run app at defined port
		appPort := os.Getenv("APP_PORT")
		if appPort == "" {
			appPort = "8888"
		}
		log.Println(router.Run(":" + appPort))
		return nil
	}

	// Init Cli
	cliCommands := cmd.NewCommand(dbService)
	app.Commands = cliCommands
	err := app.Run(os.Args)
	if err != nil {
		panic(app)
	}
}
