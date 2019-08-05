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

type Poster interface {
	Post(url string, form map[string]string) string
}

const url  = "http://www.imooc.com"
func downloade(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster)  {
	poster.Post(url, map[string]string{
		"name": "ccmouse",
		"course": "golang",
	})
}

type RetriverPoster interface {
	Retriever
	Poster
}

func session(s RetriverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another fake imook.com",
	})
	return s.Get(url)
}

func inspect(r Retriever)  {
	fmt.Println("inspecting", r)
	fmt.Printf("> %T, %v\n", r, r)
	fmt.Print("> Type Switch")
	switch v := r.(type){
	case *mooc.Retriver:
		fmt.Println(v)
		fmt.Println("> Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println(v)
		fmt.Println("> UserAgent:", v.UserAgent)
	}
	fmt.Println()
}

func main() {

	retriver  := mooc.Retriver{"this is fake imooc.com"}
	var r Retriever
	r = &retriver
	inspect(r)

	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut: time.Minute,
	}
	inspect(r)
	//fmt.Println(downloade(r))


	// Type assertion
	if mockRetriver, ok := r.(*mooc.Retriver); ok {
		fmt.Println(mockRetriver.Contents)
	} else {
		fmt.Println("not a mook retriver")
	}

	fmt.Println("Try a session:")
	fmt.Println(session(&retriver))



}
