package testingapi

import (
	"fmt"
	"html/template"
	"os"
	"time"
)

type LunchMeal struct {
	Jidla []Meal `json:"jidla"`
}

type Meal struct {

}

type Allergen struct {

}

func buildLunches(date time.Time) template.HTML {

	// Skip Weekends
	if int(date.Weekday()) == 0 || int(date.Weekday()) == 6 {
		return template.HTML(lunchUnavailableString)
	}
	// Protoype logic to select an id for a meal
	_, week := date.ISOWeek()
	lunchId := 0
	if week%2 == 0 {
		lunchId = 5
	}
	lunchId += int(date.Weekday())
	//
	lunchString := ""
	fileBytes, err := os.ReadFile("json/meals.json")
	if err != nil {
		lunchString = "Error"
		fmt.Println("Error when reading file")
	}
	lunchString = string(fileBytes)

	return template.HTML(lunchString)
}

func buildBurza() {

}

func buildFooter() template.HTML {
	footerString := fmt.Sprintf(
		`<div align="center" class="topMenu topMenuMinWidth bottomBar textGrey noPrint" role="contentinfo"> Změna
            jídelníčku vyhrazena. | iCanteen %s
| <span class="topMenuItem">poslední přihlášení: <span id="PoslLogin"
            style="font-weight: bold;">2025-01-01 00:00:00.0 IP: 1.1.1.1</span></span>
</div>`, baseVersion)
	return template.HTML(footerString)
}
