package main

import "fmt"

func main() {
	//默认管道可以读写
	//var  chan1 chan int
	//声明为只写
	var chan2 chan<- int
	chan2 = make(chan int, 3)
	chan2 <- 20
	//num := <-chan2 error : 只写 不能读
	fmt.Print(chan2)
	//声明为只读
	var chan3 <-chan int
	chan3 = make(chan int,3)
	num2 := <-chan3
	//chan3 <- 30 error 只读不能写
	fmt.Print(num2)

	//只读只写只存在于函数调用

}
