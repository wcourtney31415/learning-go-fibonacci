package main

import "fmt"

func main() {
	list := []int{0, 1}
	for i := 0; i < 20; i++ {
		var listLength = len(list)
		var last = listLength - 1
		var secondToLast = listLength - 2
		var sum = last + secondToLast
		list = append(list, sum)
	}
	fmt.Println(list)
}
