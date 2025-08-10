package testingapi

var profileOrders map[string]int
var profileCredit float64

func initOrdersState() {
	profileOrders = make(map[string]int)
	profileCredit = BASE_CREDIT

	// Example: set today's date
	//today := time.Now().Format("2006-01-02")
	//profileOrders[today] = 2 // value between 0 and 3

	//profileOrders["2025-08-09"] = 3

	// Access a value
	//if val, ok := profileOrders[today]; ok {
	//fmt.Printf("Value for %s: %d\n", today, val)
	//}

	// Print the whole map
	//fmt.Println(profileOrders)
}
