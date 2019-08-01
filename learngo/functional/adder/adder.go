package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

func main() {
	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("0 + 1 + ... + %d = %d\n", i, a(i))
	}

	fmt.Println(a(1))
	fmt.Println(a(2))
	fmt.Println(a(3))
	fmt.Println(a(4))
}
