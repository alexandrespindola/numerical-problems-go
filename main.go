package main

import "fmt"

func main() {
	numbersList := []int{}
	for i := 1; i <= 100 ; i++ {
		if i % 3 == 0 {
			numbersList = append(numbersList, i)
		}
	}
	fmt.Println(numbersList)
}