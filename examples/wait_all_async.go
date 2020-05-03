package examples

import (
	"fmt"
	"time"
)

func WaitAllAsync() {
	// Para nuestro ejemplo vamos a utilizar select con tres canales
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)

	// Cada canal va a recibir un valor despues de cierto tiempo,
	// para simular operaciónes bloqueadas de procesos externos
	// en goroutines concurrentes.
	go func() {
		fmt.Println("iniciando: channel 5s")
		time.Sleep(time.Second * 5)
		fmt.Println("fin: channel 5s")
		c1 <- "channel 5s"
	}()

	go func() {
		fmt.Println("iniciando: channel 7")
		time.Sleep(time.Second * 7)
		fmt.Println("fin: channel 7")
		c2 <- "channel 7s"
	}()

	go func() {
		fmt.Println("iniciando: channel 3")
		time.Sleep(time.Second * 3)
		fmt.Println("fin: channel 3")
		c3 <- "channel 3s"
	}()

	// Vamos a usar `select` para esperar ambos valores
	// simultaneamente, y mostraremos cada uno conforme llegue.
	for i := 0; i < 3; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("fin: ", msg1)
		case msg2 := <-c2:
			fmt.Println("fin: ", msg2)
		case msg3 := <-c3:
			fmt.Println("fin: ", msg3)
		}
	}
}
