package testingapi

import "github.com/gin-gonic/gin"

// Api creating dummy icanteen instance for testing/review
func Register(router *gin.Engine) {
	router.LoadHTMLGlob("./assets/templates/**/*")
	app := router.Group("testing")
	app.GET("/", testingRoot)
	app.POST("/j_spring_security_check", testingSecurity)
	app.GET("/faces/secured/main.jsp", testingDay)
	app.GET("/faces/secured/month.jsp", testingMonth)
	app.GET("/faces/secured/burza.jsp", testingBurza)
	app.GET("/login", testingLogin)
	app.GET("/web/setting", testingSetting)
}
