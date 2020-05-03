package examples

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Product struct {
	UniqueID             string   `json:"uniqueID"`
	PartNumber           string   `json:"partNumber"`
	Name                 string   `json:"name"`
	FullImage            string   `json:"fullImage"`
	Images               []string `json:"images"`
	ShortDescription     string   `json:"shortDescription"`
	LongDescription      string   `json:"longDescription"`
	XCatEntryCategory    string   `json:"xCatEntryCategory"`
	ProductString        string   `json:"productString"`
	IsMarketplaceProduct bool     `json:"isMarketplaceProduct"`
	URL                  string   `json:"url"`
	ThumbnailImage       string   `json:"thumbnailImage"`
	Single               bool     `json:"single"`
}

type MealsResponse struct {
	Meals []struct {
		IDMeal            string      `json:"idMeal"`
		StrMeal           string      `json:"strMeal"`
		StrDrinkAlternate interface{} `json:"strDrinkAlternate"`
		StrCategory       string      `json:"strCategory"`
		StrArea           string      `json:"strArea"`
		StrInstructions   string      `json:"strInstructions"`
		StrMealThumb      string      `json:"strMealThumb"`
		StrTags           interface{} `json:"strTags"`
		StrYoutube        string      `json:"strYoutube"`
		StrIngredient1    string      `json:"strIngredient1"`
		StrIngredient2    string      `json:"strIngredient2"`
		StrIngredient3    string      `json:"strIngredient3"`
		StrMeasure1       string      `json:"strMeasure1"`
		StrMeasure2       string      `json:"strMeasure2"`
		StrMeasure3       string      `json:"strMeasure3"`
		StrSource         string      `json:"strSource"`
		DateModified      interface{} `json:"dateModified"`
	} `json:"meals"`
}

func Requests() {

	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)

	go func() {
		fmt.Println("iniciando: cinco")

		resp, err := http.Get("https://simple.ripley.cl/api/v2/products/by-id/11792925")

		if err != nil {
			fmt.Println("error: cinco")
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		prod := Product{}
		jsonErr := json.Unmarshal(body, &prod)

		if jsonErr != nil {
			fmt.Println("error: cinco")
		}

		fmt.Println("despertando: ", prod.Name)
		c1 <- "cinco"
	}()
	go func() {
		fmt.Println("iniciando: siete")

		resp, err := http.Get("https://www.themealdb.com/api/json/v1/1/random.php")

		if err != nil {
			fmt.Println("error: siete")
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		mealsResponse := MealsResponse{}
		jsonErr := json.Unmarshal(body, &mealsResponse)

		if jsonErr != nil {
			fmt.Println("error: cinco")
		}

		fmt.Println("despertando: ", mealsResponse.Meals[0].StrMeal)
		c2 <- "siete"
	}()
	go func() {
		fmt.Println("iniciando: tres")
		time.Sleep(time.Second * 3)

		resp, err := http.Get("https://www.themealdb.com/api/json/v1/1/random.php")

		if err != nil {
			fmt.Println("error: siete")
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		mealsResponse := MealsResponse{}
		jsonErr := json.Unmarshal(body, &mealsResponse)

		if jsonErr != nil {
			fmt.Println("error: tres")
		}

		fmt.Println("despertando: ", mealsResponse.Meals[0].StrMeal)

		c3 <- "tres"
	}()

	// Vamos a usar `select` para esperar ambos valores
	// simultaneamente, y mostraremos cada uno conforme llegue.
	for i := 0; i < 3; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("recibido", msg1)
		case msg2 := <-c2:
			fmt.Println("recibido", msg2)
		case msg3 := <-c3:
			fmt.Println("recibido", msg3)
		}
	}
}
