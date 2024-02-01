package controls

import "fmt"

func loopBasicFor() {
	for num := 1; num <= 100; num++ {
		fmt.Println("El número es: ", num)
	}
	fmt.Println()
}

func loopWithContinueFor() {
	sum := 0
	for num := 1; num <= 100; num++ {
		if num%5 == 0 {
			continue
		}
		sum += num
	}
	fmt.Println("The sum of 1 to 100, but excluding numbers divisible by 5, is", sum)
	fmt.Println()

}

func loopForSimilarALoopWhile() {
	num := 3
	for num != 5 {
		fmt.Printf("El número %d es menor que 5\n", num)
		num += 1
	}
	fmt.Println()
}

func FizzBuzz() {
	fizz := "fizz"
	buzz := "buzz"

	for num := 1; num <= 100; num++ {
		switch {
		case num%3 == 0:
			fmt.Printf("%s \n", fizz)
		case num%5 == 0:
			fmt.Printf("%s \n", buzz)
		default:
			fmt.Printf("%s%s \n", fizz, buzz)
		}
	}
}

func findprimes(number int) bool {
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}

	if number > 1 {
		return true
	} else {
		return false
	}
}

func PrintLoops() {
	loopBasicFor()
	loopForSimilarALoopWhile()
	loopWithContinueFor()
	FizzBuzz()
}
