package main

import "fmt"

func main() {
	//значение по умолчанию
	var num0 int

	fmt.Println(num0) //ноль по умолчанию

	var num1 int = 1

	var num2 = 20
	fmt.Println(num0, num1, num2)

	//короткое объявление переменной
	//только для новых переменных
	num := 30
	num += 1
	fmt.Println("+=", num)

	num += 2

	fmt.Printf("+= %d", num)

	//++num нету
	num++
	fmt.Println("++", num)

	//cameCase - принятый стиль
	userIndex := 10
	//under_score - не принято
	user_index := 10
	fmt.Println(userIndex, user_index)

	//объявление нескольких переменных
	var weight, height int = 1, 1
	fmt.Printf("weight %d, height %d", weight, height)
	fmt.Println()
	//присваивание в существующеи переменные
	weight, height = 11, 21
	fmt.Printf("weight %d, height %d", weight, height)
	fmt.Println()

	weight, height = 33, 34
	fmt.Printf("weight %d, height %d", weight, height)
	fmt.Println()

	//короткое присваивание
	//хотя-бы одна переменная даолжна быть новой!
	//weight, height := 12, 22 есть ошибка
	weight, age := 12, 22

	fmt.Println(weight, height, age)
	fmt.Println()

	var test string
	fmt.Println(test)
}
