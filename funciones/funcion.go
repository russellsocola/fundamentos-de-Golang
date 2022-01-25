package main

import "fmt"

func main() {
	suma() // estoy llamando a la funcion suma desde el main
	defer zapato("Adidas", 43.5)
	//ingresare datos en una variable

	n := 13
	n1 := 12
	n2 := 18
	n3 := 20
	//envio numeros a la funcion y me retornara e imprimira una cadena.
	fmt.Println(notas(n, n1, n2, n3))
}

// las funciones se utilizan para tener un codigo mucho mas ordenado
// las funcioes se declarand de la siguiente manera

func suma() {

	var n, m int
	n = 10
	m = 32
	fmt.Println("la suma de dos numeros es ", m+n)
}

// ahora usare parametros en las funciones
func zapato(marca string, talla float64) {
	if marca != "" {
		fmt.Println(" Tienes buenos gustos ")
		fmt.Printf("Tu talla es numero %g\n", talla)
	} else {
		fmt.Println("Vuelve a ingresar, por favor")
	}
}

// retornar una funcion

func notas(a, b, c, d int) string {
	prom := (a + b + c + d) / 4
	var calificacion string
	switch {
	case prom < 13:
		calificacion = "Desaprobado"
	case prom < 15:
		calificacion = "Aprobado por la minima"
	case prom < 18:
		calificacion = "Aprobado con buen nivel de aprobechamiento"
	case prom < 20:
		calificacion = "La Nasa te estta buscando"
	default:
		calificacion = "Ya rindete"
	}
	return calificacion
}
