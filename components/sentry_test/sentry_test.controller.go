package sentrytest

import "github.com/gin-gonic/gin"

//	@Summary		Crashes the server and sends a report to Sentry
//	@Description	Crash the server
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	string	"ok"
//	@Router			/crash [get]
func testCrash(context *gin.Context) {
	context.String(200, "ok")
	panic("y tho")
}
