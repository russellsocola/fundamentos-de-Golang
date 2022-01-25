package main

import (
	"fmt"
	"time"
)

func main() {
	var h string
	go rutina()
	defer fmt.Println("Ingrese un numero")
	fmt.Scan(&h)
	fmt.Println("Ingresaste:", h)
}
func rutina() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}
