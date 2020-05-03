package examples

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/csolarz-ml/catch-up-go/examples/models"
)

func WaitAllRequests() {
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)

	go func() {
		fmt.Println("iniciando: llamada ripley")

		resp, err := http.Get("https://simple.ripley.cl/api/v2/products/by-id/11792925")

		if err != nil {
			fmt.Println("error: llamada ripley")
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		prod := models.Product{}
		jsonErr := json.Unmarshal(body, &prod)

		if jsonErr != nil {
			fmt.Println("error: llamada ripley")
		}

		fmt.Println("fin llamada ripley: ", prod.Name)
		c1 <- "ripley"
	}()

	go func() {
		fmt.Println("iniciando: llamada meal api")

		resp, err := http.Get("https://www.themealdb.com/api/json/v1/1/random.php")

		if err != nil {
			fmt.Println("error: llamada meal api")
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		mealsResponse := models.MealsResponse{}
		jsonErr := json.Unmarshal(body, &mealsResponse)

		if jsonErr != nil {
			fmt.Println("error: llamada meal api")
		}

		fmt.Println("finalizado llamada meal api: ", mealsResponse.Meals[0].StrMeal)
		c2 <- "meal-api"
	}()

	go func() {
		fmt.Println("iniciando: llamada meal api (5s delay)")
		time.Sleep(time.Second * 5)

		resp, err := http.Get("https://www.themealdb.com/api/json/v1/1/random.php")

		if err != nil {
			fmt.Println("error: llamada meal api (5s delay)")
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		mealsResponse := models.MealsResponse{}
		jsonErr := json.Unmarshal(body, &mealsResponse)

		if jsonErr != nil {
			fmt.Println("error: llamada meal api (5s delay)")
		}

		fmt.Println("fin llamada meal api (5s delay): ", mealsResponse.Meals[0].StrMeal)

		c3 <- "meal-api (5s)"
	}()

	// Vamos a usar `select` para esperar ambos valores
	// simultaneamente, y mostraremos cada uno conforme llegue.
	for i := 0; i < 3; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("fin func: ", msg1)
		case msg2 := <-c2:
			fmt.Println("fin func: ", msg2)
		case msg3 := <-c3:
			fmt.Println("fin func: ", msg3)
		}
	}
}
