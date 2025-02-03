package testingapi

import "github.com/gin-gonic/gin"

// Api creating dummy icanteen instance for testing/review
func Register(router *gin.Engine) {
	router.LoadHTMLGlob("./components/testing_api/templates/*")
	//app := router.Group("testing")
	router.GET("/", testingRoot)
	router.POST("/j_spring_security_check", testingSecurity)
	router.GET("/faces/secured/main.jsp", testingDay)
	router.GET("/faces/secured/month.jsp", testingMonth)
	router.GET("/faces/secured/burza.jsp", testingBurza)
	router.GET("/login", testingLogin)
	router.GET("/web/setting", testingSetting)
}
