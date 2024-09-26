package main

import "fmt"

func main() {
	/*
		//создание
		var buf0 []int             //len=0, cap=0
		buf1 := []int{}            //len=0, cap=0
		buf2 := []int{42}          //len=1, cap=1
		buf3 := make([]int, 0)     //len=0, cap=0
		buf4 := make([]int, 5)     //len=5, cap=5
		buf5 := make([]int, 5, 10) //len=5, cap=10

		println(buf0, buf1, buf2, buf3, buf4, buf5)

		//обращение к элементам
		someInt := buf2[0]

		//ошибка при выполнении
		//panic: runtime error: index out of range
		//someOtherInt := buf2[1]

		fmt.Println(someInt)

		fmt.Println(buf5[3]) //Что такое capacity как оно себя ведет, изучить?

		//добавление элементов
		var buf []int
		buf = append(buf, 9, 10) //len=2, cap=2
		buf = append(buf, 12)    //len=3. cap=4

		//добавление другого слайса
		otherBuf := make([]int, 3)
		buf = append(buf, otherBuf...) // len=6, cap=8

		fmt.Println(buf, otherBuf)

		//просмотр информации о слайсе
		var bufLen, bufCap int = len(buf), cap(buf)

		fmt.Println(bufLen, bufCap)

		//----------------
		var bufTest []int
		bufTest = append(bufTest, 1, 2, 3)
		fmt.Println(bufTest)

		bufTest2 := []int{4, 5, 6}
		fmt.Println(bufTest2)

		bufTest3 := append(bufTest, bufTest2...)
		fmt.Println(bufTest3)

		bufTestLen, bufTestCap := len(bufTest), cap(bufTest)
		fmt.Println(bufTestLen, bufTestCap)
	*/
	//прочиать про lenth и capacity

	/*
		s := make([]int, 3, 5)
		fmt.Println("До добавления:", len(s), cap(s)) // Длина: 3, Емкость: 5

		s = append(s, 10, 20)                            // Добавляем элементы, но не превышаем емкость
		fmt.Println("После добавления:", len(s), cap(s)) // Длина: 5, Емкость: 5

		s = append(s, 30)                                        // Теперь емкость будет превышена, произойдет перераспределение
		fmt.Println("После второго добавления:", len(s), cap(s)) // Длина: 6, Емкость: 10

		s = append(s, 10, 20, 30, 40, 50, 60)
		fmt.Println("После второго добавления:", len(s), cap(s)) //емкасть увеличивается на x2
	*/

	//удаление из среза
	s := []int{1, 2, 3, 4, 5}
	indexToRemove := 2 // Удалим элемент с индексом 2 (число 3)

	s = append(s[:indexToRemove], s[indexToRemove+1:]...) // Соединяем части среза

	fmt.Println("После удаления:", s) // [1 2 4 5]

	//посмотреть в chatgpt примеры удаления значений в срезах
}
