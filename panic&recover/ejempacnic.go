package main

import (
	"fmt"
	"reflect" //importo el paquete que me permite saber con que tipo de variable estoy trabajando
	"strconv" //me permite convertir de cadena a entero(float,boll,etc)y viseversa.
)

func main() {
	text := 2
	//declaro el valory tipo de variable
	/*
		defer func() {
			recuba := recover()
			if recuba != nil {
				fmt.Println(recuba)
			}
		}()
	*/
	fmt.Println(reflect.TypeOf(text)) //pinto el tipo de variable(entero*int*)
	cade := strconv.Itoa(text)        //la cambio a cadena
	fmt.Println(reflect.TypeOf(cade)) //pinto el nuevo tippo(cadena*string*)
	fmt.Println(cade)                 //pinto el contenido de la variable
	cade1, err := strconv.Atoi(cade)

	if err != nil {
		panic(err)
	}
	fmt.Println(cade1, err)
	fmt.Println("fin")
	fmt.Println(cadenaAnumero("2"))
	fmt.Println(cadenaAnumero("j"))
	fmt.Println("se ejecuta si funciona el panic")
}
func cadenaAnumero(cadena string) int {
	defer func() {
		recuperado := recover()
		if recuperado != nil {
			fmt.Println(recuperado)
		}

	}()
	numero, err := strconv.Atoi(cadena)
	if err != nil {
		panic(err)
	}
	return numero
}
