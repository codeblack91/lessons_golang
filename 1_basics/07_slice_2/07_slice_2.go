package main

import "fmt"

func main() {
	buf := []int{1, 2, 3, 4, 5}
	fmt.Println(buf)

	//получение среза, указывающего на ту же память
	sl1 := buf[1:4] // [2, 3, 4]
	sl2 := buf[:2]  // [1,2]
	sl3 := buf[2:]  //[3,4,5]
	fmt.Println(sl1, sl2, sl3)

	sl4 := buf[0:2]
	fmt.Println(sl4)

	newBuf := buf[:]
	newBuf[0] = 9

	//newBuf теперь указывает на другие данные
	newBuf = append(newBuf, 6)

	// buf    = [9, 2, 3, 4, 5], не изменился
	// newBuf = [1, 2, 3, 4, 5, 6], изменился
	newBuf[0] = 1
	fmt.Println("buf", buf)
	fmt.Println("newBuf", newBuf)

	var emptyBuf []int
	copied := copy(emptyBuf, buf)
	fmt.Println(copied, emptyBuf)

	// можно копировать в часть существующего слайса
	ints := []int{1, 2, 3, 4}
	copy(ints[1:3], []int{5, 6}) // ints = [1, 5, 6, 4]
	fmt.Println(ints)
}
