package coches

import "fmt"

type Lanborgini struct {
	Modelo string
}

//Lanborgini es una funcion publica.
func (l Lanborgini) Arrancar() {
	fmt.Println("Soy un Lamborgini")
}
