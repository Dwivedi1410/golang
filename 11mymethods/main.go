package main

import "fmt"

func main() {
	fmt.Println("Hello this is a struct section")

	//Note:- Their is no inheritence in golang , no super and their is no concept of parent and child in the golang.

	Prakash := User{"Prakash", "pd123@dev.com", false, 14}
	fmt.Println(Prakash)

	fmt.Printf("Thease are the values of the struct in more detail %+v\n", Prakash);

	fmt.Printf("Name of the struct is %v and age of the struct is %v\n", Prakash.Name, Prakash.Age)

	Prakash.GetStatusOfTheUser();
	Prakash.MailChange()
	fmt.Printf("Mail of the use after calling the MailChange method is %v", Prakash.Email)//this will give the same value as before it was even after calling the changeMail function because a copy of the struct is passed not the address.
}

//This is how we define a struct
type User struct{
	Name string
	Email string
	Verified bool
	Age int
}

func (u User) GetStatusOfTheUser() {
	fmt.Println("Is user verified \n:", u.Verified)
}

func (u User) MailChange() {
	u.Email = "test@mail.co"
	fmt.Printf("New mail of the user is %v\n", u.Email);
}