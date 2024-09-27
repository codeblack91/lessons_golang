package main

import "fmt"

func main() {

	//простое условие
	boolval := true
	if boolval {
		fmt.Println("boolval is true")
	}

	mapVal := map[string]string{"firstName": "Vasily"}
	//условие с блоком инициализации
	if keyValue, keyExist := mapVal["name"]; keyExist {
		fmt.Println("name =", keyValue)
	}
	/*

		//получаем только признак существа ключа
		if _, keyExist := mapVal["name"]; keyExist {
			fmt.Println("key 'name' exist")
		}

		cond := 1
		//множественные if else
		if cond == 1 {
			fmt.Println("cond is 1")
		} else if cond == 2 {
			fmt.Println("cond is 2")
		}

		s := "ku-ku-ku"
		if length := len(s); length > 3 {
			fmt.Println("Длина строки больше 3:", length)
		}

		// switch по 1 переменной
		strVal := "name"
		switch strVal {
		case "name":
			fallthrough //обязательно выполняет то что после него
		case "test", "lastName":
			// some work
		default:
			// some work
		}

		// switch как замена многим ifelse
		var val1, val2 = 2, 2
		switch {
		case val1 > 1 || val2 < 11:
			fmt.Println("first block")
		case val2 > 10:
			fmt.Println("second block")
		}
	*/
Loop: //что за оператор??????? какой-то старый синтаксис?
	for key, val := range mapVal {
		println("switch in loop", key, val)
		switch {
		case key == "lastName":
			break
			fmt.Println("dont pront this")
		case key == "firstName" && val == "Vasily":
			fmt.Println("switch - break loop here")
			break Loop
		}
	} // конец for
}
