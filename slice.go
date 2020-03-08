package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
)
const (
	iConstListSize            int    = 3
	zConstExitCode            string = "X"
)

func main() {
	var intSlice = make([]int, iConstListSize)

	re := regexp.MustCompile(`^\d+$`)
	scanner := bufio.NewScanner(os.Stdin)
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	fmt.Printf("Please enter integer to sort or X to exit: ")
	for scanner.Scan() {
		var strVal string = scanner.Text()
		if strVal == zConstExitCode {
			os.Exit(0)
		}
		match := re.Match([]byte(strVal))
		intVal, err := strconv.Atoi(strVal)
		if match == false || err != nil {
			fmt.Printf("Entered value is not correct, please enter again ")
			continue
		}
		intSlice = append(intSlice, intVal)
		sort.Sort(sort.IntSlice(intSlice))
		var result []int = intSlice[iConstListSize:len(intSlice)]
		fmt.Println(result)
		fmt.Printf("Enter value to sort:")
	}
}

