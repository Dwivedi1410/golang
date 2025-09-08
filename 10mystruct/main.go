package main

import "fmt"

func main() {
	fmt.Println("Hello this is a struct section")

	//Note:- Their is no inheritence in golang , no super and their is no concept of parent and child in the golang.

	Prakash := User{"Prakash", "pd123@dev.com", false, 14}
	fmt.Println(Prakash)

	fmt.Printf("Thease are the values of the struct in more detail %+v", Prakash);

	fmt.Printf("Name of the struct is %v and age of the struct is %v", Prakash.Name, Prakash.Age)


}

//This is how we define a struct
type User struct{
	Name string
	Email string
	Verified bool
	Age int
}