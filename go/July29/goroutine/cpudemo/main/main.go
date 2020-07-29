package main

import (
	"fmt"
	"runtime"
)

func main() {
	cpuNum := runtime.NumCPU()
	fmt.Println(cpuNum)
	runtime.GOMAXPROCS(cpuNum)

}