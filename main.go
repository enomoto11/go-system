package main

import "time"

func main() {
	duration := 10 * time.Second

	channel := time.After(duration)

	println("start")
	<-channel

	println("end")
}
