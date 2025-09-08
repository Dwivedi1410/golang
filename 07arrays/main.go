package main

import "fmt"

func main() {
	fmt.Printf("Hello this is a array section")

	var arr = [3]int{1,2,3};   //fixed size array declaration
	fmt.Println(arr);  //[1,2,3]

	var arr2 = [...]int{3,4,5,2}  //variable length array declaration
	fmt.Println(arr2);  //[3,4,5,2]

	arr3 := [2]string{"hello", "world"}
	fmt.Println(arr3)   //[hello world]


	//you can also initialize the specific element in an array

	arr4 := [5]int{0:1, 3:2}
	fmt.Println(arr4)  //[1,0,0,2,0]

	//you can use len() function to find the length of the array
	fmt.Println(len(arr4));  //5
}