package arrays

import "fmt"

// Arrays o matrices se pueden inicializar de la siguiente manera

var myArray [3]int // Aca declaramos este array de enteros con una capacidad de 3 posiciones, pero no esta inicializada
// La salida sería [0 0 0]

// Inicialización sin especificar un número especifico
var cities = [...]string{"New York", "Paris", "Berlin", "Madrid"}

// Inicialización con posiciones especificas
var names = [5]string{1: "Flavio", 3: "Pepe el grillo"}

func PrintInitializationArrays() {
	fmt.Println(myArray)
	fmt.Println("Cities:", cities)
	fmt.Println(names)
}
