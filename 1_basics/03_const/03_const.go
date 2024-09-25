package main

import "fmt"

const pi = 3.141
const (
	hello = "Привет"
	e     = 2.718
)

const (
	zero = iota
	_    //пустая, переменая, пропуск iota
	two
	three // = 3
)
const (
	_         = iota             //пропускам первое значение
	KB uint64 = 1 << (10 * iota) // 1 << (10 * 1) = 1024
	MB                           // 1 << (10 * 2) = 1048576
)
const (
	//нетипизированная константа
	year = 2017
	//типизированная константа
	yearTyped int = 2017
)

func main() {
	var month int32 = 13
	fmt.Println(month + year)

	fmt.Printf("pi: %f", pi) //разобать вывод в fmt float и uint64
	fmt.Println()

	fmt.Printf("const hello: %s, const e: %f", hello, e)
	fmt.Println()

	fmt.Printf("zero: %d, two: %d, three %d", zero, two, three)
	fmt.Println()

	fmt.Printf("KB: %b, MB: %d", KB, MB)
	fmt.Println()

	fmt.Print(yearTyped)
	fmt.Println()
}
