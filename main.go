package main

import "fmt"

func main() {
	user := User{}
	user.ID = 1
	user.FirstName = "Yudi"
	user.LastName = "Setiawan"
	user.Email = "kolonel.yudisetiawan@gmail.com"
	user.IsActive = true

	user2 := User{
		ID:        2,
		FirstName: "Link",
		LastName:  "Awakening",
		Email:     "link@nintendo.com",
		IsActive:  false,
	}

	user3 := User{3, "Eka", "Pablo", "eka@gmail.com", false}

	fmt.Println(user)
	fmt.Println(user.FirstName)
	fmt.Println(user2)
	fmt.Println(user2.FirstName)
	fmt.Println(user3)
	fmt.Println(user3.FirstName)
}

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}
