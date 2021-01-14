package main

import (
	"fmt"
	"time"
)

func main() {
	go goodCount("dave")
	goodCount("flynn")
}

func goodCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is Good", i)
		time.Sleep(time.Second)
	}
}
