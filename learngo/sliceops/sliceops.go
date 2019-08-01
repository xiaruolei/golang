package main

import "fmt"

func printSlice(s []int)  {
	fmt.Printf("%v, len=%d, cap=%d\n", s, len(s), cap(s))
}
func main() {
	/*
	Create slice
	 */
	var s []int
	for i := 0; i < 100; i++ {
		//fmt.Println(s)
		//printSlice(s)
		s = append(s, 2 * i + 1)
	}
	fmt.Println(s)

	s1 := []int{2, 4, 6, 8}
	printSlice(s1)

	s2 := make([]int, 10)
	printSlice(s2)

	s3 := make([]int, 10, 32)
	printSlice(s3)

	/*
	Copy slice
	 */
	copy(s2, s1) //s1 -> s2
	printSlice(s2)

	/*
	Delete elements from slice
	 */
	fmt.Println("Delete elements from slice")
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	/*
	Popping from front
	 */
	fmt.Println("Popping from front")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front)
	printSlice(s2)

	/*
	Popping from tail
	 */
	fmt.Println("Popping from tail")
	tail := s2[len(s2) - 1]
	s2 = s2[:len(s2) - 1]
	fmt.Println(tail)
	printSlice(s2)

}
