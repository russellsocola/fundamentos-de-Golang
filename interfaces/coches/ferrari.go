package coches

import "fmt"

type Ferrari struct {
	Modelo string
}

//Arrancar() es una funcion publica
func (f Ferrari) Arrancar() {
	fmt.Println("Soy un Ferrari")
}
