package main

import (
	"fmt"

	"github.com/russellsocola/gocurso/interfaces/coches"
)

func main() {
	defer fmt.Println("La eh liado, la paqueteria me esta confundiendo.")
	var (
		f coches.Ferrari //declaro la variable F como tipo (paquete+type""")
		l coches.Lanborgini
	)
	//uso la variable antes declaradas
	f.Modelo = "nuevo"
	l.Modelo = "viejo"

	Comprar(f)
	Comprar(l)
}

//Comprar es una funcion publica
func Comprar(q coches.Vehiculos) {
	q.Arrancar()
}
