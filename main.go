package main

import (
	"fmt"
	//"github.com/DevOps-PcNay/Go_Licito/Utilitario"
	"log"
	"net/http" // Servidor Web
	"net/url"  // Usado para generar la URL
)

func main() {
	// Definiendo la ruta principal

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Mi pagina Principal, utilizando Air para autoejecutarse ")

		// Creando la ruta.
		// Se agrega los datos de la ruta.
		host := "www.my-empresa.com" // "localhost:3034"
		protocolo := "http"
		ruta_nueva := "/carrito"
		parametros := map[string]string{"id": "1", "nombre": "Galleta"}
		Generar_url(ruta_nueva, host, protocolo, nil)

		log.Println(Generar_url(ruta_nueva, host, protocolo, parametros))

		/*

			// url.Parse ; url = Es un paquete de Go.
			// Creando la URL
			ruta_nueva, err := url.Parse("/carrito")

			if err != nil {
				panic("Ocurrio Un Error")
			}

			// Definiendo las opciones requeridas para crear la "ruta"
			ruta_nueva.Host = host
			ruta_nueva.Scheme = protocolo

			// Agregando los parametros.
			nuevo_parametro := ruta_nueva.Query()
			nuevo_parametro.Add("id", "5")
			nuevo_parametro.Add("nombre", "Teclados")

			// Agregando a la URL el nuevo parametro creado.
			ruta_nueva.RawQuery = nuevo_parametro.Encode()

			// Imprime la URL con los parmetros
			log.Println(ruta_nueva.String())
		*/

	})

	http.HandleFunc("/productos", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Pagina De Productos")
		// Para imprimir la URL pero sin el dominio.
		log.Println("Obteniendo la URL ", r.URL)
		log.Println("Obteniendo los parametros ", r.URL.RawQuery)
		// Obteniendo el valor de las variables que se envian en los parametros
		log.Println("Id Producto = ", r.URL.Query().Get("idProduct"))
		log.Println("Nombre = ", r.URL.Query().Get("nombre"))

		// Agregando una variable de http: precio = 10
		// Se utilizan los mapas
		nuevo_valor := r.URL.Query()
		nuevo_valor.Add("precio", "10")
		log.Println("Contenido Variable Http nueva ", nuevo_valor.Encode())

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

func Generar_url(parametro, host, protocolo string, url_generar map[string]string) string {
	// Se agrega los datos de la ruta.
	ruta_nueva, err := url.Parse(parametro)

	if err != nil {
		panic("Ocurrio Un Error")
	}
	ruta_nueva.Host = host
	ruta_nueva.Scheme = protocolo

	// Agregando los parametros.
	nuevo_parametro := ruta_nueva.Query()

	for key, value := range url_generar {
		nuevo_parametro.Add(key, value)

	}
	//nuevo_parametro.Add("id", "5")
	//nuevo_parametro.Add("nombre", "Teclados")

	// Agregando a la URL el nuevo parametro creado.
	ruta_nueva.RawQuery = nuevo_parametro.Encode()

	// Imprime la url modificada
	//log.Println(ruta_nueva.String())

	return ruta_nueva.String()

} // func generar_url(parametro, host, pro
