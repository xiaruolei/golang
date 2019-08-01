package main

import "fmt"

func main() {
	/*
	key does not contain slice, map, function
	 */

	/*
	Create map
	 */
	m := map[string]string {
		"name": "ccmouse",
		"course": "golang",
		"site": "imooc",
		"quality": "notbad",
	}

	m2 := make(map[string]int) // m2 == empty map

	var m3 map[string]int // m3 == nil

	fmt.Println(m, m2, m3)

	/*
	Traversing map
	 */

	fmt.Println("Traversing map: ")
	for k, v := range(m) {
		fmt.Println(k, v)
	}

	//fmt.Println("Getting values")
	//courseName, ok := m["course"]
	//fmt.Println(courseName, ok)
	//caurseName, ok := m["caurse"]
	//fmt.Println(caurseName, ok)

	if causeName, ok := m["cause"]; ok {
		fmt.Println(causeName)
	} else {
		fmt.Println("Key does not exist")
	}

	/*
	Delete values
	 */

	delete(m, "name")



}
