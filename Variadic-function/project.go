package main

import "fmt"

type User struct{
   name string
}

func addUser(users[]User,newUser ...User)[]User{
	users=append(users,newUser...)
	return users
}

func main(){

	users:=[]User{}

	users=addUser(users,
		User{"amit"},
		User{"rakhes"})
	fmt.Println(users)

	
}