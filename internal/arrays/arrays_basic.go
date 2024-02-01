package arrays

import "fmt"

// Arrays o matrices se pueden inicializar de la siguiente manera

var myArray [3]int // Aca inicializamos este array de enteros con una capacidad de 3 posiciones

// Inicialización sin especificar un número especifico
var cities = [...]string{"New York", "Paris", "Berlin", "Madrid"}

func PrintArray() {
	fmt.Println(myArray)
	fmt.Println("Cities:", cities)
}
