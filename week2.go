package main

import "fmt"

//  task 2
// func square(n *int) {
// 	*n = *n * 10
// }
// task 3
// func increment(nums *[]int) {
// 	// for i, n := range *nums {
// 	// 	(*nums)[i] = n * 2
// 	// }
// 	for i, n := range *nums {
// 		(*nums)[i] = n + 1
// 	}

// }

// day 9

type Person struct {
	name string
	age  int
}

func bitthday(p *Person) {
	p.age += 1
}

func main() {

	// day 9

	p := Person{name: "Bekzat", age: 19}
	fmt.Print(p.name)

	bitthday(&p)
	fmt.Println(p.age)

	var u Person
	u.name = "ALI"

	fmt.Print(u.name)
	// day 8

	// task 3

	// numbers := []int{1, 2, 3, 4, 5}
	// increment(&numbers)
	// fmt.Print(numbers)

	// task 2

	// x := 5
	// square(&x)
	// fmt.Println(x)

	// task 1
	// x := 10
	// p := &x

	// *p = 20

	// fmt.Println(*p)
	// fmt.Println(p)

	// fmt.Println(x)
}
