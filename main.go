package main

import "fmt"

func main() {

	//day 7
	students := map[string]int{
		"Ali":    85,
		"Bekzat": 90,
		"Dana":   78,
	}

	// fmt.Println("Изначально:", students)

	students["Dana"] = 88   // обновляем оценку
	delete(students, "Ali") // удаляем Али

	// fmt.Println("После изменений:", students)

	score := students["Aru"]
	fmt.Println(score) // 👉 Вернёт 0

	//  4 day
	// task1
	// numbers := [5]int{10, 20, 30, 40, 50}
	// sum := 0

	// for _, value := range numbers {
	// 	sum += value
	// }

	// fmt.Println("Сумма элементов массива:", sum)

	// task2

	// languages := []string{"Go", "Python", "Java"}
	// languages = append(languages, "Kotlin")

	// fmt.Println("Языки программирования:", languages)

	// // task3
	// for i, lang := range languages {
	// 	fmt.Println("Индекс:", i, "Язык:", lang)
	// }

	// 3day

	// printInfo("Бекзат", 19)

	// result := multiply(4, 5)
	// fmt.Println("Результат умножения:", result)

	// fmt.Println("Число 4 чётное?", isEven(4)) // true
	// fmt.Println("Число 7 чётное?", isEven(7)) // false

	// task2

	// language := "go"

	// switch language {
	// case "go":
	// 	fmt.Print("Go — супер!")
	// case "python":
	// 	fmt.Print("Python — гибкий")
	// case "java":
	// 	fmt.Print("Java — надёжный")
	// default:
	// 	fmt.Println("")
	// }

	// task1
	// var temp int
	// fmt.Scan(&temp)

	// if temp > 30 {
	// 	fmt.Print("Очень жарко")
	// }
	// else if temp >= 15 {
	// 	fmt.Print("Тепло")
	// } else {
	// 	fmt.Print("Холодно")
	// }
}

func printInfo(name string, age int) {
	fmt.Println("Имя:", name, "Возраст:", age)
}
func isEven(n int) bool {
	return n%2 == 0
}
func multiply(a int, b int) int {
	return a * b
}
