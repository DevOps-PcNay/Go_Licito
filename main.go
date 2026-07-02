package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Definiendo la ruta principal
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Mi pagina Principal, utilizando Air para autoejecutarse ")
	})
	http.HandleFunc("/productos", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Pagina De Productos")
	})

	http.HandleFunc("/categorias", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Pagina De Categorias")
		//http.Redirect(w, r, "/", 301)
	})

	http.HandleFunc("/noencontro", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "Pagina De Categorias")
		http.NotFound(w, r)
	})

	http.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "Pagina De Categorias")
		http.Error(w, "Error En El Servidor", 500)
	})

	// Ejecutando el servidor Web de Go
	log.Println("Servidor Corriendo ...")
	http.ListenAndServe(":3034", nil)

}
