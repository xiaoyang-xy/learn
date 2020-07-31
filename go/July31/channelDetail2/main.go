package main

import "fmt"

func main() {
	//使用select解决管道去数据阻塞问题
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprint(i)
	}
	//在实际项目中，可能不好确定关闭那些管道防止阻塞
	//使用 }select可以解决
	for{
		select {
		case v := <-intChan :  //
			fmt.Print(v,"  ")
		case v := <- stringChan :
			fmt.Print(v,"  ")
		default:
			fmt.Print("qubudao")
			return

		}
	}
}
