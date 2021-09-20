package main

import "fmt"

func main() {
	s := make([]int, 2)
	s[0] = 0
	s[1] = 1
	for i := 0; i < 20; i++ {
		var a = len(s) - 1
		var b = len(s) - 2
		var sum = a + b
		s = append(s, sum)
	}
	fmt.Println(s)
}
