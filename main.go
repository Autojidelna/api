package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"database/sql"

	_ "github.com/lib/pq"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
)

const (
	dbHost     = "localhost"
	dbPort     = 5433
	dbUser     = "tom"
	dbPassword = "tom"
	dbName     = "tom"
)

var db *sql.DB

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	app := gin.Default()
	// Once it's done, you can attach the handler as one of your middleware
	app.Use(sentrygin.New(sentrygin.Options{
		Repanic: true,
	}))

	// Ping test
	app.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})

	// Get user value
	app.GET("/user/:name", func(context *gin.Context) {
		user := context.Params.ByName("name")
		context.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := app.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	authorized.POST("admin", func(context *gin.Context) {
		user := context.MustGet(gin.AuthUserKey).(string)
		fmt.Print("user: "+user+"\n", gin.Logger())

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if context.Bind(&json) == nil {
			context.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return app
}

func main() {
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
	// To initialize Sentry's handler, you need to initialize Sentry itself beforehand
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)
	dbHere, err := sql.Open("postgres", psqlInfo)
	db = dbHere
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	if gin.Mode() == gin.DebugMode {
		fmt.Println("Gin is running in debug mode")
		// Do something specific for debug mode
	} else {
		fmt.Println("Gin is running in release or test mode")
		// Do something else for release or test mode
	}
	println("Starting server")

	app := setupRouter()
	app.Run(":8080")

}
