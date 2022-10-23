package util

import "fmt"

func DividerPrinter(num int) {
	for i := 0; i < num; i++ {
		fmt.Print("-")
	}

	fmt.Println()
}