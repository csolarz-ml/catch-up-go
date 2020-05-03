package examples

import (
	"fmt"
	"time"
)

func one() <-chan string {
	r := make(chan string)

	go func() {
		fmt.Println("Iniciando channel one (2s)")
		defer close(r)

		// Simulate a workload.
		time.Sleep(time.Second * 2)
		r <- "channel one: 2s"
		fmt.Println("Fin channel one (2s)")
	}()

	return r
}

func two() <-chan string {
	r := make(chan string)

	go func() {
		fmt.Println("Iniciando channel two (10s)")
		defer close(r)

		// Simulate a workload.
		time.Sleep(time.Second * 10)
		r <- "channel two: 10s"

		fmt.Println("Fin channel two (10s)")
	}()

	return r
}

func WaitAnyAsync() {
	// Vamos a usar `select` para esperar ambos valores
	// simultaneamente, y mostraremos cada uno conforme llegue.
	var r string
	select {
	case r = <-one():
	case r = <-two():
	}

	fmt.Println(r)
}
