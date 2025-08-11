package testingapi

import (
	"encoding/json"
	"fmt"
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

func initLunches() bool {
	lunchStore = make(map[string]LunchMeal)

	// Read the raw lunches from a file
	var rawLunches map[string]LunchMeal
	fileBytes, err := os.ReadFile("assets/json/meals.json")
	if err != nil {
		fmt.Println("Error when reading meals.json file")
		return false
	}

	err = json.Unmarshal(fileBytes, &rawLunches)
	if err != nil {
		fmt.Println("Error when parsing meals.json file")
		return false
	}
	// Generate lunches for specific days
	initialDay := time.Now().AddDate(0, 0, -LUNCH_GEN_PAST_DAYS)
	for i := 0; i <= LUNCH_GEN_FUTURE_DAYS+LUNCH_GEN_PAST_DAYS; i++ {
		currentDay := initialDay.AddDate(0, 0, i)
		currentDayString := currentDay.Format(DATE_FORMAT_YYYY_DD_MM)
		_, week := currentDay.ISOWeek()
		lunchId := 0
		if week%2 == 0 {
			lunchId = 5
		}
		lunchId += int(currentDay.Weekday())
		lunchIdString := strconv.Itoa(lunchId)

		fmt.Println("[LUNCH] Generating for date:" + currentDayString + " lunchId:" + lunchIdString)

		lunchStore[currentDayString] = rawLunches[lunchIdString]
	}
	return true
}

func getLunchesDay(date time.Time) (LunchMeal, bool) {
	dateString := date.Format(DATE_FORMAT_YYYY_DD_MM)
	lunch, ok := lunchStore[dateString]
	return lunch, ok
}
