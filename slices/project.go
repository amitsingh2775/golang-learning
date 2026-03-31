package main

import (
	"fmt"
)

// what we will do
//   Hume multiple users store karne hain
//  dynamic size chahiye
// operations perform karne hain

// pahle sare users ke pass id and name hoga so i will create a cunstruct

type User struct {
	ID   int
	Name string
}

func addUser(users []User, ID int, name string) []User {
	data := User{ID, name}
	users = append(users, data)
	return users

}

// fetch list of users

func fetchList(users []User) {
	for _, u := range users {
		fmt.Println(u.ID, u.Name)
	}
}

// delete user

func deleteUser(users[]User,id int)[]User{
	// method 1.
	var updatedUsers[] User
	for _,u:=range users {
		if u.ID!=id {
			updatedUsers=append(updatedUsers,u)
		}
	}
	return updatedUsers
}

// search user
func serachUser(users[]User,name string)string{
	
	for _,u:=range users{
		if u.Name==name{
			return u.Name
		}
	}
	return "not found"
}

func main() {
	var users []User
	// reader:=bufio.NewReader(os.Stdin)
	//1. ham users ko add karegne
	for i := 0; i <= 2; i++ {
		var id int
		fmt.Scan(&id)
		var name string
		fmt.Scan(&name)

		users = addUser(users, id, name)
	}
	fmt.Println("----------")
	fetchList(users)
	fmt.Println(".................")

	fmt.Println(serachUser(users,"amit"))
	fmt.Println("..................")
	fmt.Print(deleteUser(users,2))
	// fmt.Println(users)
}
