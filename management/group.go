package management

import "fmt"

type Group struct {
	Name        string
	Admin       User
	Users       []User
	IsAvailable bool
}

func (group Group) Display() {
	fmt.Printf("Name: %s", group.Name)
	fmt.Println()
	fmt.Printf("Member count: %d", len(group.Users))
	fmt.Println()

	for _, user := range group.Users {
		fmt.Println(user.FirstName)
	}
}