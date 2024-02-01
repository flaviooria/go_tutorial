package controls

import (
	"fmt"
	"go_tutorials/internal/utils"
	"math/rand"
)

// Vamos a tratar los condicionales,
//algo que recalcar es que a diferencias de otros lenguajes no tiene ternario

func PrintConditionals() {
	numberTime := 12
	random := rand.Intn(100-1) + 1

	if numberTime >= random {
		fmt.Println("Es mayor", numberTime)
	} else {
		fmt.Println("Es menor que", numberTime)
	}

	// Con Go podemos obtener directamente una variable y trabajar con ella dentro del bloque
	//esto nos da mucha más ventaja para manipular datos si tener que declaralos antes de realizar una condición

	if numberRandom := utils.GenerateNumberRandom(); numberRandom >= random {
		fmt.Println("Es mayor", numberTime)
	} else {
		fmt.Println("Es menor que", numberTime)
	}
}
