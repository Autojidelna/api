package dbexample

import (
	"coree/ent"

	"github.com/gin-gonic/gin"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /

func Register(router *gin.Engine, db *ent.Client) {
	dbClient = db
	app := router.Group("db-example")
	app.GET("/user/:name", getUser)
	app.POST("/user", setUser)
}
