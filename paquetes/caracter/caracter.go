package caracter

import "fmt"

func Caracter(string) { // declaro la funcion con Mayuscula para que pueda ser usada desde fuera
	var (
		peso        float64
		colorpiel   string
		tipocabello string
		talla       float64
	)
	peso = 93.4
	colorpiel = "Moreno"
	tipocabello = "Ondulado"
	fmt.Println("Ingrese la talla ")
	fmt.Scanln(&talla)
	masa1, numero := indice(peso, talla)
	genero := sexo("m")
	fmt.Printf("Mis caracteristicas son las siguientes,peso %g soy de test %s(a) mi cabello es %s,\tmi genero es %s\ty mi indice de masa corportal es %.2f y %s gracias.\n", peso, colorpiel, tipocabello, genero, numero, masa1)
}

//esta funcion solo puede ser llamada desde este archivo.
func indice(peso float64, talla float64) (string, float64) { //esta es una funcion privada
	var (
		total float64
		masa  string
	)
	total = peso / (talla * talla)

	switch {
	case total < 18.5:
		masa = "bajo de lo normal"
	case total < 24.9:
		masa = "Normal"
	case total < 29.9:
		masa = "superior a lo normal"
	case total >= 30.0:
		masa = "Obesidad"
	default:
		masa = "Calculo invalido vuelva a intentar"
	}
	return masa, total

}
func sexo(genero string) string {
	switch genero {
	case "m":
		genero = "Masculino"
	case "f":
		genero = "Femenino"
	default:
		fmt.Println("vuelva a intentar")
	}
	return genero
}
