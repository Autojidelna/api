package testingapi

import "github.com/gin-gonic/gin"

// Api creating dummy icanteen instance for testing/review
func Register(router *gin.Engine) {
	router.LoadHTMLGlob("./components/testing_api/templates/*")
	app := router.Group("testing")
	app.GET("/", testingRoot)
	app.POST("/j_spring_security_check", testingSecurity)
	app.GET("/faces/secured/main.jsp", testingDay)
}
