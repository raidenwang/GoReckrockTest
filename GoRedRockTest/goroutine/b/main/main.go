package main

import (
	"fmt"
	"strconv"
)

type Goroutine struct {
	name string
}

func main() {
	var Chan_1 chan Goroutine
	Chan_1 = make(chan Goroutine, 4)
	str1 := "goroutine"
	str2 := "successful"
	for i := 0; i < 4; i++ {
		var a Goroutine
		if i < 3 {
			a.name = str1 + strconv.Itoa(i+1)
		} else {
			a.name = str2
		}
		Chan_1 <- a
	}
	close(Chan_1)
	for v := range Chan_1 {
		fmt.Println(v)
	}
}
