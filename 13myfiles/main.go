package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("This is a file section")

	content := "This content will go inside the file"

	file, err := os.Create("./lcogofile.txt") //here I am creating a file

	// if err != nil {
	// 	panic(err) //this is stop the execution of the code and throw error
	// }

	checkNilError(err)

	length, err := io.WriteString(file, content)

	checkNilError(err)

	fmt.Println(length)

	defer file.Close() // defer will close the file at the end

	readFile("./lcogofile.txt")
}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)

	checkNilError(err)

	fmt.Println("Text data inside the file is \n", string(databyte))
}

func checkNilError(err error){
	if err != nil {
		panic(err)
	}
}
