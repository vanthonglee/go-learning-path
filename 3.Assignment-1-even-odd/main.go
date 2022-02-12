package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, v := range s {
		r := ""
		if v%2 == 0 {
			r = "even"
		} else {
			r = "odd"
		}
		fmt.Printf("%d is %v \n", v, r)
	}
}
