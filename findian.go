package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("This program checks the chracter i, n and a, if the entered string has first letter i and last letter n and contains a, this \n will print" +
		"Found, other Not Found")

	// Noraml strings scan only read one string, which means if the user enter "this is the test", then it won't read entire string
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter your string : ")
	// read the user input
	zUserInput, _ := reader.ReadString('\n')
	// this user input contains the new line as well, we should remove it before processing the input
	zModifiedInput := strings.TrimSuffix(zUserInput,"\n")
	// In order to apply the logic, we will convert move all the characters into uppercase, so that we can easily validate the logic
	zModifiedUpperCaseInput :=strings.ToUpper(zModifiedInput)
	// just check the length, if it is less than 3, then we can clearly say, not found
	b := len(zModifiedUpperCaseInput) < 3
	if (b){
		fmt.Printf("Not Found\n")
	} else{
		// extract the first character using slicing
		zFirstCharacter :=zModifiedUpperCaseInput[0:1];
		// get the last character using slicing
		zLastCharacter :=zModifiedUpperCaseInput[len(zModifiedUpperCaseInput)-1:];
		// now check whether both first and last characters are matching with given condition
		if(strings.Compare(zFirstCharacter,"I") ==0 &&strings.Compare(zLastCharacter,"N") ==0 ){
		     //now check whether the string has character a
			if(strings.Contains(zModifiedUpperCaseInput,"A")){
				 fmt.Printf("Found\n")
			 }else{
				 fmt.Printf("Not Found\n")
			 }

		}else{
			fmt.Printf("Not Found\n")
		}

	}

}