package testingapi

import (
	"fmt"
	"net/http"
	"strconv"
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
// @Param		targetUrl	formData	string	true	"redirection target"
// @Accept			application/x-www-form-urlencoded
// @Success		200
// @Router			/testing/j_spring_security_check [post]
func testingSecurity(context *gin.Context) {
	username := context.PostForm("j_username")
	password := context.PostForm("j_password")
	xsrfToken := context.PostForm("_csrf")
	targetUrl := context.PostForm("targetUrl")

	infoString := "[SECURITY] username:" + username + " password:" + password + " xsrfToken:" + xsrfToken

	sessionId, err := generateSession(username, password)
	if err != nil {
		context.SetCookie("remember-me", "", -1, "/", "", false, false)
		context.Redirect(http.StatusFound, "/testing/login")
		fmt.Println(infoString + " sessionId: AUTH FAILED")
		return
	}
	fmt.Println(infoString + " sessionId:" + sessionId)

	context.SetCookie("XSRF-TOKEN", uuid.NewString(), 0, "/", "", true, true)
	context.SetCookie(COOKIE_SESSION_ID, sessionId, 0, "/", "", true, true)

	context.Redirect(http.StatusFound, targetUrl)
}

// @tags Testing API
// @Summary		Day
// @Description	Day
// @Produce		html
// @Success		200
// @Router			/testing/faces/secured/main.jsp [get]
func testingDay(context *gin.Context) {
	sessionId, err := context.Cookie(COOKIE_SESSION_ID)
	if err != nil {
		context.Status(http.StatusForbidden)
		return
	}
	username, err := getSessionUsername(sessionId)
	if err != nil {
		context.Status(http.StatusForbidden)
		return
	}

	dayString := context.Query("day")
	date, err := time.Parse("2006-01-02", dayString)
	if err != nil {
		date = time.Now()
	}

	context.HTML(http.StatusOK, "main.jsp.html", gin.H{
		"Lunches": buildLunches(date, username),
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
// @Router			/testing/web/setting [get]
func testingSetting(context *gin.Context) {
	// Session ID
	sessionId, err := context.Cookie(COOKIE_SESSION_ID)
	if err != nil {
		context.Status(http.StatusForbidden)
		return
	}
	username, err := getSessionUsername(sessionId)
	if err != nil {
		context.Status(http.StatusForbidden)
		return
	}
	credit, _ := getUserCredit(username)
	printer := message.NewPrinter(language.Czech)
	creditString := printer.Sprintf("%.2f", credit)
	context.HTML(http.StatusOK, "setting.2.html", gin.H{
		"Footer": buildFooter(),
		"Credit": creditString,
	})
}

func testingOrder(context *gin.Context) {
	// sessionId & username
	sessionId, err := context.Cookie(COOKIE_SESSION_ID)
	if err != nil {
		context.Status(http.StatusForbidden)
		return
	}
	username, err := getSessionUsername(sessionId)
	if err != nil {
		context.Status(http.StatusForbidden)
		return
	}
	// Day
	mealDate, err := time.Parse("2006-01-02", context.Query("day"))
	if err != nil {
		context.String(http.StatusBadRequest, "Invalid Query Parameter DAY! Must resolve to DATE in 2006-01-02 format!")
		fmt.Println("Invalid Query Parameter DAY! Must resolve to DATE in 2006-01-02 format!")
	}
	// Meal ID
	mealIndex, err := strconv.Atoi(context.Query("ID"))
	if err != nil {
		context.String(http.StatusBadRequest, "Invalid Query Parameter ID! Must resolve to INT32!")
		fmt.Println("Invalid Query Parameter ID! Must resolve to INT32!")
	}
	// Type
	transactionType := context.Query("type")
	switch transactionType {
	case "delete":
		setUserOrder(username, mealDate, 0)
	case "make":
		setUserOrder(username, mealDate, mealIndex)
	case "reorder":
		setUserOrder(username, mealDate, mealIndex)
	default:
		context.String(400, "Invalid Query Parameter TYPE! Must be 'make', 'reorder' or 'delete'!")
		fmt.Println("Invalid Query Parameter TYPE! Must be 'make', 'reorder' or 'delete'!")
	}
	context.Status(200)
}

func testingLogin(context *gin.Context) {
	_, errXsfr := context.Request.Cookie("XSRF-TOKEN")
	_, errJsid := context.Request.Cookie(COOKIE_SESSION_ID)
	if errXsfr == http.ErrNoCookie {
		context.SetCookie("XSRF-TOKEN", uuid.NewString(), 0, "/", "", true, true)
	}
	if errJsid == http.ErrNoCookie {
		context.SetCookie(COOKIE_SESSION_ID, "NOT-LOGGED-IN", 0, "/", "", true, true)
	}

	context.HTML(http.StatusOK, "login.html", gin.H{
		"Footer": buildFooter(),
	})
}
