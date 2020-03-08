package main

import (
	"fmt"
)


func main(){
	fmt.Printf("This program is check the floating point truncation ....\n")
	var fUserInput float64
	fmt.Printf("Please enter the float number:")
	fmt.Scan(&fUserInput)
	var iUserInput int64
	iUserInput = int64(fUserInput)
	fmt.Printf("Truncated value [%d] \n",iUserInput)
}