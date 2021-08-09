package main

import (
	"fmt"
	"time"
)

func say(c chan bool) {
	for i := 0; i < 5; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(i)
	}

	c <- true // this is required to wait to finished all routines
}

func main() {
	channel := make(chan bool)

	fmt.Println("start go routine")
	go say(channel)
	fmt.Println("end go routine")

	b := <-channel
	fmt.Println(b)
}
