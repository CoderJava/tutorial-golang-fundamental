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

	resultUser1 := displayUser(user)
	resultUser2 := displayUser(user2)
	resultUser3 := displayUser(user3)
	fmt.Println(resultUser1)
	fmt.Println(resultUser2)
	fmt.Println(resultUser3)
}

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

func displayUser(user User) string {
	result := fmt.Sprintf("Name: %s %s, Email: %s", user.FirstName, user.LastName, user.Email)
	return result
}