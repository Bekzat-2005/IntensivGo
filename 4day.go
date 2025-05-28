package main

import "fmt"

type Student struct {
	name  string
	grade int
}

func (s Student) studentInfo() {
	fmt.Println("name: ", s.name, " grade: ", s.grade)
}

func (s Student) exam() bool {
	return s.grade >= 70
}

func (s Student) Level() string {
	switch {
	case s.grade >= 90:
		return "A"
	case s.grade >= 70:
		return "B"
	case s.grade >= 50:
		return "C"
	default:
		return "F"
	}
}

func main() {

	students := make(map[string]Student)

	students["Ali"] = Student{name: "Ali", grade: 75}
	students["Miras"] = Student{name: "Miras", grade: 69}

	for _, student := range students {
		student.studentInfo()
		fmt.Print(student.Level())
		if student.exam() {
			fmt.Println("Ok")
		} else {
			fmt.Println("No Leta")
		}
	}

}
