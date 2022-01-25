package main

import "fmt"

type Persona struct {
	nombre   string
	edad     int
	estatura float64
	otros    []string
}
type rentangle struct {
	altura float64
	ancho  float64
}

func (rec *rentangle) area() float64 {
	return rec.altura * rec.ancho
}

func main() {
	var per Persona
	per.edad = 20
	per.nombre = "Russell"
	per.estatura = 1.75
	per.otros = append(per.otros, "futuro Inge")

	fmt.Println(per)

	par1 := Persona{
		"Kevin",
		28,
		1.79,
		[]string{"null"},
	}

	fmt.Println(par1)

	//Envio datos a la estrucura
	rec2 := rentangle{2, 2}
	//pideo que la estructura envie datos a la funcion para luego pintarlos en la pantalla
	fmt.Println(rec2.area())
}
