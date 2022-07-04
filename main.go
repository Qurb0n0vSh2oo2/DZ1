package main

import (
	"fmt"
	"time"
)

func main() {
	Minute := 10

	go Tack(Minute)
	go Tick(Minute)

	time.Sleep(time.Second * time.Duration(Minute))
}

func Tick(Minute int) {
	for i := 0; i <= Minute; i++ {
		fmt.Println("tick")
		<-time.After(time.Second)
	}
}

func Tack(Minute int) {
	for i := 0; i <= Minute/5; i++ {
		<-time.After(time.Second * 5)
		fmt.Println("tack")
	}
}