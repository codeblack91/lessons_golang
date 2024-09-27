package main

import "fmt"

func getSomeVars() string {
	fmt.Println("getSomeVars execution")
	return "getSomeVars result"
}

// дефер работает вверх по стеку вызовов
func main() {
	defer fmt.Println("After work") //дергается вторым
	defer func() {                  //дергается првым
		fmt.Println(getSomeVars())
	}()
	fmt.Println("Some userful work")
}
