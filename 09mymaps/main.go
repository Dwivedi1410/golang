package main

import "fmt"

func main() {
	fmt.Println("This is a map section")
 
	//This is how we declare maps
	courses := make(map[string]string)
	courses["js"] = "javascript"
	courses["c++"] = "cplusplus"
	courses["py"] = "python"

	fmt.Println(courses)


	//How we can delete any value from the map

	delete(courses, "py");
	fmt.Println(courses)


	//How we can access the value of any key
	fmt.Println(courses["c++"])  //cplusplus


	// loop for maps

	for key, value := range courses{
		fmt.Printf("%v is the short form of %v \n", key, value)
	}

	//if you don't want to use the value then you can use '_'

	for _, value := range courses{
		fmt.Printf("these are the values %v\n", value);
	}

}