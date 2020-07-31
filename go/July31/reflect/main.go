package main

import (
	"fmt"
	"reflect"
)

func test(b interface{})  {
	rVal := reflect.ValueOf(b)
	iVal := rVal.Interface()
	v := iVal.(int)
	fmt.Println(v)
}
func main() {
	var a int = 100
	var str string = "123"
	fmt.Println(reflect.ValueOf(a),reflect.TypeOf(a))
	fmt.Println(reflect.ValueOf(str),reflect.TypeOf(str))
	test(a)
}