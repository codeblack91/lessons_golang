package main

import "fmt"

func main() {
	//размер массива является частью его типа
	//инициализация значениями по-умолчанию
	var a1 [3]int
	fmt.Println("a1 short", a1)
	fmt.Printf("a1 short %v\n", a1)
	fmt.Printf("a1 full %#v\n", a1)

	const size = 2
	var a2 [2 * size]bool
	fmt.Println("a2", a2)

	//оперделение размера при объявлении
	a3 := [...]int{1, 2, 3}
	fmt.Println("a2", a3)

	//---------------------------------------
	test := [...]string{"Nemov", "Vadim", "Igorevich"}
	fmt.Printf("Name: %s, Middlename: %s, Surname: %s", test[0], test[1], test[2])

	var test2 [3]int
	fmt.Print(test2)

	const capacity int = 2

	var test3 [capacity * 10]int //нужна именно константа, если int то ошибка
	fmt.Printf("Array exists: %d", test3)
}
