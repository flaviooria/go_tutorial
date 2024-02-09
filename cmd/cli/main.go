package main

import (
	"fmt"
	"go_tutorials/internal/arrays"
	"go_tutorials/internal/basictypes"
	"go_tutorials/internal/exercises"
	flow_control "go_tutorials/internal/flow-control"
	"go_tutorials/internal/maps"
	"go_tutorials/internal/structs"
	"go_tutorials/internal/variables"
)

func main() {

	basictypes.PrintTypes()
	fmt.Println()
	variables.PrintVariables()
	fmt.Println()
	flow_control.PrintFlowControl()
	fmt.Println()
	arrays.PrintInitializationArrays()
	fmt.Println()
	arrays.AccessPosition()
	fmt.Println()
	arrays.AddNewElements()
	fmt.Println()
	arrays.RemoveItemFromArray()
	fmt.Println()
	maps.PrintMaps()
	fmt.Println()
	structs.PrintStructs()
	fmt.Println()
	structs.PrintStructsInsert()
	fmt.Println()
	structs.PrintStructPointer()
	fmt.Println()
	structs.PrintStrucEncodedAndDecoded()
	fmt.Println()
	exercises.PrintFibonnaci()
}
