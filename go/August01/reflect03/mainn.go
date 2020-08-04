package main

import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Score float32
	Sex string
}

func (s Monster) Print()  {
	fmt.Println("start")
	fmt.Println(s)
	fmt.Println("end")
}
func (s Monster) GetSum(n1,n2 int) int {
	return n1 + n2
}

func (s Monster) Set(name string,age int,score float32,sex string){
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
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
	val.Method(1).Call(nil)
	var params []reflect.Value
	params = append(params,reflect.ValueOf(10))
	params = append(params,reflect.ValueOf(40))
	res := val.Method(0).Call(params)
	fmt.Println(res[0].Int())
}

func main() {
	var a Monster = Monster{
		Name:  "test",
		Age:   2,
		Score: 22,
		Sex:   "nv",
	}
	testStruct(a)
}