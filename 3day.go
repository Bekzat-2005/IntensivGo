package main

import "fmt"

//  task 2
// func sum(nums []int) int {
// 	a := 0
// 	for i := range nums {
// 		a += i
// 	}
// 	return a
// }

// task 4

// func printMap(n map[string]int) {
// 	for name, score := range n {
// 		fmt.Println(name, score)
// 	}
// }

// task 5
// func average(n map[string]int) float64 {
// 	sum := 0
// 	for _, score := range n {
// 		sum += score
// 	}
// 	return float64(sum) / float64(len(n))
// }

// task 7

type School struct {
	students map[string]int
}

func (s School) addStudents(name string, grade int) {
	s.students[name] = grade
}

func main() {

	// task 7
	s := School{students: map[string]int{}}
	s.addStudents("Bekzat", 20)
	fmt.Print(s.students)

	//  task 6

	// languages := map[string][]string{
	// 	"Bekzat": {"Go", "Kotlin"},
	// 	"Ali":    {"Null"},
	// }

	// for name, lang := range languages {
	// 	fmt.Println(name, " Have ", lang)
	// }

	// task 5

	// students := map[string]int{
	// 	"Bekzat": 20,
	// 	"Ali":    17,
	// }
	// avg := average(students)
	// fmt.Print(avg)

	// task 4
	// students := map[string]int{
	// 	"Bekzat": 20,
	// 	"Ali":    17,
	// }
	// printMap(students)

	//  task 3 map

	// students := map[string]int{
	// 	"Bekzat": 20,
	// 	"Ali":    17,
	// }

	// students["Aslan"] = 15
	// delete(students, "Ali")

	// fmt.Print(students)

	//  ------------------------------------
	// task 2
	// nums := []int{1, 2, 3, 4, 5}
	// println(sum(nums))

	// task 1
	// nums := []int{1, 2, 3, 4, 5}
	// for i := range nums {
	// 	nums[i] *= 2
	// }
	// fmt.Print(nums)
}
