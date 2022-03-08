package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// Aula Google I/O 2012 - Go Concurrency Patterns

// <-chan - canal de somente leitura

func Titulo(urls ...string) <-chan string {
	c := make(chan string)

	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}
	return c
}

func main() {
	t1 := Titulo("https://www.amazon.com", "https://www.youtube.com")

	t2 := Titulo("https://www.cod3r.com.br", "https://www.google.com")

	fmt.Println("Primeiros: ", <-t1, "|", <-t2)
	fmt.Println("Segundos: ", <-t1, "|", <-t2)

}
