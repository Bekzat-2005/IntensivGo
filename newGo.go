package main

// task 2
// func square(n *int) {
// 	*n = *n * *n
// }

// task 3
// func increment(nums *[]int) {
// 	for i := range *nums {
// 		(*nums)[i] += 1
// 	}
// }

// task 4
// type Person struct {
// 	name string
// 	age  int
// }

// func birthday(p *Person) {
// 	p.age += 1
// }

func main() {

	// task 6

	// numbers := [5]int{10, 20, 30, 40, 50}
	// sum := 0

	// for _, v := range numbers {
	// 	sum += v
	// }

	// fmt.Print(sum)

	// task 5

	// students := map[string]int{
	// 	"Ali":   17,
	// 	"Aslan": 15,
	// }

	// students["Ramazan"] = 10
	// delete(students, "Ali")
	// fmt.Print(students)

	// task 4

	// p := Person{"Bekzat", 20}
	// birthday(&p)
	// fmt.Print(p)

	// task 3

	// numbers := []int{1, 2, 3}
	// increment(&numbers)
	// fmt.Print(numbers)

	// task 2

	// x := 5
	// square(&x)
	// fmt.Println(x)

	//  task 1
	// x := 10
	// p := &x // p — указатель на x

	// fmt.Println("x =", x)
	// fmt.Println("p =", p)   // адрес
	// fmt.Println("*p =", *p) // разыменование — значение по адресу

	// *p = 20 // изменяем x через указатель
	// fmt.Println("x после изменения:", x)

}
