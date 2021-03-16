package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan bool, 1)

	go func(ch chan bool) {
		for true {
			time.Sleep(time.Second)
			ch <- true
		}
	}(ch)
	go func(ch chan bool) {
		for true {
			<-ch
			fmt.Println("1 consume")
		}
	}(ch)
	go func(ch chan bool) {
		for true {
			<-ch
			fmt.Println("2 consume")
		}
	}(ch)

	time.Sleep(time.Hour)
}
