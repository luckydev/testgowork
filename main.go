package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)
	c4 := make(chan int)
	c5 := make(chan int)
	c6 := make(chan int)
	c7 := make(chan int)

	start := time.Now()

	go work("w1", 2, 10, c1)
	go work("w2", 1, 15, c2)
	go work("w3", 1, 13, c3)
	go work("w4", 2, 7, c4)
	go work("w5", 1.0, 18, c5)
	go work("w6", 1.0, 18, c6)
	go work("w7", 1.0, 18, c7)

	<-c1
	<-c2
	<-c3
	<-c4
	<-c5
	<-c6
	<-c7

	fmt.Println("Time to complete: " + time.Since(start).String())
}

func work(worker string, waiting int64, limit int, end chan int) {
	var latest int
	for i := 0; i < limit; i++ {
		st := fmt.Sprintf("%v printing for %d time", worker, i)
		latest = i
		fmt.Println(st)
		time.Sleep(time.Duration(waiting) * time.Second)
	}

	fmt.Println(">>> " + worker + " is done")
	end <- latest
}
