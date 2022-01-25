package main

import (
	"log"
	"net/http"
)

func main() {
	// servidor web de forma tradicional
	//Ruta a nuestro servidor que va a ser la raiz "/" y que debe hacer al realizar esa peticion.
	http.Handle("/", http.FileServer(http.Dir("public")))
	log.Println("Ejecutado server en http://localhot:8080")

	//se carga el servidor diciendole que se sirva en el puerto
	//si el puerto esta ocupado va a devolver un error y para capturar el error lo hacemos con un log
	log.Println(http.ListenAndServe(":8080", nil))

	/*servidor web mejorado con un multiplexor(ayuda a relsolver lasa solucitudes)
	server mux resuelve mas peticiones
	mux:= http.NewServerMux()
	fs:=http.FileServer(http.Dir("public"))
	mux.Handle("/",fs)
	log.Println("Ejecutando el server en http://localhost:8080")
	log.Println(http.ListenAndServer(":8080",mux))
	*/
}
