package variables

import "fmt"

// Asignación de variables

// Se usa la var keyword
// var variablename type = value, Se puede declarar dentro o fuera de una función

var myString string = "Hola Amigos"

// Usando asignación corta
// variablename := value, aunque existen reglas para usar esto y esque se deben de declarar siempre dentro de una función
// myNumber := 100

// Asignación de multiple variables
var a, b, c, d int = 1, 3, 5, 7

var greetHello, amount = myString, 100.0 // Aunque hay que tener en cuenta de que si se le añade los tipos no se puede aplicar

func PrintVariables() {
	fmt.Println("Usando paquete de variables")
	fmt.Println(myString)
	myNumber := 100
	fmt.Println(myNumber)
	fmt.Println(a, b, c, d)
	fmt.Println(greetHello)
	fmt.Println(amount)
}
