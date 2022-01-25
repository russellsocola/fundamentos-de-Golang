package main

import "fmt"

func main() {
	//en go los array son de un tamalo fijo y mismo tipo de datos
	var alumnos [4]string
	alumnos[0] = "Socola"
	alumnos[1] = "Manrrique"
	alumnos[2] = "Anhuaman"
	alumnos[3] = "Sotomayor"

	fmt.Println(alumnos)
	//forma abreviada de declarar un arreglo
	notas := [4]int{10, 12, 14, 20}

	fmt.Println(alumnos[0], notas[0])
	// se usa para contar la cantidad de elementos en un arreglo
	tama := len(alumnos)

	fmt.Println(tama)

	fmt.Println(alumnos[0:2])

}
