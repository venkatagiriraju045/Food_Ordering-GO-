package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	var count int = 1

	for i := 1; i <= num && count <= 3; i++ {
		if num%i == 0 {
			fmt.Println(i)
			count++

		}
	}
}
