package main

import (
	"fmt"
	"imooc/learngo/retriever/mooc"
	"imooc/learngo/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
} 

func downloade(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

func inspect(r Retriever)  {
	fmt.Printf("%T, %v\n", r, r)
	switch v := r.(type){
	case mooc.Retriver:
		fmt.Println(v)
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println(v)
		fmt.Println("UserAgent:", v.UserAgent)
	}


}

func main() {
	var r Retriever
	r  = mooc.Retriver{"this is fake imooc.com"}
	inspect(r)

	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut: time.Minute,
	}
	inspect(r)
	//fmt.Println(downloade(r))


	// Type assertion
	if mockRetriver, ok := r.(mooc.Retriver); ok {
		fmt.Println(mockRetriver.Contents)
	} else {
		fmt.Println("not a mook retriver")
	}




}
