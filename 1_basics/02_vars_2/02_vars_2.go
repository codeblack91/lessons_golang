package main

import "fmt"

func main() {
	//int - платформозависимый тип, 32/64
	var i int = 10
	fmt.Printf("i: %d", i)
	fmt.Println()

	var autoInt = -10
	fmt.Println(autoInt)

	//int8, int16, int32, int64
	var bigInt int64 = 1<<32 - 1 //побитовая операция? что это?
	fmt.Print(bigInt)
	fmt.Println()

	//платформозависимый тип, 32/64
	var unsignedInt uint = 123 //uint это 32 битный тип, вроде у него меньше емкость, загугли
	fmt.Printf("unsignedInt: %d", unsignedInt)
	fmt.Println()

	//uint8, uint16, uint32, uint64
	var unsignedBigInt uint64 = 1<<64 - 1 //вопрос
	unsignedBigInt = 0x01fe               //вопрос
	unsignedBigInt = 0b00001010101        //вопрос
	fmt.Println(i, autoInt, bigInt, unsignedInt, unsignedBigInt)

	//float32, float64
	var pi float32 = 3.141
	var e = 2.718
	goldRatio := 1.618
	fmt.Println(pi, e, goldRatio)

	//bool
	var b bool // false по-умолчанию
	var isOk bool = true
	var success = false
	cond := true

	fmt.Println(b, isOk, success, cond)

	//complex64, complex128
	var c complex128 = -1.1 + 7.1i //что есть комплексные числа? не понимаю
	c2 := -1.1 + 7.12i

	fmt.Println(c, c2)

}
