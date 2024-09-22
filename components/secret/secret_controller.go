package secret

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @tags Protected
// @Summary Get A Secret if Authenticated
// @Security ApiKeyAuth
// @Description	Return secret
// @Accept			text/plain
// @Success		200	{string}	string	"secret"
// @Router			/protected/secret [get]
func secret(context *gin.Context) {
	context.String(http.StatusOK, "Big secret")
}

// @tags Protected
// @Summary Get Profile
// @Description Get the user profile information after authenticating with Firebase.
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]string
// @Router /protected/profile [get]
func profile(context *gin.Context) {
	// Access the verified token from context
	firebaseUser, _ := context.Get("firebaseUser")

	context.JSON(http.StatusOK, gin.H{
		"message":      "You are authenticated",
		"firebaseUser": firebaseUser,
	})
}
