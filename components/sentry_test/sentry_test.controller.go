package sentrytest

import "github.com/gin-gonic/gin"

// @tags Sentry Test API
// @Summary		Crashes the server and sends a report to Sentry
// @Description	Crash the server
// @Accept			json
// @Produce		json
// @Success		200	{string}	string	"ok"
// @Router			/sentry/crash [get]
func testCrash(context *gin.Context) {
	context.String(200, "crash successful")
	panic("y tho")
}
