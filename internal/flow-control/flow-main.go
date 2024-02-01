package flow_control

import (
	"fmt"
	"go_tutorials/internal/flow-control/controls"
)

func PrintFlowControl() {
	fmt.Println("Mostrando condicionales")
	controls.PrintConditionals()
	fmt.Println()
	fmt.Println("Mostrando switchs")
	controls.PrintSwitchs()
	fmt.Println()
	fmt.Println("Mostrando bucles for")
	controls.PrintLoops()
}
