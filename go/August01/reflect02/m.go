package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num = 100
	fn := reflect.ValueOf(&num)
	fn.Elem().SetInt(200)
	fmt.Println(num)
}