package main

import (
	"fmt"
	"reflect"
)

func test(b interface{})  {
	//return 值  类型interface
	rVal := reflect.ValueOf(b)
	rTyp := reflect.TypeOf(b)
	fmt.Println(rVal,rTyp,rVal.Kind(),rTyp.Kind())
	//value.kind()  return type
	//return b的值  return interface
	iVal := rVal.Interface()
	v := iVal.(int)
	fmt.Printf("%T",v)
}

type Student struct {
	Age int
	Name string
}
func testStruct(b interface{})  {
	//return 值  类型interface
	rVal := reflect.ValueOf(b)
	fmt.Println(rVal)
	//return b的值  return interface
	iVal := rVal.Interface()
	fmt.Println(rVal.Kind())	//return struct
	v := iVal.(Student)
	fmt.Printf("%T",v)
	fmt.Println(v.Age,v.Name)
}
func main() {
	//变量、接口、 reflect.Value是可以相互转换
	a := 1
	student := Student{
		Age:  1,
		Name: "1",
	}
	test(a)
	testStruct(student)
	}
