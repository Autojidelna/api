package health

import "github.com/gin-gonic/gin"

// Api for testing Sentry is setup correctly
func Register(router *gin.Engine) {
	router.GET("/health", health)
}
