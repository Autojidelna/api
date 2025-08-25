package testingapi

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Api creating dummy icanteen instance for testing/review
func Register(router *gin.Engine) {
	ok := initUsers()
	if !ok {
		fmt.Printf("Aborted initialization of Testing API!!! Failed initUsers")
		return
	}
	ok = initLunches()
	if !ok {
		fmt.Printf("Aborted initialization of Testing API!!! Failed initLunches")
		return
	}

	initSessions()
	initOrdersState()

	router.LoadHTMLGlob("./assets/templates/**/*")
	app := router.Group("testing")
	app.GET("/", testingRoot)
	app.GET("/faces/secured/db/dbProcessOrder.jsp", testingOrder)
	app.POST("/j_spring_security_check", testingSecurity)
	app.GET("/faces/secured/main.jsp", testingDay)
	app.GET("/faces/secured/month.jsp", testingMonth)
	app.GET("/faces/secured/burza.jsp", testingBurza)
	app.GET("/login", testingLogin)
	app.GET("/web/setting", testingSetting)

}
