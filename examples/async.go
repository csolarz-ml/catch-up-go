package examples

import (
	"fmt"
	"time"
)

func Async() {
	// Para nuestro ejemplo vamos a utilizar select con tres canales
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)

	// Cada canal va a recibir un valor despues de cierto tiempo,
	// para simular operaciónes bloqueadas de procesos externos
	// en goroutines concurrentes.
	go func() {
		fmt.Println("iniciando: cinco")
		time.Sleep(time.Second * 5)
		fmt.Println("despertando: cinco")
		c1 <- "cinco"
	}()

	go func() {
		fmt.Println("iniciando: siete")
		time.Sleep(time.Second * 7)
		fmt.Println("despertando: siete")
		c2 <- "siete"
	}()

	go func() {
		fmt.Println("iniciando: tres")
		time.Sleep(time.Second * 3)
		fmt.Println("despertando: tres")
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
