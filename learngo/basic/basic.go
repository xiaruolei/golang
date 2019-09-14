package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	aa = 3
	ss = "kkk"
	bb = true
)

func variableZeroValue()  {
	var a int
	var s string
	fmt.Printf("%d, %q\n", a, s)
}

func variableIniticalValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func varialeTypeDeduction()  {
	var a, b, c, d = 3, 4, true, "def"
	fmt.Println(a, b, c, d)
}

func variableShorter()  {
	a, b, c, d := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, d)
}

func euler()  {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
	fmt.Println(cmplx.Pow(math.E, 1i * math.Pi) + 1)
	fmt.Printf("%3f", cmplx.Exp(1i * math.Pi) + 1)
}

func triangle()  {
	var a, b int = 3, 4
	fmt.Println(calTriangle(a, b))
}

func calTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a * a + b * b)))
	fmt.Println(c)
	return c
}

func consts()  {
	const (
		filename  = "abc.txt"
		a, b = 3, 4
	)
	var c int
	c = int(math.Sqrt(a * a + b * b))
	fmt.Println(filename, c)
}

func enums()  {
	const (
		cpp = iota
		_
		python
		golang
		javascript
	)
	fmt.Println(cpp, python, golang, javascript)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("Hello World")
	variableZeroValue()
	variableIniticalValue()
	varialeTypeDeduction()
	variableShorter()
	println(aa, ss, bb)
	euler()
	triangle()
	consts()
	enums()
}
