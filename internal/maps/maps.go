package maps

import "fmt"

// Vamos hacer uso de los maps que son diccionarios o hash

// Inicialización de un map

func PrintMaps() {
	// Para crear un map se usa la keyword map y luego se le indica entre corchetes el tipo de la clave y luego el tipo del valor
	myMap := map[string]int{}

	// Otra manera de inicializar un map
	myMap2 := make(map[string]int)

	// Asignanciones en un map
	myMap["Flavio"] = 25
	myMap["Lau"] = 25

	myMap2["Jepeto"] = 65
	myMap2["Pinocho"] = 10

	fmt.Println("Mostrando los maps y sus respectivos datos: ")
	fmt.Println(myMap)
	fmt.Println(myMap2)

	fmt.Println("iterando sobre los datos de un map: ")

	for k, v := range myMap {
		fmt.Println("La clave es: ", k)
		fmt.Println("El valor es: ", v)
	}

	fmt.Println("Eliminando un dato del map")

	delete(myMap2, "Pinocho")
	fmt.Println(myMap2)

}
