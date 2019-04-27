package main

import (
	"fmt"
	"os/user"
)

func main() {
	curUser, _ := user.Current() //igonre error
	fmt.Println("Current User Id", curUser.Uid)
	fmt.Println("Current User Name", curUser.Username)
	fmt.Println("Current Grouo Id", curUser.Gid)

}
