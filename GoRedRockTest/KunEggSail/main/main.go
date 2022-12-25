package main

import "fmt"

func main() {
	var n int
	fmt.Print("请输入船的个数")
	fmt.Scanln(&n)
	var time map[int]int
	var ship map[int]map[int]int
	var kind = make(map[int]int, 100000)
	time = make(map[int]int, n)
	ship = make(map[int]map[int]int)
	var t int
	var k int
	for i := 0; i < n; i++ {

		fmt.Print("该船的出发时间")
		fmt.Scanln(&t)
		time[i] = t
		fmt.Print("该船蛋的批数")
		fmt.Scanln(&k)
		ship[i] = make(map[int]int, k)
		for j := 0; j < k; j++ {
			var kind int
			if j < k-1 {
				fmt.Scan(&kind)
			} else {
				fmt.Scanln(&kind)
			}
			ship[i][j] = kind
		}
	}
	var count = 1
	for i := 0; i < n; i++ {
		start := find(time, time[i])
		kind[0] = ship[start][0]
		for j := start; j < n; j++ {
			len := len(ship[j])
			for m := 0; m < len; i++ {
				count = Judge(ship[j][m], kind, count)
			}
		}
	}
	fmt.Printf("count=%v\n", count)
}
func find(time map[int]int, now int) int {
	for i, v := range time {
		if now-86400 <= 0 {
			return 0
		} else {
			if v >= now-86400 && v <= now {
				if v == now {
					return i
				}
			}
		}
	}
	return 0
}
func Judge(n int, kind map[int]int, count int) int {
	for _, v := range kind {
		if v == n {
			count++
			kind[count-1] = n
			break
		}
	}
	return count
}
