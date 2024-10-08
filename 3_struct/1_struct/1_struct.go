package main

import "fmt"

type Person struct {
	Id      int
	Name    string
	Address string
}

type Account struct {
	Id int
	// Name    string
	Cleaner func(string) string //функция в структуре???
	Owner   Person
	Person
}

func main() {
	// полное объявление структуры
	var acc Account = Account{
		Id: 1,
		// Name: "vasya",
		Person: Person{
			Name:    "Василий",
			Address: "Москва",
		},
	}
	fmt.Printf("%#v\n", acc)

	// короткое объявление структуры
	acc.Owner = Person{2, "Romanov Vasily", "Moscow"}

	fmt.Printf("%#v\n", acc)

	//fmt.Println(acc.Name)
	//fmt.Println(acc.Person.Name)

	//-----------------------------
	var acc2 Account
	people := Person{
		Id:      2,
		Name:    "Vadim Nemov",
		Address: "Samara",
	}
	acc2 = Account{
		Id:    1,
		Owner: people,
	}

	fmt.Println(acc2)
}
