// go:build ignore
package main

import "fmt"

// https://qiita.com/ta1m1kam/items/fc798cdd6a4eaf9a7d5e#closing-channels
func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 5; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done //Block until we receive a notification from the worker on the channel.
	// https://gobyexample.com/channel-synchronization

	close(done) // don't need ?
}
