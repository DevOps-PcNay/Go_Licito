package main

import "fmt"

func main() {
	resultado1 := sumar(6, 7)
	fmt.Println("La suma es ", resultado1)
	resultado2 := sumar(10, 20)
	fmt.Println("La suma es ", resultado2)

	// Funciones con operaciones.
	resultado_operacion := operacion(10, 20, "multiplicacion")
	fmt.Println("El resultado es", resultado_operacion)

	// Retorno de varios valores.
	suma, resta, multiplicacion := super_operaciones(20, 40)

	fmt.Println("Retorno de tres valores : ", suma, resta, multiplicacion)

}

func sumar(numero1, numero2 int) int {
	suma := numero1 + numero2
	return suma
}

func operacion(n1, n2 int, calculo string) int {
	/*
		if calculo == "suma" {
			return n1 + n2
		} else if calculo == "resta" {
			return n1 - n2
		} else if calculo == "multiplicacion" {
			return n1 * n2
		} else {
			return 0
		}
	*/

	// Nos muestra advertencia, indicando que se debe hacer con "switch"
	switch calculo {
	case "suma":
		return n1 + n2
	case "resta":
		return n1 - n2
	case "multiplicacion":
		return n1 * n2
	default:
		return 0
	}
}

// Retornando 3 valores:
func operaciones(n1, n2 int) (int, int, int) {
	suma := n1 + n2
	resta := n1 - n2
	multiplicacion := n1 * n2
	return suma, resta, multiplicacion
}

// No tiene implicito la palabra "return"

func super_operaciones(n1, n2 int) (suma1 int, resta1 int, multiplicacion1 int) {
	suma1 = n1 + n2
	resta1 = n1 - n2
	multiplicacion1 = n1 * n2
	return
	// Min 12:56

}
