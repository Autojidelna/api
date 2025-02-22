package testingapi

import (
	"fmt"
	"html/template"
	"os"
	"time"

	"encoding/json"
)

type LunchMeal struct {
	Jidla []Meal `json:"jidla"`
}

type Meal struct {
	Nazev    string `json:"nazev"`
	Varianta string `json:"varianta"`
}

type Allergen struct {
}

func buildLunches(date time.Time) template.HTML {
	fmt.Println(date)
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
	fmt.Println(lunchId)
	//
	lunchString := ""
	fileBytes, err := os.ReadFile("json/meals.json")

	if err != nil {
		lunchString = "Error: Read File Failed"
		fmt.Println("Error when reading file")
	}

	var lunches map[string]LunchMeal
	err = json.Unmarshal(fileBytes, &lunches)
	if err != nil {
		lunchString = "Error: Invalid JSON"
		fmt.Println("Error when reading file")
	}

	lunchString = string(fileBytes)
	lunchIdString := fmt.Sprintf("%d", lunchId)
	dayMeals := lunches[lunchIdString]
	fmt.Println("Jidla")
	lunchString = dayMeals.Jidla[0].Nazev

	return template.HTML(lunchString)
}

func buildMeals(date time.Time, meals []LunchMeal) string {
	mealsString := fmt.Sprintf(
		`<div id="orderContent%s" class="orderContent">
    <div class="jidelnicekMain">
        <div class="jidelnicekItem " role="group">
            <div class="jidelnicekItemWrapper">
                <div class="jidWrapLeft">
                    <a href="#" class="btn button-link button-link-main maxbutton ordered"
                        onClick="ajaxOrder(this, 'db/dbProcessOrder.jsp?time=1737310888909&amp;token=;ID=755744&amp;day=2025-01-21&amp;type=delete&amp;week=&amp;terminal=false&amp;keyboard=false&amp;printer=false', '2025-01-21', 'ordered');"
                        role="button">
                        <span class="button-link-align">zrušit<span class="warning"
                                style="color:black">&nbsp;1&nbsp;ks</span></span>
                        <span class="smallBoldTitle button-link-align">Oběd 1</span><span
                            class="button-link-align">za</span>
                        <span class="important warning button-link-align"
                            title="Cena objednaného jídla">42,00&nbsp;Kč</span>
                        <span class="button-link-tick">
                            <i class="fa fa-check fa-2x"
                                title="Máte objednáno<b>&nbsp;1&nbsp;ks</b> á <b>42,00&nbsp;Kč</b>"></i>
                        </span>
                    </a>
                </div>
                <div class="jidWrapCenter" id="menu-3-day-2025-01-21">
                    Polévka hovězí s játrovou rýží * , hovězí maso námořnické *směs , přidává se
                    víno a vejce , rýže dušená * , čaj * , sirup , voda , salátový bar <span
                        class="textGrey">(<span
                            title="Obiloviny obsahující lepek - nejedná se o celiakii, výrobky z nich"
                            class="textGrey">Obiloviny</span><span>, </span><span
                            title="Vejce a výrobky z nich - patří mezi potraviny ohrožující život"
                            class="textGrey">Vejce</span><span>, </span><span
                            title="Mléko  a  výrobky  z  něj - patří mezi potraviny ohrožující život"
                            class="textGrey">Mléko</span><span>, </span><span
                            title="Skořápkové  plody a výrobky z nich - jedná se o všechny druhy ořechů"
                            class="textGrey">Skořápkové plody</span><span>, </span><span
                            title="Celer a výrobky z  něj" class="textGrey">Celer</span><span>,
                        </span><span title="Hořčice a  výrobky z ní"
                            class="textGrey">Hořčice</span>)</span>
                    <br>
                </div>
                <div id="icons28428" class="icons jidWrapRight">
                    <i class="far fa-clock fa-2x inlineIcon"
                        title="výdej&nbsp;od:&nbsp;<b>11:00:00</b>&nbsp;do:&nbsp;<b>14:40:00</b><br/> objednat&nbsp;do:&nbsp;<b>20.01.2025 15:00:00</b><br/> zrušit&nbsp;do:&nbsp;<b>20.01.2025 15:00:00</b>"></i>
                </div>
            </div>
        </div>
        <div class="jidelnicekItem " role="group">
            <div class="jidelnicekItemWrapper">
                <div class="jidWrapLeft">
                    <a href="#" class="btn button-link button-link-main maxbutton enabled"
                        onClick="ajaxOrder(this, 'db/dbProcessOrder.jsp?time=1737310888909&amp;token=;ID=5&amp;day=2025-01-21&amp;type=reorder&amp;week=&amp;terminal=false&amp;keyboard=false&amp;printer=false', '2025-01-21', 'enabled');"
                        role="button">
                        <span class="button-link-align">přeobjednat</span>
                        <span class="smallBoldTitle button-link-align">Oběd 2</span><span
                            class="button-link-align">s doplatkem</span>
                        <span class="important warning button-link-align"
                            title="Rozdíl ceny oproti stávající objednávce&nbsp;Cena při objednání jídla:&nbsp;42,00&nbsp;Kč">0,00&nbsp;Kč</span>
                    </a>
                </div>
                <div class="jidWrapCenter" id="menu-4-day-2025-01-21">
                    Polévka hovězí s játrovou rýží * , kuře na způsob bažanta * , brambory
                    šťouchané , čaj * , sirup , voda , salátový bar , pro celiaky-polévka s rýží
                    a masem <span class="textGrey">(<span
                            title="Obiloviny obsahující lepek–nejedná se o celiakii, výrobky z nich"
                            class="textGrey">Obiloviny</span><span>, </span><span
                            title="Vejce a výrobky z nich - patří mezi potraviny ohrožující život"
                            class="textGrey">Vejce</span><span>, </span><span
                            title="Mléko  a  výrobky  z  něj - patří mezi potraviny ohrožující život"
                            class="textGrey">Mléko</span><span>, </span><span
                            title="Skořápkové  plody a výrobky z nich – jedná se o všechny druhy ořechů"
                            class="textGrey">Skořápkové plody</span><span>, </span><span
                            title="Celer a výrobky z  něj" class="textGrey">Celer</span><span>,
                        </span><span title="Hořčice a  výrobky z ní"
                            class="textGrey">Hořčice</span>)</span>
                    <br>
                </div>
                <div id="icons28429" class="icons jidWrapRight">
                    <i class="far fa-clock fa-2x inlineIcon"
                        title="výdej&nbsp;od:&nbsp;<b>11:00:00</b>&nbsp;do:&nbsp;<b>14:40:00</b><br/> objednat&nbsp;do:&nbsp;<b>20.01.2025 15:00:00</b><br/> zrušit&nbsp;do:&nbsp;<b>20.01.2025 15:00:00</b>"></i>
                </div>
            </div>
        </div>
    </div>
</div>`, date.Format("2006-01-02"))
	return mealsString
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
