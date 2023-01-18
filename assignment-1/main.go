package main

import "fmt"

func main() {
	s10 := []int{}
	for i := 0; i < 11; i++ {
		s10 = append(s10, i)
	}
	fmt.Println(s10)
	for _, num := range s10 {
		if num%2 == 0 {
			fmt.Printf("%v is even\n", num)
		} else {
			fmt.Printf("%v is odd\n", num)
		}
	}
}
