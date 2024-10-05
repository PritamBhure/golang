package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {

	fmt.Printf("Worker %d Started work\n", i)
	fmt.Printf("Worker %d End work\n", i)
	defer wg.Done() // kam khatAM HONE KE BAD wo sync method ko signal dega kham khatam
}

func main() {
	fmt.Println("Learning Sync waitgroup")

	var wg sync.WaitGroup
	//started 3 workers go routine

	for i := 1; i <= 3; i++ {
		wg.Add(1) // increament the wait group count
		go worker(i, &wg)
	}

	fmt.Println("Worker task completed ")
}
