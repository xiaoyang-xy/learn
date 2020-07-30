package main

import "fmt"

func main() {
	intChan := make(chan int, 10)
	intChan <- 100
	intChan <- 200
	intChan <- 300
	intChan <- 400
	intChan <- 500
	//关闭管道 只能读不能写
	close(intChan)
	//不关闭管道遍历管道时  fatal error: all goroutines are asleep - deadlock!
	for i := range	intChan {
		fmt.Print("i = ",i," ")
	}
	//close(intChan)
}
