package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

/**/
func main() {
	fmt.Println("Guess the no in Dice. Enter q if want to Quit.")

	var input string
	result := rand.Intn(6)

	// loop with do while condition
	for {
		fmt.Print("Guess the number: ")
		fmt.Scan(&input)

		if strings.EqualFold(input, "q") {
			break
		}

		num, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Invalid Input. Please enter a valid no")
			continue

		}

		if num < 0 || num > 6 {
			fmt.Println("Invalid Input. Please enter a valid no. Input should be between 1 and 6")
		}

		if num == result {
			fmt.Println("Congratulations....!!! You gussed the correct value: ", result)

		} else {
			fmt.Println("Bad luck....!!! The correct value: ", result)
		}
	}

}
