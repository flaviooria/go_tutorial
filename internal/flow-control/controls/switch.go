package controls

import (
	"fmt"
	"go_tutorials/internal/utils"
)

func getRandomCity() string {
	myCities := [...]string{"Lima", "Buenos Aires", "Jaén", "Madrid"}

	return myCities[utils.GetNumberRandomWithLimit(len(myCities))]
}

func basicSwitch() {
	city := getRandomCity()

	switch city {
	case "Lima":
		fmt.Println("la ciudad es", city)
	case "Buenos Aires":
		fmt.Println("la ciudad es", city)
	case "Jaén":
		fmt.Println("la ciudad es", city)
	case "Madrid":
		fmt.Println("la ciudad es", city)
	default:
		fmt.Println("La ciudad no existe")
	}
}

func invokeFuncOnSwitch() {

	// Podemos llamar a una función directamente dentro del switch

	switch getRandomCity() {
	case "Lima":
		fmt.Println("la ciudad es Lima")
	case "Buenos Aires":
		fmt.Println("la ciudad es Buenos Aires")
	case "Jaén":
		fmt.Println("la ciudad es Jaén")
	case "Madrid":
		fmt.Println("la ciudad es Madrid")
	default:
		fmt.Println("La ciudad no existe")
	}
}

func nextStepLogicalSwitch() {

	// Usamos la keyworkd fallthrough ya que a diferencia de otro lenguajes donde se tiene un brake en Go eso es implicto
	// por lo tanto usamos el falltrough para que siga la lógica en el siguiente caso
	city := getRandomCity()

	switch city {
	case "Lima":
		fallthrough
	case "Buenos Aires":
		fmt.Println("Lima y Buenos Aires son ciudades de Latino América")
	case "Jaén":
		fallthrough
	case "Madrid":
		fmt.Println("Jaén y Madrid son ciudades de Europa")
	default:
		fmt.Println("La ciudad no existe")
	}
}

func PrintSwitchs() {
	basicSwitch()
	invokeFuncOnSwitch()
	nextStepLogicalSwitch()
}
