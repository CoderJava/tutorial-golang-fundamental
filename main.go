package main

import (
	"fmt"
	"golang-fundamental/management"
)

func main() {
	user := management.User{}
	user.ID = 1
	user.FirstName = "Yudi"
	user.LastName = "Setiawan"
	user.Email = "kolonel.yudisetiawan@gmail.com"
	user.IsActive = true
	user2 := management.User{
		ID:        2,
		FirstName: "Link",
		LastName:  "Awakening",
		Email:     "link@nintendo.com",
		IsActive:  false,
	}
	user3 := management.User{
		ID:        3,
		FirstName: "Eka",
		LastName:  "Pablo",
		Email:     "eka@gmail.com",
		IsActive:  false,
	}
	fmt.Println(user.Display())
	fmt.Println(user2.Display())
	fmt.Println(user3.Display())

	users := []management.User{
		user,
		user2,
		user3,
	}
	group := management.Group{
		Name:        "Gamer",
		Admin:       user,
		Users:       users,
		IsAvailable: true,
	}
	group.Display()
}
