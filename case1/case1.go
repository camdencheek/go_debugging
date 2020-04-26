package main

import "fmt"

func main() {
	a := make([]int, 5)
	b := []int{0, 1, 2, 3, 4}

	for _, val := range b {
		a = append(a, val)
	}

	fmt.Printf("%#v\n", b)
}
