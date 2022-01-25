package main

import (
	"errors"
	"fmt"
)

func main() {
	var name string
	name = "Russell"
	fmt.Println(comprobar(name))

}
func comprobar(n string) (_string, err error) {

	if n == "Russell" {
		err = errors.New("el nombre Ingresado ya esta usado.")
		_string = errors.New("vuelve a intentar")
	} else {
		fmt.Println(n)
	}
	return
}
