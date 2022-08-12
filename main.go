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

	users := []User{
		user,
		user2,
		user3,
	}
	group := Group{
		Name:        "Gamer",
		Admin:       user,
		Users:       users,
		IsAvailable: true,
	}
	displayGroup(group)
}

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

type Group struct {
	Name        string
	Admin       User
	Users       []User
	IsAvailable bool
}

// func displayUser(user User) string {
// 	output := fmt.Sprintf("Name: %s %s, Email: %s", user.FirstName, user.LastName, user.Email)
// 	return output
// }

func displayGroup(group Group) {
	fmt.Printf("Name: %s", group.Name)
	fmt.Println()
	fmt.Printf("Member count: %d", len(group.Users))
	fmt.Println()

	for _, user := range group.Users {
		fmt.Println(user.FirstName)
	}
}
