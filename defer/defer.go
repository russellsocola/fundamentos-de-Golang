package main

import "fmt"

func main() {
	fmt.Println("Establecer enlace con el servidor")
	defer conec() //este sera el ultimo defer en ser pintado
	//
	fmt.Println("Pedir datos del servidor")
	defer extraer() //sera el antepenultimo defer en ser pintado

	defer cerra() //este sera el primer defer en ser pintado
}

func conec() {
	fmt.Println("Se establecio conecciòn")
}
func extraer() {
	fmt.Println("Se extraen datos de la base ")
}
func cerra() {
	fmt.Println("Se a concluido la conecciòn")
}
