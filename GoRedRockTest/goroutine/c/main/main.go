package main

import (
	"fmt"
	"time"
)

func powerf(x int, n int) int {
	ans := 1

	for n > 0 {
		ans *= x
		n--
	}
	return ans
}

func putNUm(NumChan chan int) {
	for i := 1; i <= 999999; i++ {
		NumChan <- i
	}
	close(NumChan)
}
func Power(NumChan chan int, PowerChan chan int, exitchan chan bool) {
	//var flag bool
	for {
		time.Sleep(time.Millisecond * 5)
		num, ok := <-NumChan
		//flag = false
		if !ok {
			break
		}
		var n = num
		var t, count = 1, 0
		switch true {
		case n < 10:
			t = 1
		case n >= 10 && n < 100:
			t = 2
		case n >= 100 && n < 1000:
			t = 3
		case n >= 1000 && n < 10000:
			t = 4
		case n >= 10000 && n < 100000:
			t = 5
		case n >= 100000 && n < 1000000:
			t = 6
		}
		for i := 0; i < t; i++ {
			count += powerf(n%10, t)
			n = n / 10
		}
		if count == num {
			PowerChan <- num
		}
	}
	//fmt.Println("有一个协程取不到而退出")
	exitchan <- true
}
func main() {
	NumChan := make(chan int, 90000)
	PowerChan := make(chan int, 40000)
	exitchan := make(chan bool, 2000)
	start := time.Now().Unix()
	go putNUm(NumChan)
	for i := 0; i < 2000; i++ {
		go Power(NumChan, PowerChan, exitchan)
	}
	go func() {
		for i := 0; i < 2000; i++ {
			<-exitchan
		}
		close(PowerChan)
		end := time.Now().Unix()
		fmt.Printf("耗时=%v\n", end-start)
	}()
	for {
		res, ok := <-PowerChan
		if !ok {
			break
		}
		fmt.Printf("自幂数有%v\n", res)
	}
	fmt.Println("main线程退出")
}
