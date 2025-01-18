package dbexample

import (
	"autojidelna/ent"

	"github.com/gin-gonic/gin"
)

// Api for testing that the database is setup correctly
func Register(router *gin.Engine, db *ent.Client) {
	dbClient = db
	app := router.Group("db")
	app.GET("/user/:name", getUser)
	app.POST("/user", setUser)
}
