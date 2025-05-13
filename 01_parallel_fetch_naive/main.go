package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	urls := []string{
		"http://www.aftonbladet.se",
		"http://www.apple.com",
		"http://www.svd.se",
	}
	for _, url := range urls {
		fetchURL(url, &wg)
	}
	fmt.Println("Duration:", time.Since(start))
}

func fetchURL(url string, wg *sync.WaitGroup) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("Status Code:", resp.StatusCode)
}
