package main

import "fmt"

func main() {
	//инициализация при создании

	var user map[string]string = map[string]string{
		"name":     "Vadim",
		"lastName": "Nemov",
	}

	//сразу с нужной емкостью
	profile := make(map[string]string, 10)

	//количество элементов
	mapLength := len(user)
	fmt.Printf("%d %+v\n", mapLength, profile)

	// если ключа нет - вернет значение по умолчанию для типа
	mName := user["middleName"]
	fmt.Println("mName:", mName)

	// проверка на существование ключа
	mName, mNameExist := user["middleName"]
	fmt.Println("mName:", mName, "mNameExist:", mNameExist)

	// пустая переменная - только проверяем что ключ есть
	_, mNameExist2 := user["middleName"]
	fmt.Println("mNameExist2", mNameExist2)

	// удаление ключа
	delete(user, "lastName") //для слайса нету метода удаления, только для map?
	fmt.Printf("%#v\n", user)

	//----------------------------------
	var testMap = map[string]string{"Key": "Value"}
	fmt.Println(testMap["Key"])

	delete(testMap, "Key")
	fmt.Println(testMap["Key"])
}
