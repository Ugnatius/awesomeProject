package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i <= 5; i++ {
		fmt.Println("New itteration ", i)
		time.Sleep(500 * time.Millisecond)
	}
}
