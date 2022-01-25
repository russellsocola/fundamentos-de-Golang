package main

import "fmt"

func main() {
	nombres := make([]string, 0)
	nombres = append(nombres, "Lufy", "Russell", "Jose") //mediante
	fmt.Println(len(nombres), cap(nombres))
	nombres = append(nombres, "Ignacio")
	fmt.Println(len(nombres), cap(nombres))
	nombres = append(nombres, "Emely")
	fmt.Println(len(nombres), cap(nombres))
	nombres = append(nombres, "Kiara")
	fmt.Println(len(nombres), cap(nombres))
	nombres = append(nombres, "Josimar")
	fmt.Println(len(nombres), cap(nombres))
	nombres = append(nombres, "Gerald")
	fmt.Println(nombres)

	var ap string
	var slicen1 []string
	slicen1 = append(slicen1, "Russell", "Nahùm", "Socola", "Reyes")

	//segunda forma de utilizar un slice
	var n int
	fmt.Scanln(&n)
	ap = slicen1[n]
	fmt.Println(ap)
	fmt.Println(len(slicen1))

}
