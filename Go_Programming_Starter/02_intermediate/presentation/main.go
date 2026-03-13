package main

import (
	"fmt"
)

func main() {
	// how to remove item from a slice

	mySlice := []int{3, 5, 7}
	newSlice := append(mySlice[:1], mySlice[2:]...)

	fmt.Println(newSlice)

	// Real World Application in API responses
	// type User struct {
	// 	id       int
	// 	name     string
	// 	email    string
	// 	password string
	// }

	// var users []User
	// users = append(users, User{1, "Sundra Bomjan", "bomjan@gmail.com", "nopass"})
	// fmt.Println(users)
}
