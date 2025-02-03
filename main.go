package main

//	@title			APPE autojidelna example server
//	@version		1.0
//	@description	This is an example for all api server projects in App Elevate

// 	@tag.name 	DB Example API
// 	@tag.description	API for testing that the database is setup correctly
// 	@tag.name	Sentry Test API
// 	@tag.description	Api for testing Sentry is setup correctly
// 	@tag.name	Health Check
// 	@tag.description	Health Check for this API
// 	@tag.name	Protected
// 	@tag.description	Endpoints that require authentication
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

import (
	dbexample "autojidelna/components/db_example"
	"autojidelna/components/health"
	sentrytest "autojidelna/components/sentry_test"
	testingapi "autojidelna/components/testing_api"
	"autojidelna/ent"
	"context"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"

	"fmt"

	_ "github.com/lib/pq"

	_ "autojidelna/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var dbClient *ent.Client

func setupRouter() *gin.Engine {
	app := gin.Default()
	app.Use(sentrygin.New(sentrygin.Options{
		Repanic: true,
	}))

	// register modules
	testingapi.Register(app)
	dbexample.Register(app, dbClient)
	sentrytest.Register(app)
	health.Register(app)

	return app

	app.GET("/", func(context *gin.Context) {
		context.Redirect(http.StatusTemporaryRedirect, "/docs/index.html")
	})
	app.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return app
}

func getVariable(key string, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		value = defaultValue
	}
	return value
}

func getFileVariable(key string, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if ok {
		valueBytes, err := os.ReadFile(value)
		ok = err == nil
		if ok {
			value = string(valueBytes)
		}
	}
	if !ok {
		value = defaultValue
	}
	return value
}

func getIntVariable(key string, defaultValue int) int {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return intValue
}

func getCredentials() (string, int, string, string, string) {
	dbHost := getVariable("POSTGRES_HOST", "localhost")
	dbUser := getFileVariable("POSTGRES_USER_FILE", "postgres")
	dbPassword := getFileVariable("POSTGRES_PASSWORD_FILE", "postgres")
	dbName := getFileVariable("POSTGRES_DB_FILE", "postgres")
	dbPort := getIntVariable("POSTGRES_PORT", 5432)
	return dbHost, dbPort, dbUser, dbPassword, dbName
}

func initDatabase() {
	dbHost, dbPort, dbUser, dbPassword, dbName := getCredentials()
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)
	for {
		db, err := ent.Open("postgres", psqlInfo)
		if err != nil {
			log.Printf("failed opening connection to postgres: %v", err)
		}

		// Run the auto migration tool.
		if err := db.Schema.Create(context.Background()); err != nil {
			log.Printf("failed creating schema resources: %v", err)
		}
		dbClient = db
		time.Sleep(time.Second)
		if dbClient != nil && err == nil {
			break
		}
	}

}

func initSentry() {
	// To initialize Sentry's handler, you need to initialize Sentry itself beforehand
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:           "https://0a46799bef1e6ceb83bc77eba5c5aaea@o4507799131258880.ingest.de.sentry.io/4507928244256848",
		EnableTracing: true,
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for tracing.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}
}

func main() {
	initSentry()
	initDatabase()
	defer dbClient.Close()

	app := setupRouter()
	port := getVariable("APP_PORT", "80")
	port = ":" + port
	println("App is running on http://localhost" + port)
	app.Run(port)

}
