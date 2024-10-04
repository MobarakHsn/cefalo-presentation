package main

import (
	"fmt"
	"net/http"
	"sync"
)

func sendRequest(wg *sync.WaitGroup, url string) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("Response Status Code: %d\n", resp.StatusCode)
}

func main() {
	url := "http://localhost:8080/" // Adjust this to your endpoint
	n := 1000                       // Number of concurrent requests
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go sendRequest(&wg, url)
	}

	wg.Wait()
	fmt.Println("All requests completed.")
}
