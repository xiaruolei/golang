package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string)  {
	//base := "/Users/ruolei/go/src/imooc/learngo/loop/"
	base := "loop/"
	file, err := os.Open(base + filename)
	if err != nil {
		panic(err)
	}

	printFileContents(file)
}

func printFileContents(reader io.Reader)  {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	fmt.Println(convertToBin(5),
		convertToBin(13),
		convertToBin(0),
		)

	printFile("abc.txt")
	s := `abc"d"
	kkkk
	123

	p`
	printFileContents(strings.NewReader(s))
}
