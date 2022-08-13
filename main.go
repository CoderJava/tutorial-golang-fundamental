package main

import "fmt"

func main() {
	student := Student{
		ID:   1,
		Name: "Agung Setiawan",
		GPA:  3.7,
	}
	var student2 *Student = &Student{
		ID:   2,
		Name: "Eko Kurniawan",
		GPA:  3.8,
	}
	fmt.Println(student.Name)
	student.graduate()
	fmt.Println(student.Name)
	fmt.Println("====================")
	fmt.Println(student2.Name)
	student2.graduate()
	fmt.Println(student2.Name)
}

type Student struct {
	ID   int
	Name string
	GPA  float32
}

func (student *Student) graduate() {
	student.Name = student.Name + " S.T"
}
