package main

import (
	"fmt"
	"reflect"
)

type Cal struct {
	Num1 int
	Num2 int
}

func (this Cal) GetSub(n1,n2 int) int {
	return n1 - n2
}

func testStruct(b interface{})  {
	typ := reflect.TypeOf(b)
	val := reflect.ValueOf(b)
	kd := val.Kind()
	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}
	num := val.NumField()
	fmt.Printf("struct has %d fields\n",num)
	for i := 0; i < num; i++ {
		fmt.Println(i,val.Field(i))
		tagVal := typ.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Println(i,tagVal)
		}
	}
	numOfMethod := val.NumMethod()
	fmt.Println(numOfMethod)
	//Method 根据字典序排序
	var params []reflect.Value
	params = append(params,reflect.ValueOf(10))
	params = append(params,reflect.ValueOf(40))
	res := val.Method(0).Call(params)
	fmt.Println(res[0].Int())
}

func main() {
	var a Cal = Cal{
		Num1: 10 ,
		Num2: 5 ,
	}
	testStruct(a)

}
