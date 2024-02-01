package main

import (
	"fmt"
	"go_tutorials/internal/arrays"
	"go_tutorials/internal/basictypes"
	flow_control "go_tutorials/internal/flow-control"
	"go_tutorials/internal/variables"
)

func main() {

	basictypes.PrintTypes()
	fmt.Println()
	variables.PrintVariables()
	fmt.Println()
	flow_control.PrintFlowControl()
	fmt.Println()
	arrays.PrintArray()
}
