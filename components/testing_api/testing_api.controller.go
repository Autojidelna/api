package testingapi

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// @tags Testing API
// @Summary		Root URL redirections
// @Description	Redirects from the root URL to login or day page.
// @Success		300
// @Router			/testing [get]
func testingRoot(context *gin.Context) {
	context.Redirect(http.StatusFound, "/testing/login")
}

// @tags Testing API
// @Summary		Security
// @Description	Security
// @Param		j_username	formData	string	true	"Username for login"
// @Param		j_password	formData	string	true	"Password for login"
// @Param		_csrf		formData	string	true	"CSRF token"
// @Accept			application/x-www-form-urlencoded
// @Success		200
// @Router			/testing/j_spring_security_check [post]
func testingSecurity(context *gin.Context) {
	username := context.PostForm("j_username")
	password := context.PostForm("j_password")
	xsrfToken := context.PostForm("_csrf")
	targetUrl := context.PostForm("targetUrl")

	//! Very half-baked no worries all WIP
	if !(username == "user" && password == "password") {
		context.SetCookie("remember-me", "", -1, "/", "", false, false)
		context.Redirect(http.StatusFound, "/testing/login")
		return
	}
	context.SetCookie("XSRF-TOKEN", uuid.NewString(), 0, "/", "", true, true)
	context.SetCookie("JSESSIONID", "RANDOM123", 0, "/", "", true, true)

	fmt.Println("username:" + username + " password:" + password + " xsrfToken:" + xsrfToken)
	context.Redirect(http.StatusFound, targetUrl)
}

// @tags Testing API
// @Summary		Day
// @Description	Day
// @Produce		html
// @Success		200
// @Router			/testing/faces/secured/main.jsp [get]
func testingDay(context *gin.Context) {
	dayString := context.Query("day")
	date, err := time.Parse("2006-01-02", dayString)
	if err != nil {
		date = time.Now()
	}

	context.HTML(http.StatusOK, "main.jsp.html", gin.H{
		"Lunches": buildLunches(date),
		"Footer":  buildFooter(),
	})
}

// @tags Testing API
// @Summary		Month
// @Description	Month
// @Produce		html
// @Success		200
// @Router			/testing/faces/secured/month.jsp [get]
func testingMonth(context *gin.Context) {
	context.HTML(http.StatusOK, "month.jsp.html", gin.H{
		"Footer": buildFooter(),
	})
}

// @tags Testing API
// @Summary		Burza
// @Description	Burza
// @Produce		html
// @Success		200
// @Router			/testing/faces/secured/burza.jsp [get]
func testingBurza(context *gin.Context) {
	context.HTML(http.StatusOK, "burza.jsp.html", gin.H{
		"Footer": buildFooter(),
	})
}

// @tags Testing API
// @Summary		Login
// @Description	Login
// @Produce		html
// @Success		200
// @Router			/testing/faces/login.jsp [get]
func testingSetting(context *gin.Context) {
	printer := message.NewPrinter(language.Czech)
	creditString := printer.Sprintf("%.2f", baseCredit)
	context.HTML(http.StatusOK, "setting.2.html", gin.H{
		"Footer": buildFooter(),
		"Credit": creditString,
	})
}

func testingLogin(context *gin.Context) {
	_, errXsfr := context.Request.Cookie("XSRF-TOKEN")
	_, errJsid := context.Request.Cookie("JSESSIONID")
	if errXsfr == http.ErrNoCookie || errJsid == http.ErrNoCookie {
		context.SetCookie("XSRF-TOKEN", uuid.NewString(), 0, "/", "", true, true)
		context.SetCookie("JSESSIONID", "RANDOM123", 0, "/", "", true, true)
	}

	context.HTML(http.StatusOK, "login.html", gin.H{
		"Footer": buildFooter(),
	})
}
