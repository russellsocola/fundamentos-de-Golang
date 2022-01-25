package main

import "fmt"

func main() {
	var (
		n  int
		n1 int
		n2 int
	)
	//asigno valor a las variables
	n = 2
	n1 = 3
	n2 = 8
	//proceso
	if n <= n2 {
		fmt.Println("estoy pintando desde el if")
		fmt.Println(n, n1, n2)
	} else {
		fmt.Println("estoy pintando desde el else")
		fmt.Println(n, n1, n2)
	}
}
