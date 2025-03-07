package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Poetry Real-Time Service is Running")
	for {
		time.Sleep(time.Hour)
	}
}
