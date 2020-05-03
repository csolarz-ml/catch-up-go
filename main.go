package main

import (
	"fmt"
	"time"

	"github.com/csolarz-ml/catch-up-go/examples"
)

func countSecs(time time.Time) int {
	return 1
}

func main() {
	for {
		fmt.Println()
		fmt.Println("Ingrese opciones")
		fmt.Println("1.- Ejecutar funciones en paralelo")
		fmt.Println("2.- Ejecutar requests en paralelo")

		var input string
		fmt.Scanln(&input)
		fmt.Print(input)

		fmt.Println()

		if input == "0" {
			break
		} else if input == "1" {
			start := time.Now()
			examples.Async()
			elapsed := time.Since(start)
			fmt.Printf("Async took %s", elapsed)
		} else if input == "2" {
			start := time.Now()
			examples.Requests()
			elapsed := time.Since(start)
			fmt.Printf("Requests took %s", elapsed)
		}
	}
}
