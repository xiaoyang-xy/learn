package main

import (
	"fmt"
	"time"
)

func sayHello()  {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello world")
	}
}

func test()  {
	defer func() {
		if err := recover(); err != nil	{
			fmt.Println(err)
		}
	}()
	var mymap map[int]string
	//mymap = make(map[int]string,10)
	mymap[0] = "qihongwu"
	fmt.Print(mymap)
}
func main() {
	go sayHello()
	go test()
	time.Sleep(time.Second * 10)
}
