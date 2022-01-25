package main

import "fmt"

func main() {

	hero := map[string]map[string]string{
		"saitama": {
			"Oficio": "Heroe por pasatiempo",
			"Tipo":   "B",
		},
		"sonic": {
			"Oficio": "Asecino a sueldo ",
			"Tipo":   "B",
		},
	}
	if tem, Heroe := hero["sonic"]; Heroe {
		fmt.Println(tem["Oficio"], " clasificacion ", tem["Tipo"])
	}

	//
	StudentAge := make(map[string]int)

	StudentAge["Russell"] = 20
	StudentAge["Carla"] = 22
	StudentAge["Nycoll"] = 20
	StudentAge["Iria"] = 19

	fmt.Println(StudentAge)
	fmt.Println(len(StudentAge))
	fmt.Println(StudentAge["Russell"])
}
