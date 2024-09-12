package main

//	@title			APPE COREE example server
//	@version		1.0
//	@description	This is an example for all api server projects in App Elevate

// 	@tag.name 	DB Example API
// 	@tag.description	API for testing that the database is setup correctly
// 	@tag.name	Sentry Test API
// 	@tag.description	Api for testing Sentry is setup correctly
// 	@tag.name	Health Check
// 	@tag.description	Health Check for this API

//	@host		localhost:8080
//	@BasePath	/

import (
	"context"
	dbexample "coree/components/db_example"
	"coree/components/health"
	sentrytest "coree/components/sentry_test"
	"coree/ent"
	"log"
	"net/http"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"

	"fmt"

	_ "github.com/lib/pq"

	_ "coree/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	dbHost     = "localhost"
	dbPort     = 5433
	dbUser     = "tom"
	dbPassword = "tom"
	dbName     = "tom"
)

var dbClient *ent.Client

func setupRouter() *gin.Engine {
	app := gin.Default()
	app.Use(sentrygin.New(sentrygin.Options{
		Repanic: true,
	}))

	// register modules
	dbexample.Register(app, dbClient)
	sentrytest.Register(app)
	health.Register(app)

	if gin.Mode() == gin.DebugMode {
		fmt.Println("Gin is running in debug mode creating swagger docs")
		app.GET("/", func(context *gin.Context) {
			context.Redirect(http.StatusMovedPermanently, "/docs/index.html")
		})
		app.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	} else {
		fmt.Println("Gin is running in release mode")
	}

	return app
}

func initDatabase() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)
	db, err := ent.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	// Run the auto migration tool.
	if err := db.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	dbClient = db
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
	println("App is running on http://localhost:8080")
	app.Run(":8080")

}
