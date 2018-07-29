package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// Procurar sobre Go Concurrency Patterns

//  <-chan = somente leitura

func getTitle(urls ...string) <-chan string {
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
	t1 := getTitle("https://www.cod3r.com.br", "https://www.google.com")
	t2 := getTitle("https://www.amazon.com.br", "https://www.youtube.com")

	fmt.Println("Primeiro:", <-t1, "|", <-t2)
	fmt.Println("Segundo:", <-t1, "|", <-t2)

}
