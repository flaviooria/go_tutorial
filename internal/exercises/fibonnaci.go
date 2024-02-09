package exercises

import (
	"fmt"
)

func fibonnaci(num int) {

	// Esta es mi forma jajaja igual sale pero haciendo un poquillo de trampa

	prev := 1
	current := 0
	acum := 0
	listNumbers := []int{}

	if num < 2 {
		listNumbers = append(listNumbers, []int{0, 1}...)
	} else {
		for i := 1; i < num-2; i++ {
			if prev <= 1 {
				prev += i
			} else {
				prev += current
			}

			acum = prev
			current = i

			listNumbers = append(listNumbers, acum)
		}

		listNumbers = append([]int{0, 1, 1}, listNumbers...)
	}

	fmt.Println(listNumbers)
}

func fibonacci2(num int) {

	// Otra forma esta si es sin hacer trampilla

	prev := 0
	current := 1
	nextNum := 0
	listNumbers := []int{}

	for i := 1; i <= num; i++ {
		listNumbers = append(listNumbers, prev)
		nextNum = prev + current
		prev = current
		current = nextNum
	}

	fmt.Println(listNumbers)

}

func PrintFibonnaci() {
	fibonnaci(7)
	fibonacci2(7)
}
