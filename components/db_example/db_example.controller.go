package dbexample

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary		Create a user
// @Description	create a user using post
// @Accept			json
// @Produce		json
// @Param			user	body		UserCreate	true	"User object"
// @Success		200	{object}	UserGet		"ok"
// @Router			/db-example/user [post]
func setUser(context *gin.Context) {
	var input UserCreate
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := createUser(context, dbClient, input.Name, input.Age)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Respond with the created user
	context.JSON(http.StatusOK, user)
}

// @Summary		Get a User by name
// @Description	Get a User by name
// @Accept			json
// @Produce		json
// @Param			name		path		string		true	"User name"
// @Success		200	{object}	UserGet		"ok"
// @Router			/db-example/user/{name} [get]
func getUser(context *gin.Context) {
	user := context.Params.ByName("name")
	resultUser, err := queryUser(context, dbClient, user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, resultUser)
}
