package secret

import (
	auth "coree/_auth"

	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine) {
	group := router.Group("protected")
	group.Use(auth.FirebaseAuthMiddleware())
	group.GET("/secret", secret)
	group.GET("/profile", profile)
}
