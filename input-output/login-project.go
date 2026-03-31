package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	adminUserName := "admin"
	adminPassword := "admin123"

	attempts := 3
	check:=false
	for attempts > 0 {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("enter your username:")
		username, _ := reader.ReadString('\n')
		username = strings.TrimSpace(username)
		fmt.Println("enter your password:")
		password, _ := reader.ReadString('\n')
		password = strings.TrimSpace(password)

		if username == adminUserName && password == adminPassword {
			check=true
			fmt.Println("login successful")
			break
			
		} else {
			fmt.Println("invaild credentials, try again", "attempts left:", attempts-1)
			attempts--
		}

	}
	if !check{
	fmt.Println("login failed")
	}

}
