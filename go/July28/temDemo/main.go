package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(len(os.Args))
	for i, i2 := range os.Args {
		fmt.Println(i,i2)
	}
}
