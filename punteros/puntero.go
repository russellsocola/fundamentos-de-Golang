package main

import "fmt"

func main() {
	var hola, hello *string
	tre := "Hola, quiero ser programdor "
	// imprimimos la variable de forma tradicional
	fmt.Println(tre)
	//acceso desde diferentes variables, no al contenido de la memoria si no a la direccion
	hola = &tre
	hello = &tre
	fmt.Println(hola, "\n", hello)

	*hola = " altero el contenido de esta direccion"
	//al alterar el contenido en memoria ambas variables pintaran el nuevo valor
	fmt.Println(*hello)
	fmt.Println(*hola)
}
