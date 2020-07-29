package main

import "fmt"

type Cat struct {
	Name string
	Age int
}
func main() {
	var intChan chan int
	intChan = make(chan int, 3)
	//地址不同  管道中的数据
	fmt.Println(intChan,&intChan)
	//向管道写数据
	intChan<- 10
	num := 211
	intChan<- num
	//查看管道的长度和容量
	fmt.Println(len(intChan),cap(intChan))
	//从管道冲读取数据

	var num2 int
	num2 = <-intChan
	fmt.Println(num2,len(intChan))

	var mapChan chan map[string]string
	mapChan = make(chan map[string]string,10)
	m1 := make(map[string]string,20)
	m1["c1"] = "beijing"
	m1["c2"] = "shanghai"

	m2 := make(map[string]string,20)
	m2["c1"] = "chengdu"
	m2["c2"] = "chongqing"
	mapChan<-m1
	mapChan<-m2
	var allChan chan interface{}
	allChan = make(chan interface{},10)
	cat1 :=Cat{
		Name: "myj",
		Age:  12,
	}
	allChan<-cat1
	cat11 := <-allChan
	fmt.Println(cat11)
	//类型断言
	a := cat11.(Cat)
	fmt.Println(a.Name)
	close(allChan)
}