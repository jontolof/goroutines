package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	fmt.Println("Start")
	var wg sync.WaitGroup
	urls := []string{
		"http://www.aftonbladet.se",
		"http://www.apple.com",
		"http://www.svd.se",
	}
	for _, url := range urls {
		wg.Add(1)
		go fetchURL(url, &wg)
	}
	wg.Wait()
}

func fetchURL(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("Status Code:", resp.StatusCode)
}
