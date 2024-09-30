package main

import "fmt"

type Person struct {
	Id   int
	Name string
}

// не изменит оригинальной структуры, для котрой вызван
func (p Person) UpdateName(name string) {
	p.Name = name
}

// изменяет оригинальную структуру
func (p *Person) SetName(name string) {
	p.Name = name
}

type Account struct {
	Id   int
	Name string
	Person
}

func (p *Account) SetName(name string) {
	p.Name = name
}

type MySlice []int

func (sl *MySlice) Add(val int) {
	*sl = append(*sl, val)
}

func (sl *MySlice) Count() int {
	return len(*sl)
}

type Employee struct {
	Id   int
	Name string
}

// ---------------------------
func (v *Employee) SetNameEmployee(n string) {
	v.Name = n
}

func (v *Employee) SetIdEmployee(i int) {
	v.Id = i
}

func (v *Employee) PrintInfo() {
	fmt.Printf("Id: %d, Name: %s", v.Id, v.Name)
}

func main() {

	//------------------------
	emp := new(Employee)
	emp.SetIdEmployee(1)
	emp.SetNameEmployee("Vadim")

	emp.PrintInfo()
	/*
		pers := new(Person)
		pers.SetName("Vasily Romanov")

		var acc Account = Account{
			Id:   1,
			Name: "vasya",
			Person: Person{
				Id:   2,
				Name: "Vasily Romanov",
			},
		}

		acc.SetName("romanov.vasily")
		acc.Person.SetName("Test")
	*/
}
