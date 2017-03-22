// GoSort project main.go
package main

import (
	"GoSort/utils"
	"fmt"
)

func main() {

	//	insert := utils.InsertSort{}
	//	bubble := utils.BubbleSort{}
	quick := utils.QuickSort{}
	data := []int{3, 25, 37, 41, 2, 8, 7, 9, 14, 32}
	for _, item := range data {
		fmt.Print(item, " ")
	}

	//	insert.Sort(data)
	//	bubble.Sort(data)
	quick.Sort(data)
	fmt.Println("\nsort result: ")

	for _, item := range data {
		fmt.Print(item, " ")
	}
	fmt.Println()

}
