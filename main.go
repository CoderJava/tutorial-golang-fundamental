package main

import "fmt"

func main() {
	student := Student{
		ID:   1,
		Name: "Agung Setiawan",
		GPA:  0,
	}
	fmt.Println(student.Name)
	graduate(&student)
	fmt.Println(student.Name)
}

type Student struct {
	ID   int
	Name string
	GPA  float32
}

func graduate(student *Student) {
	student.Name = student.Name + " S.T"
}
