package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name string `json:"name"`
	Age int	`json:"age"`
	Birthday string
	Sal float64
	Skill string
}

func testStruct(){
	var monster Monster = Monster{
		Name:     "yang",
		Age:      12,
		Birthday: "2012/01/01",
		Sal:      1770.2,
		Skill:    "fuck",
	}
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}

func testmap()  {
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "yang"
	a["age"] = 12
	a["address"] = "beijing"
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}

func testslice()  {
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"] = "111"
	m1["age"] = 13
	slice = append(slice,m1)
	var m2 map[string]interface{}
	m2 = make(map[string]interface{})
	m2["name"] = "111"
	m2["age"] = 13
	slice = append(slice,m2)
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))

}

func testNum()  {
	var num1 float64 = 1.123
	data,err := json.Marshal(num1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}
func main() {
	testStruct()
	testmap()
	testslice()
	testNum()
}
