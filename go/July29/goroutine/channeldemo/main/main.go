package main

import (
	"fmt"
	"sync"
	"time"
)

var mymap = make(map[int]int,10)
var lock sync.Mutex
func getNum(a int) {
	sum := 1
	for i := 1; i <= a;i++ {
		sum *= i
	}
	lock.Lock()
	mymap[a] = sum
	lock.Unlock()
}
func main()  {
	for i := 1; i <= 20;i++ {
		go getNum(i)
	}
	time.Sleep(time.Second)
	lock.Lock()
	for i, i2 := range mymap {
		fmt.Println(i,i2)
	}
	lock.Unlock()
}
