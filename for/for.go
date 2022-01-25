package main

import "fmt"

func main() {
	/*

		var n int
		name := "russell"
		n = 6

		for i := 0; i < n; i++ {
			fmt.Println(i)
		}
		fmt.Scanln(&name)
		for i := 0; *&name == name; i++ {
			fmt.Println(name)
			break
		}
		for i := 5; i >= 0; i-- {
			if i == 3 {
				fmt.Println("Se salta el valor 3")
				continue
			}
			fmt.Println(i)
		}
		fmt.Println("")
	*/
	//Uso una forma simplificada de para rellenar mi arreglo
	mat := [3][3]int{{18, 13, 13}, {23, 23, 43}, {23, 32, 54}}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println("[", i, ";", j, "]") //imprimo la posicion en la que se encuentra recorrinendo el arreglo
			fmt.Println(mat[i][j])           //imprimo el valor guardado en la pocicion
		}
	}
	// usare un nuevo for con el fin de imorimir los valores en forma de matriz de 3 x 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(mat[i][j], " ")
		}
		fmt.Println("")
	}
}
