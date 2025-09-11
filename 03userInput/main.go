package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	msg := "Hello welcome to my Pizza shop"
	fmt.Println(msg, "\n")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter the rating of our pizza :")

	//Note:- golang dosn't have a try catch to catch the errors
	//golang treat errors as a variables like true and false.
	//golang use comma ok or comma error syntax

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanx for the rating, ", input);
	fmt.Printf("Type of input is %T", input);    //string
}
