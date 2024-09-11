package main

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

import (
	"context"
	"coree/ent"
	"log"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"

	"coree/ent/user"

	"fmt"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
)

const (
	dbHost     = "localhost"
	dbPort     = 5433
	dbUser     = "tom"
	dbPassword = "tom"
	dbName     = "tom"
)

type UserInput struct {
	Name string `json:"name" binding:"required"`
}

var dbClient *ent.Client

func QueryUser(ctx context.Context, client *ent.Client, name string) (*ent.User, error) {
	u, err := client.User.
		Query().
		Where(user.Name(name)).
		// `Only` fails if no user found,
		// or more than 1 user returned.
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}

func CreateUser(ctx context.Context, client *ent.Client, name string) (*ent.User, error) {
	u, err := client.User.
		Create().
		SetAge(30).
		SetName(name).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}

func setupRouter() *gin.Engine {
	if gin.Mode() == gin.DebugMode {
		fmt.Println("Gin is running in debug mode")
		// Do something specific for debug mode
	} else {
		fmt.Println("Gin is running in release or test mode")
		// Do something else for release or test mode
	}
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
		resultUser, err := QueryUser(context, dbClient, user)
		if err != nil {
			context.JSON(http.StatusOK, gin.H{"user": resultUser, "status": "error"})
			return
		}
		context.JSON(http.StatusOK, gin.H{"user": resultUser, "status": "ok"})
	})

	app.POST("/users", func(c *gin.Context) {
		var input UserInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Save the user to the database
		user, err := dbClient.User.
			Create().
			SetName(input.Name).
			Save(context.Background())

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Respond with the created user
		c.JSON(http.StatusOK, gin.H{
			"id":   user.ID,
			"name": user.Name,
		})
	})
	app.GET("/foo", func(ctx *gin.Context) {
		// sentrygin handler will catch it just fine. Also, because we attached "someRandomTag"
		// in the middleware before, it will be sent through as well
		panic("y tho")
	})

	authorized := app.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

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

	println("Starting server")

	app := setupRouter()
	app.Run(":8080")

}
