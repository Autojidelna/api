package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @tags Health Check
// @Summary		Health check for the server
// @Description	Check for the health of the server
// @Accept			json
// @Success		200	{string}	string	"ok"
// @Router			/health [get]
func health(context *gin.Context) {
	context.String(http.StatusOK, "ok")
}
