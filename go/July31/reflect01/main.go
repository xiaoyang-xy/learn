package main

import (
	"fmt"
	"reflect"
)

func reflectTest(b interface{}){
	rTyp := reflect.TypeOf(b)
	fmt.Println(rTyp)
	rVal := reflect.ValueOf(b)
	fmt.Printf("%v %T",rVal,rVal)
}
func main() {
	var a  = 10
	reflectTest(a)
}