package main

import (
	"fmt"
)

func main() {
	saludar()
	saludar_dinamico("Jose", "Lopez", 40)
}

func saludar() {
	fmt.Println("Hola Mundo")
}

func saludar_dinamico(nombre, apellido string, edad int8) {
	fmt.Println("Hola", nombre, apellido, " edad ", edad)
}
