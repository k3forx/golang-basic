package main

import (
	"log"
	"math/rand"
	"time"
)

const numPool = 1000

func PrintText(s string) {
	log.Println(s)
}

func RandomNumber(n int) int {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(n)
	return value
}

func CalculateValue(initChan chan int) {
	randomNumber := RandomNumber(numPool)
	initChan <- randomNumber
}

func main() {
	PrintText("Hi")

	initChan := make(chan int)
	defer close(initChan)

	go CalculateValue(initChan) // Channels are convenient way to pass values from one package to another package

	num := <-initChan
	log.Println(num)
}
