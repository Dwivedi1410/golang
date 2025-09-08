package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to my Pizza app")
	fmt.Println("Please enter the ratig of pizza between 1 to 5")
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')

	if err != nil {
		println("Something went wrong")
	} else {
		// newRating, _ := strconv.Atoi(input) //this will give me error
		newRating, _ := strconv.Atoi(strings.TrimSpace(input))  //Here I am converting string to integer
		println("Thanx for the rating of ", newRating+1)  
	}
}
