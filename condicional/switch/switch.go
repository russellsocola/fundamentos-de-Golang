package main

import "fmt"

func main() {

	var n int

	fmt.Println("Ingrese un numero del 1-5")
	fmt.Scanln(&n)

	switch n {
	case 1:
		fmt.Println("Ingreso el numero UNO")
	case 2:
		fmt.Println("Ingreso el numero DOS")
	case 3:
		fmt.Println("Ingreso el numero TRES")
	case 4:
		fmt.Println("Ingreso el numero CUATRO")
	case 5:
		fmt.Println("Ingreso el numero CINCO")
	default:
		fmt.Println("El numero ingresado es un numero invalido")
	}
}
