package dbexample

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// setting the db tag
// @tags	DB Example API
// @Summary		Create a user
// @Description	create a user using post
// @Accept			json
// @Produce		json
// @Param			user	body		UserCreate	true	"User object"
// @Success		200	{object}	UserGet		"ok"
// @Failure		404	{object}	map[string]interface{}		"User not found"
// @Failure		500	{object}	map[string]interface{}		"Internal server error"
// @Router			/db/user [post]
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
// @tags	DB Example API
// @Accept			json
// @Produce		json
// @Param			name		path		string		true	"User name"
// @Success		200	{object}	UserGet		"ok"
// @Failure		404	{object}	map[string]interface{}		"User not found"
// @Failure		500	{object}	map[string]interface{}		"Internal server error"
// @Router			/db/user/{name} [get]
func getUser(context *gin.Context) {
	user := context.Params.ByName("name")
	resultUser, err := queryUser(context, dbClient, user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, resultUser)
}
