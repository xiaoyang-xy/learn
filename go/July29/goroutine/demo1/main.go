package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

func test() {
	for i := 1; i <= 10;i++ {
		fmt.Println("hello world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
func main() {
	fmt.Println(runtime.NumCPU())
	go test()
	for i := 1; i <= 10 ;i++ {
		fmt.Println("main hello world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	 }

}
