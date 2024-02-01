package basictypes

import "fmt"

// Datatypes in Go

// Strings

var myString string = "Mi string"

// Aqui no añadimos el tipo ya que Go lo puede inferir
var myNumber = 1

// float
var myFloat float32 = 3.56

// boolean
var myBool bool = false

func PrintTypes() {
	fmt.Println("Usando paquete de basictypes")
	fmt.Println("Un string: ", myString)
	fmt.Println("Un número: ", myNumber)
	fmt.Println("Un float: ", myFloat)
	fmt.Println("Un booleani", myBool)
}
