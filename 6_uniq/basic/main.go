package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	var prev string
	for in.Scan() {
		txt := in.Text()
		if txt == prev {
			continue
		}

		if txt < prev { //если в консоли вывели число меньше предыдущего, то будет паника
			panic("file not serted")
		}
		prev = txt
		fmt.Println(txt)
	}
}
