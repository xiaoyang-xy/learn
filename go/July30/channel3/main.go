package main

import "fmt"

var i int
func writeData(numChan chan int){
	for i = 1; i <= 10; i++ {
		numChan <- i
	}
	close(numChan)
}

func getNum(num int) int{
	res := 0
	for b := 1; b <= num; b++ {
		res += b
	}
	return res
}
func readData(numChan chan int,resChan chan int,exitChan chan bool){
	for {
		a, ok := <-numChan
		if !ok {
			break
		}
		res := getNum(a)
		resChan <- res

	}
	close(resChan)
	exitChan <- true
	close(exitChan)
}

func main() {
	resChan := make(chan int,10)
	numChan := make(chan int, 10)
	exitChan := make(chan bool,1)
	go writeData(numChan)
	for true {
		if i == 10 {
			break
		}
	}
	go readData(numChan,resChan,exitChan)
	for true {
		_,ok := <- exitChan
		if !ok{
			break
		}
	}
	for i2 := range resChan {
		fmt.Printf("%d ",i2)
	}

}
