package testingapi

import (
	"fmt"
	"html/template"
	"strconv"
	"time"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func buildLunches(date time.Time, username string) template.HTML {
	// Skip Weekends
	if int(date.Weekday()) == 0 || int(date.Weekday()) == 6 {
		return template.HTML(LUNCH_UNAVAILABLE_STRING)
	}

	dayMeals, ok := getLunchesDay(date)
	if !ok {
		return template.HTML(LUNCH_UNAVAILABLE_STRING)
	}
	lunchString := buildMeals(username, date, dayMeals.Jidla)

	return template.HTML(lunchString)
}

func buildMeals(username string, date time.Time, meals []Meal) string {
	// Begin Main Wrapper
	dateString := date.Format("2006-01-02")
	mealsString := fmt.Sprintf(
		`<div id="orderContent%s" class="orderContent"><div class="jidelnicekMain">`,
		dateString,
	)
	// Generate Meal HTML
	for index, mealItem := range meals {
		// Begin Item Wrapper
		mealsString += `<div class="jidelnicekItem " role="group"><div class="jidelnicekItemWrapper">`
		// Meal Interaction Primary - Order/Cancel
		/// State Logic
		orderState, err := getUserOrder(username, date)
		if err != nil {
			fmt.Println("This shouldn't happen, somehow username doesn't exist, but the sessionId matches?! buildMeals")
			orderState = 0
		}
		mealIndex := index + 1
		if int(time.Since(date).Hours()) < -ORDER_CUTOFF_HOURS {
			mealItem.LzeObjednat = true
		}
		if mealIndex == orderState {
			mealItem.Objednano = true
		} else {
			mealItem.Objednano = false
		}

		/// Frontend Logic
		state, action := gatherStateAction(mealItem)
		printer := message.NewPrinter(language.Czech)
		priceString := printer.Sprintf("%.2f", mealItem.Cena)
		orderConfirmString := ""
		if state == "ordered" {
			orderConfirmString = "Máte objednáno"
		}
		mealLink := buildMealLink(username, date, mealItem, mealIndex)

		mealsString += fmt.Sprintf(
			`
			<div class="jidWrapLeft">
				<a href="#" class="btn button-link button-link-main maxbutton %s"
					onClick="%s"
					role="button">
					<span class="button-link-align">%s<span
						style="color:black">&nbsp;1&nbsp;ks</span></span>
					<span class="smallBoldTitle button-link-align">%s</span><span
						class="button-link-align">za</span>
					<span class="important warning button-link-align"
						title="Cena objednaného jídla">%s&nbsp;Kč</span>
					<span class="button-link-tick">
						<i class="fa fa-check fa-2x"
							title="%s<b>&nbsp;1&nbsp;ks</b> á <b>%s&nbsp;Kč</b>"></i>
					</span>
				</a>
			</div>`,
			// Last Part is diff when not ordered, not just the orderConfirmString, but it hopefully isnt important
			state, mealLink, action, mealItem.Varianta, priceString, orderConfirmString, priceString,
		)
		// Meal Name and Allergens
		mealsString += fmt.Sprintf(
			`
			<div class="jidWrapCenter" id="menu-3-day-%s">
				%s
				<span
				class="textGrey">%s</span>
				<br>
			</div>`,
			dateString, mealItem.Nazev, buildAllergens(mealItem.Alergeny),
		)
		// Meal Interaction Secondary - Burza
		//!! Not Implemented Yet
		mealsString += fmt.Sprintf(
			`
			<div id="icons28428" class="icons jidWrapRight">
				<i class="far fa-clock fa-2x inlineIcon"
					title="výdej&nbsp;od:&nbsp;<b>11:00:00</b>&nbsp;do:&nbsp;<b>14:40:00</b><br/> objednat&nbsp;do:&nbsp;<b>20.01.2025 15:00:00</b><br/> zrušit&nbsp;do:&nbsp;<b>20.01.2025 15:00:00</b>"></i>
			</div>`,
		)
		// End Item Wrapper
		mealsString += `</div></div>`
	}
	// End Main Wrapper
	mealsString += `</div></div>`

	return mealsString
}

func gatherStateAction(meal Meal) (string, string) {
	stateString := ""  // ordered | enabled | disabled
	actionString := "" // (nelze) zrušit | (nelze) objednat | přeobjednat
	if !meal.LzeObjednat {
		stateString = "disabled"
		if meal.Objednano {
			actionString = "nelze zrušit"
		} else {
			actionString = "nelze objednat"
		}
	} else {
		stateString = "enabled"
		if meal.Objednano {
			stateString = "ordered"
			actionString = "zrušit"
		} else {
			actionString = "objednat"
		}
	}
	return stateString, actionString
}

func buildMealLink(username string, date time.Time, meal Meal, mealIndex int) string {
	// There are 3 types: "make" - to order, "delete" - to cancel, "reorder" - to reorder
	dateString := date.Format("2006-01-02")
	orderState, err := getUserOrder(username, date)
	if err != nil {
		fmt.Println("This shouldn't happen, somehow username doesn't exist, but the sessionId matches?! buildMealLink")
		orderState = 0
	}
	transactionType := "make"
	if orderState > 0 {
		if orderState == mealIndex {
			transactionType = "delete"
		} else {
			transactionType = "reorder"
		}
	}
	return fmt.Sprintf(
		`ajaxOrder(this, 'db/dbProcessOrder.jsp?time=1737310888909&amp;token=;&amp;ID=%s&amp;day=%s&amp;type=%s&amp;week=&amp;terminal=false&amp;keyboard=false&amp;printer=false', '2025-01-21', 'ordered')`,
		strconv.Itoa(mealIndex), dateString, transactionType,
	)
}

func buildAllergens(allergens []Allergen) string {
	allergensString := `(`
	for _, allergenItem := range allergens {
		allergensString += fmt.Sprintf(
			`<span 
            title="%s"
            class="textGrey">%s</span><span>, </span>`,
			allergenItem.Popis, allergenItem.Nazev,
		)
	}
	allergensString += `)`
	return allergensString
}

func buildBurza() {

}

func buildFooter() template.HTML {
	footerString := fmt.Sprintf(
		`<div align="center" class="textGrey noPrint">iCanteen %s | 2025-01-01 00:00:00 | &copy; <a href="https://www.z-ware.cz">Z-WARE s.r.o.</a> 2003-2021</div>`,
		BASE_VERSION,
	)

	return template.HTML(footerString)
}
