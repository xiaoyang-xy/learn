package main

import (
	"fmt"
	"time"
)

func putNum(intChan chan int)  {
	for i := 1; i <= 8000; i++ {
		intChan <- i
	}
	close(intChan)
}

func primeNum(intChan chan int ,primeChan chan int,exitChan chan bool){
	var flag bool
	for true {
		num,ok := <- intChan
		if !ok{
			break
		}
		flag = true
		for i := 2; i < num; i++ {
			if num % i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeChan <- num
		}
	}
	fmt.Print("exit")
	//这个时候不能关闭primeChan 可能还有协程在运行
	exitChan <- true
}

func main() {
	start := time.Now().Unix()
	intChan := make(chan int,1000)
	primeChan := make(chan int,2000)//放入结果
	//标示退出的管道
	exitChan := make(chan bool,4)
	//开启一个协程 向intChan存数据
	go putNum(intChan)
	//开启四个协程，从intChan取数据 并判断是否为素数，如果是 放入primeChan
	for i := 0; i < 4; i++ {
		go primeNum(intChan,primeChan,exitChan)
	}

	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		close(primeChan)
	}()
	for true {
		res, ok := <- primeChan
		if !ok{
			break
		}
		fmt.Print(res,"  ")
	}
	end := time.Now().Unix()
	fmt.Printf("%v %v",end,start)


}
