package main

import (
	"fmt"
	"time"
)

func fetchUser() string {
	time.Sleep(time.Millisecond*100);
	return "Vikash"
}

func fetchUserLikes(id string) int {
	time.Sleep(time.Millisecond*150);
	return 10
}

func fetchUserMatch(id string) string {
	time.Sleep(time.Millisecond*100);
	return "unknown"
}

// sync pattern

func main() {
	start :=  time.Now()

	userName := fetchUser()
	likes := fetchUserLikes(userName)
	match := fetchUserMatch(userName)

	fmt.Println("likes : ",likes)
	fmt.Println("match : ",match)
	fmt.Println("time of completetion : ",time.Since(start))

}