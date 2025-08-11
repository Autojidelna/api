package testingapi

// Behavior constants, these change some basic parameters
const (
	BASE_CREDIT           = 469.69
	BASE_PRICE            = 42
	ORDER_CUTOFF_HOURS    = 10
	LUNCH_GEN_PAST_DAYS   = 7
	LUNCH_GEN_FUTURE_DAYS = 14
)

// Technical constants, !! only change when you know what you're doing !!
const (
	DATE_FORMAT_YYYY_DD_MM = "2006-01-02"
	COOKIE_SESSION_ID      = "JSESSIONID"
	BCRYPT_COST            = 16
	BASE_VERSION           = "2.18.03"
)

// HTML constants
const (
	LUNCH_UNAVAILABLE_STRING = `<div id="orderContent2025-02-22" class="orderContent">
    <div class="textGrey">
        <span class="fa-stack fa-lg">
            <i class="far fa-square fa-stack-2x"></i>
            <i class="fa fa-exclamation fa-stack-1x"></i>
        </span>
    Litujeme, ale na vybrany den nejsou zadana v jidelnicku zadna jidla. Vyberte jiny den, nebo kontaktujte vasi jidelnu a pozadejteji o schvaleni jidelnicku.
	</div>
</div>`
)
