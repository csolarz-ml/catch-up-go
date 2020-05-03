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
		fmt.Println("1.- Ejecutar funciones en paralelo y esperar que todos terminen")
		fmt.Println("2.- Ejecutar requests en paralelo y esperar que todos terminen")
		fmt.Println("3.- Ejecutar funciones en paralelo y esperar que termine solo uno")

		var input string
		fmt.Scanln(&input)
		fmt.Print(input)

		fmt.Println()

		if input == "0" {
			break
		} else if input == "1" {
			start := time.Now()

			examples.WaitAllAsync()

			elapsed := time.Since(start)
			fmt.Printf("WaitAllAsync se ejecuto en %s", elapsed)
		} else if input == "2" {
			start := time.Now()

			examples.WaitAllRequests()

			elapsed := time.Since(start)
			fmt.Printf("WaitAnyAsync se ejecuto en %s", elapsed)
		} else if input == "3" {
			start := time.Now()

			examples.WaitAnyAsync()

			elapsed := time.Since(start)
			fmt.Printf("WaitAnyAsync se ejecuto en %s", elapsed)
		}
	}
}
