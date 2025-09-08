//why we need pointers:- sometimes when we pass some arrays and veriables to some functions they are not get pass propely (copy of verible is passed not the actual veriable)that's why we use pointers.

package main

import "fmt"

func main(){
	fmt.Println("Hello this is pointer section")

	var ptr *int
	fmt.Println("This is the value of my poiner ", ptr)  //nil

	val := 23
	ptr = &val
	fmt.Println("This is the value of pointer after the assignment of the address ", ptr)  //0xc00000a088 
	fmt.Println("This is the value of pointer after derefrencing ", *ptr);  //23

	val2 := 12
	ptr2 := &val2
	fmt.Println("This is the value of the second pointer ", ptr2)  //0xc00008c078
	fmt.Println("This is the value of the second pointer after derefrencing ", *ptr2)  //12

	*ptr2 = *ptr2 + 2;
	fmt.Println("This is the new value of val2 ", val2); //14
}