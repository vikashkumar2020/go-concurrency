package main

import (
	"fmt"
	"sync"
	"time"
)

func fetchUser() string {
	time.Sleep(time.Millisecond * 100)
	return "Vikash"
}

func fetchUserLikes(id string, channel chan any, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Millisecond * 150)
	channel <- 11
}

func fetchUserMatch(id string, channel chan any, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Millisecond * 100)
	channel <- "unknown"
}

func main() {
	start := time.Now()
	userName := fetchUser()
	respch := make(chan any,2) // buffered

	var wg sync.WaitGroup
	wg.Add(2)

	go fetchUserLikes(userName, respch, &wg)
	go fetchUserMatch(userName, respch, &wg)

		wg.Wait()
		close(respch)

	for resp := range respch {
		fmt.Println("resp", resp)
	}

	fmt.Println("time of completion : ", time.Since(start))
}
