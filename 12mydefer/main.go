package main

import "fmt"

func main() {
	//defer:- if you put defer before anyline then it will get placed at the end just before the closing of the scope
	// Note:- if their are too many defers then they will follow the last in first out(LIFO)

	// defer fmt.Println("Hello")
	// fmt.Println("World")

	/*
	output :- 	World
				Hello
	*/

	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")
	fmt.Println("Hello")
	count()


	/*
	output :- 	4
				3
				2
				1
				0
				three
				two
				one
	*/


}

func count(){
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}