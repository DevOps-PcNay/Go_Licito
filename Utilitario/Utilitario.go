package Utilitario

import (
	//"fmt"
	"io" // Para desplegar los datos en Bytes
	//"log"
	"net/http" //
	"net/url"  // Usado para generar la URL"
)

// Se coloca la primer letra en mayusculas "Generar_url" para que pueda ser accesada por otros modulos, paquetes

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

// Peticion "Get" del protocolo "http"
func Request_Get(metodo, url string) (resolve_get string) {

	request, err := http.NewRequest(metodo, url, nil)
	if err != nil {
		panic("Hubo un problmea con el Request ")
	}

	cliente := &http.Client{}
	resolve, err := cliente.Do(request)
	if err != nil {
		panic("Hubo un problmea con el Cliente ")
	}

	// Converte informacion Hexadecimal a caracteres legibles
	bytes, err := io.ReadAll(resolve.Body)
	if err != nil {
		panic("Error al recuperar los datos")
	}

	//log.Println("Status", resolve.Status)
	//log.Println("Body", string(bytes))

	//"resolve_get" = Se declara como vaor que retorna la funcion.
	resolve_get = string(bytes)

	return resolve_get

} // func Request(metodo, url string) {
