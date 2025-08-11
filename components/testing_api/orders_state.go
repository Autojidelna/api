package testingapi

import (
	"errors"
	"time"
)

var usersCredits map[string]float64
var usersOrders map[string]map[string]int

func initOrdersState() {
	usersCredits = make(map[string]float64)
	usersOrders = make(map[string]map[string]int)

	for _, username := range getAllUsers() {
		// INIT credits
		credit, _ := getUserBaseCredit(username) // we can be sure the user exists here
		usersCredits[username] = credit
		// INIT orders
		initialDay := time.Now().AddDate(0, 0, -LUNCH_GEN_PAST_DAYS)
		usersOrders[username] = make(map[string]int)
		for i := 0; i <= LUNCH_GEN_FUTURE_DAYS+LUNCH_GEN_PAST_DAYS; i++ {
			currentDay := initialDay.AddDate(0, 0, i)
			currentDayString := currentDay.Format(DATE_FORMAT_YYYY_DD_MM)
			usersOrders[username][currentDayString] = 0
		}
	}
}

func getUserOrder(username string, date time.Time) (int, error) {
	userOrderData, ok := usersOrders[username]
	if !ok {
		return -1, errors.New("User not found")
	}
	dateString := date.Format(DATE_FORMAT_YYYY_DD_MM)
	return userOrderData[dateString], nil

}

func setUserOrder(username string, date time.Time, newOrder int) (bool, error) {
	userOrderData := usersOrders[username]
	credits, ok := usersCredits[username]
	if !ok {
		return false, errors.New("User not found")
	}

	dateString := date.Format(DATE_FORMAT_YYYY_DD_MM)
	prevOrder, ok := userOrderData[dateString]
	if !ok {
		return false, nil
	}

	prevMealPrice := 0.0
	if prevOrder != 0 {
		meals, ok := getLunchesDay(date)
		if !ok {
			return false, nil
		}
		prevMealPrice = meals.Jidla[prevOrder-1].Cena // we use 0 for no meal selected, so indexes are offset
	}
	newOrderPrice := 0.0
	if newOrder != 0 {
		meals, ok := getLunchesDay(date)
		if !ok {
			return false, nil
		}
		newOrderPrice = meals.Jidla[newOrder-1].Cena // we use 0 for no meal selected, so indexes are offset
	}

	credits += prevMealPrice - newOrderPrice
	if credits < 0 {
		return false, nil
	}

	usersCredits[username] = credits
	userOrderData[dateString] = newOrder
	usersOrders[username] = userOrderData

	return true, nil

}

func getUserCredit(username string) (float64, error) {
	credits, ok := usersCredits[username]
	if !ok {
		return 0.0, errors.New("User not found")
	}
	return credits, nil
}
