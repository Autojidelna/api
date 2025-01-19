package testingapi

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @tags Testing API
// @Summary		Root URL redirections
// @Description	Redirects from the root URL to login or day page.
// @Success		300
// @Router			/testing [get]
func testingRoot(context *gin.Context) {
	context.Redirect(http.StatusFound, "/testing/faces/secured/main.jsp")
}

// @tags Testing API
// @Summary		Security
// @Description	Security
// @Accept			json
// @Success		200
// @Router			/testing/j_spring_security_check [post]
func testingSecurity(context *gin.Context) {
	context.String(http.StatusOK, "cool ")
	fmt.Println("security")
}

// @tags Testing API
// @Summary		Day
// @Description	Day
// @Produce		html
// @Success		200
// @Router			/testing/faces/secured/main.jsp [get]
func testingDay(context *gin.Context) {
	context.HTML(http.StatusOK, "main.jsp.html", gin.H{})
}
