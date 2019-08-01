package main

import "fmt"

func updateSlice(s []int)  {
	s[0] = 100
}
func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 8, 9}
	s1 := arr[2:6]
	/*
	fmt.Println(arr)
	fmt.Println(s)

	updateSlice(s)
	fmt.Println(arr)
	fmt.Println(s)
	 */

	fmt.Println("arr: ", arr)
	fmt.Println("s1: ", s1)

	s2 := append(s1, 12)
	fmt.Println("arr: ", arr)
	fmt.Println("s1: ", s1)
	fmt.Println("s2: ", s2)


	s3 := append(s2, 13)
	fmt.Println("arr: ", arr)
	fmt.Println("s1: ", s1)
	fmt.Println("s2: ", s2)
	fmt.Println("s3: ", s3)
	//s2 := s1[3:5]
	//fmt.Println("s2: ", s2)
	//s3 := s2[0:3]
	//fmt.Println("s3: ", s3)

	//fmt.Printf("s1: %v, len(s1): %d, cap(s1): %d\n", s1, len(s1), cap(s1))
	//fmt.Printf("s2: %v, len(s2): %d, cap(s2): %d\n", s2, len(s2), cap(s2))


}
