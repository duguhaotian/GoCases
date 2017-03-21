// GoSort project main.go
package main

import (
	"GoSort/utils"
	"fmt"
)

func main() {

	data := []int{3, 25, 37, 41, 2, 8, 7, 9, 14, 32}
	insert := utils.InsertSort{}
	for _, item := range data {
		fmt.Print(item, " ")
	}

	insert.Sort(data)
	fmt.Println("\nsort result: ")

	for _, item := range data {
		fmt.Print(item, " ")
	}
	fmt.Println()

}
