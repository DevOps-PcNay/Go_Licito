package main

import (
	"errors"
	"fmt"
)

func main() {
	resultado, err := operacion(34, 50, "resta")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultado)
	}

}

func operacion(n1, n2 int, calculo string) (resultado int, err error) {

	if n1 < 0 {
		err = errors.New("Error : El numero 1 es negativo ")
	} else if n2 < 0 {
		err = errors.New("Error : El nuemro 2 es negativo ")
	}
	// Nos muestra advertencia, indicando que se debe hacer con "switch"
	// En la asignacion de la variable se omite ":=" porque se esta retornando en la cabecera d ela funcion.

	switch {
	case calculo == "suma":
		resultado = n1 + n2
	case calculo == "resta":
		resultado = n1 - n2
	case calculo == "multiplicacion":
		resultado = n1 * n2
	default:
		resultado = 0
	}

	return
}
