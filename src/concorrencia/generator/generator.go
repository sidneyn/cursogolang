package main

import (
	"cursogolanghtml/html"
	"fmt"
	_ "io/ioutil"
	_ "net/http"
)

// Google I/O 2021 - Go Concurrency Patterns

// <-chan - canal somente-leitura
/*
func titulo (urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)
			fmt.Println(html)
			//	r, _ := regexp.Compile("<title>(.*?)<\\/title>")
		//	c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}
	return c
}
*/
func main() {
	t1 := html.Titulo("https://www.cod3r.com.br", "https://www.google.com")
	t2 := html.Titulo("https://www.amazon.com.br", "https://www.youtube.com")

	fmt.Println("Primeiros: ", <-t1, "|", <-t2)
	fmt.Println("Segundos: ", <-t1, "|", <-t2)
}
