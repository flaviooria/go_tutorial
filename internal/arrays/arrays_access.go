package arrays

import "fmt"

func AccessPosition() {
	months := [...]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	segment := months[3:6]

	fmt.Println(segment)
	fmt.Println("La longitud es: ", len(segment))
	fmt.Println("La capacidad es: ", cap(segment))

	extendSegment := segment[:4]
	fmt.Println(extendSegment)

	// Como podemos ver en Go no es necesario indicarle la primera posición el entiende que se quieren los datos desde la última posición pasada
	// hacia de delante
	// segment[:5] => ["April", "May", "June", "July", "August"]

	// Y lo mismo sucederia si declaramos la primera posición, lo que hara es darte los datos desde esa posición hasta la última
	// segment[8:] => ["August", "September", "October", "November", "December"]

}

func AddNewElements() {
	numbers := []int{}

	for i := 0; i <= 10; i++ {
		numbers = append(numbers, i)
		fmt.Println("Número es: ", cap(numbers), numbers)
	}
}

func RemoveItemFromArray() {
	letters := []string{"A", "B", "C", "D", "E"}

	posRemove := 2

	if posRemove < len(letters) {
		letters = append(letters[:posRemove], letters[posRemove+1:]...)
	}

	fmt.Println(letters)
}
