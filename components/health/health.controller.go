package health

import "github.com/gin-gonic/gin"

// @tags Health Check
// @Summary		Health check for the server
// @Description	Crash the server
// @Accept			json
// @Success		200	{string}	string	"ok"
// @Router			/health [get]
func health(context *gin.Context) {
	context.String(200, "okk")
}
