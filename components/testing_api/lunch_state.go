package testingapi

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

type LunchMeal struct {
	Jidla []Meal `json:"jidla"`
}

type Meal struct {
	Nazev       string     `json:"nazev"`
	Varianta    string     `json:"varianta"`
	Objednano   bool       `json:"objednano"`
	LzeObjednat bool       `json:"lzeObjednat"`
	NaBurze     bool       `json:"naBurze"`
	Cena        float64    `json:"cena"`
	Alergeny    []Allergen `json:"alergeny"`
}

type Allergen struct {
	Nazev string `json:"nazev"`
	Popis string `json:"popis"`
}

var lunchStore map[string]LunchMeal
var lastUpdatedDate time.Time

func initLunches() bool {
	lunchStore = make(map[string]LunchMeal)
	lastUpdatedDate = time.Now()

	// Generate lunches for specific days
	initialDay := time.Now().AddDate(0, 0, -LUNCH_GEN_PAST_DAYS)
	for i := 0; i <= LUNCH_GEN_FUTURE_DAYS+LUNCH_GEN_PAST_DAYS; i++ {
		currentDay := initialDay.AddDate(0, 0, i)

		updateLunchDay(currentDay)
	}
	return true
}

func getRawLunches() (map[string]LunchMeal, bool) {
	var rawLunches map[string]LunchMeal
	fileBytes, err := os.ReadFile("assets/json/meals.json")
	if err != nil {
		fmt.Println("Error when reading meals.json file")
		return rawLunches, false
	}

	err = json.Unmarshal(fileBytes, &rawLunches)
	if err != nil {
		fmt.Println("Error when parsing meals.json file")
		return rawLunches, false
	}

	return rawLunches, true
}

/* Generates lunch for a specific day*/
func updateLunchDay(date time.Time) {
	dateString := date.Format(DATE_FORMAT_YYYY_DD_MM)
	_, ok := lunchStore[dateString]
	if ok {
		return
	}

	rawLunches, ok := getRawLunches()
	if !ok {
		return
	}

	_, week := date.ISOWeek()
	lunchId := 0
	if week%2 == 0 {
		lunchId = 5
	}
	lunchId += int(date.Weekday())
	lunchIdString := strconv.Itoa(lunchId)

	fmt.Println("[LUNCH] Generating for date:" + dateString + " lunchId:" + lunchIdString)

	lunchStore[dateString] = rawLunches[lunchIdString]
}

/* Updates the lunch database when the date changes.*/
func UpdateLunchesAndOrders() {
	currentDate := time.Now()

	if lastUpdatedDate.Day() == currentDate.Day() {
		fmt.Println("Update of lunches and orders skipped - waiting for the next day.")
		return
	}
	fmt.Println("Update of lunches and orders started.")

	// Add new lunches
	daysSince := int(math.Floor((currentDate.Sub(lastUpdatedDate).Hours() + 24) / 24))
	for i := 0; i <= daysSince; i++ {
		addDate := currentDate.AddDate(0, 0, LUNCH_GEN_FUTURE_DAYS-i)
		updateLunchDay(addDate)
		updateOrderDay(addDate)
	}

	lastUpdatedDate = currentDate
}

func getLunchesDay(date time.Time) (LunchMeal, bool) {
	dateString := date.Format(DATE_FORMAT_YYYY_DD_MM)
	lunch, ok := lunchStore[dateString]
	return lunch, ok
}
