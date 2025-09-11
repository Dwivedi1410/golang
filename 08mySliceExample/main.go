package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hello this is a slice section")

	//Note :- Their are three ways by which we can delcare a slice

	//1) by using []datatype{}
	var slice1 = []int{1, 2, 3}
	fmt.Println(slice1) //[1,2,3]

	//2) by using array
	var arr1 = [4]int{1, 2, 3, 4}
	slice2 := arr1[1:3]
	fmt.Println(slice2) //[2,3]

	//3) By using make([]type, length, capacity) if capacity is not provided then it is equivalent to the length provided
	slice3 := make([]int, 3)
	slice3 = append(slice3, 1, 2, 3) //[0 0 0 1 2 3]
	slice3[0] = 5
	slice3[1] = 6
	slice3[2] = 7
	fmt.Println(slice3) //[5 6 7 1 2 3]

	// let's check the type of this slice
	fmt.Printf("type of this slice is %T", slice3) //type of this slice is []int

	fruitList := []string{"Mango", "Banana", "Peach", "Papaya"}
	// fruitList = append(fruitList[:3])  //[Mango Banana Peach]
	// fmt.Println(fruitList);

	sort.Strings(fruitList)
	fmt.Println("This is a sorted slice ", fruitList) //[Banana Mango Papaya Peach]

	// How we can remove a value from the slice based on the index

	var courses = []string{"javascript", "react.js", "ruby", "python", "swift", "c++"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:2], courses[index+1:]...)
	fmt.Println(courses)
}
