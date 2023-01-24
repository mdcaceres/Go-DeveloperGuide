package main

import "fmt"

func main() {
	integers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	for _, num := range integers {
		if num%2 == 0 {
			fmt.Println("EVEN")
		} else {
			fmt.Println("ODD")
		}
	}
}
