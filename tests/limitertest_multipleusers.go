package main

import (
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup

	var allowed atomic.Int64
	var rejected atomic.Int64

	start := make(chan struct{})

	for user := 1; user <= 10; user++ {
		userID := user

		for request := 0; request < 20; request++ {
			wg.Add(1)

			go func() {
				defer wg.Done()

				<-start

				url := fmt.Sprintf(
					"http://localhost:8080/checklimit/%d",
					userID,
				)

				resp, err := http.Get(url)
				if err != nil {
					fmt.Println("request failed:", err)
					return
				}
				defer resp.Body.Close()

				if resp.StatusCode == http.StatusOK {
					allowed.Add(1)
				} else if resp.StatusCode == http.StatusTooManyRequests {
					rejected.Add(1)
				}
			}()
		}
	}

	// Release all 200 goroutines.
	close(start)

	wg.Wait()

	fmt.Println("allowed:", allowed.Load())
	fmt.Println("rejected:", rejected.Load())
}
