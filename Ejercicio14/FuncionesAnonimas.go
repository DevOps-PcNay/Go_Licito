package main

import (
	"fmt"
)

func main() {
	// Funciones anonima, se autojecuta con solo colocarle "()" al terminar de definir.

	func() {
		fmt.Println("Soy una funcion anomima")
	}() // Se autoejecuta la funcion

	// Asignando una funcion anonima a una variable
	saludo := func() {
		fmt.Println("Es otra funcion anonima, se asigna a una variable ")
	}

	saludo() // Usando la variable que le asigno una funcion.

	nombrecursos("Pedro Noagels", "Go Land", "Python", "NodeJs", "JavaScript")

	// Manejando punteros.
	var a int = 4
	triplicar(&a) //Se pasa por referencia, puntero,
	fmt.Println(a)

} //func main() {

// Pasando "n" valores a una funcion
func nombrecursos(profesor string, cursos ...string) {
	// cursos = Es un "slices"
	// Para recorrer el slides
	// _ = Para que no se muestre el indice.

	for _, v := range cursos {
		fmt.Println("Mi curso que enseno es ", v, "MI nombre es ", profesor)
	}

} // func nombrecursos(cursos ...string){

// Variables manaejadas por valor.
// "*" = Es desreferenciacion de la variable
func triplicar(n1 *int) {
	*n1 = *n1 * 3
	fmt.Println("El triple del numero es : ", *n1)
	// n1 = Muestra la direccion de memoria
	// *n1 = Muesta el valor que contiene
}
