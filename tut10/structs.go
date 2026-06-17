package main

import (
	"fmt"
	"time"
)

type lipoo struct {
	string
	int
}
type User struct {
	userid int
	name   string
}
type order struct {
	id        int
	title     string
	desc      string
	createdAt time.Time
	tipoy     struct {
		id   int
		name string
	}
}
type Student struct {
	personalDetails struct { // Anonymous inner structure for personal details
		name       string
		enrollment int
	}
	GPA float64 // Standard field
}

func (o *order) updateTitle(title string) {
	o.title = title
}

func neworder(Id int, Title string, Desc string, CreatedAt time.Time) *order {
	return &order{id: Id, title: Title, desc: Desc, createdAt: CreatedAt}

}

func readptr(o *order) {
	fmt.Println(*o)
}
func main() {
	student := Student{
		personalDetails: struct {
			name       string
			enrollment int
		}{
			name:       "A",
			enrollment: 12345,
		},
		GPA: 3.8,
	}
	// Displaying values
	fmt.Println("Name:", student.personalDetails.name)
	fmt.Println("Enrollment:", student.personalDetails.enrollment)
	fmt.Println("GPA:", student.GPA)

	newlipo := lipoo{"vipu;", 001}
	fmt.Println(newlipo)
	// o := order{
	// 	id:        1,
	// 	title:     "nike shoes",
	// 	desc:      "some prod",
	// 	createdAt: time.Now(),

	// }
	// //fmt.Println(o)
	// readptr(&o)
	// nworder := order{
	// 	id:        01,
	// 	title:     "bag",
	// 	desc:      "some product",
	// 	createdAt: time.Now(),
	// }
	// fmt.Println(nworder)

	// nworder.updateTitle("new prod")
	// fmt.Println(nworder)
	// firstorder := neworder(1, "new bag", "new something", time.Now())
	// fmt.Println(firstorder)
}

// Golang program to illustrate
// the nested structure
// package main

// import "fmt"

// // Creating structure
// type Student struct {
// 	name   string
// 	branch string
// 	year   int
// }

// // Creating nested structure
// type Teacher struct {
// 	name    string
// 	subject string
// 	exp     int
// 	details Student
// }

// func main() {

// 	// Initializing the fields
// 	// of the structure
// 	result := Teacher{
// 		name:    "teacher",
// 		subject: "Java",
// 		exp:     5,
// 		details: Student{"john", "CSE", 2},
// 	}

// 	// Display the values
// 	fmt.Println("Details of the Teacher")
// 	fmt.Println("Teacher's name: ", result.name)
// 	fmt.Println("Subject: ", result.subject)
// 	fmt.Println("Experience: ", result.exp)

// 	fmt.Println("\nDetails of Student")
// 	fmt.Println("Student's name: ", result.details.name)
// 	fmt.Println("Student's branch name: ", result.details.branch)
// 	fmt.Println("Year: ", result.details.year)
// }
