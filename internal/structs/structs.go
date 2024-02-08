package structs

import (
	"encoding/json"
	"fmt"
)

// Los structs son un tipo de datos en la cual viene representada por una colección de campos

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

// Usando nuestro struct

var flavio Employee

func PrintStructs() {
	flavio = Employee{ID: 1, FirstName: "Flavio", LastName: "Oria", Address: "Juanito Perez 10"}

	// Para acceder a una propiedad de un struct se hace mediante el uso del punto
	println(flavio.FirstName)

}

// Podriamos también insertar datos a otro struct sin tener que declararlos de esta manera, de esta manera podrianos refactorizar código

type InformationEmployee struct {
	AccountBankNumber int
	Nif               string
	Employee
}

func PrintStructsInsert() {

	flavio := InformationEmployee{AccountBankNumber: 43543543, Nif: "F12345", Employee: Employee{ID: 1, FirstName: "Flavio"}}
	flavio.LastName = "Oria"
	flavio.Address = "Juanito Perez 10"

	fmt.Println(flavio)
}

// Podemos tambien crear una copia de un struct creando un puntero con el operador "&" y así modificariamos sus valores

func PrintStructPointer() {
	flavio = Employee{ID: 1, FirstName: "Flavio", LastName: "Oria", Address: "Juanito Perez 10"}

	fmt.Println("Soy flavio sin apuntar aún", flavio.FirstName)
	flavioClon := &flavio

	flavioClon.FirstName = "Flavio Clone"

	fmt.Println("Soy flavio ya cambiado con puntero", flavio.FirstName)

}

// Codificando y decodificando un struct
type Employee2 struct {
	ID        int
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname,omitempty"`
	Address   string
}

func PrintStrucEncodedAndDecoded() {

	employee := Employee2{ID: 1, FirstName: "Raul", LastName: "Navas", Address: "Amauta 7"}

	encoded, error := json.Marshal(employee)

	if error != nil {
		return
	}

	var decoded Employee
	error_decoded := json.Unmarshal(encoded, &decoded)

	if error_decoded != nil {
		return
	}

	fmt.Printf("%s\n", encoded)
	fmt.Println()
	fmt.Printf("%v", decoded)
}
