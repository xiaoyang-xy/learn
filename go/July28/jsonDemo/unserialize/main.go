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
func reserialize() {
	var monster Monster
	str := "{\"name\":\"yang\",\"age\":12,\"Birthday\":\"2012/01/01\",\"Sal\":1770.2,\"Skill\":\"fuck\"}"
	err := json.Unmarshal([]byte(str),&monster)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(monster)
}
func remap(){
	var a map[string]interface{}
	str := "{\"name\":\"yang\",\"age\":12,\"Birthday\":\"2012/01/01\",\"Sal\":1770.2,\"Skill\":\"fuck\"}"
	err := json.Unmarshal([]byte(str),&a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)
}
func reslice()  {
	var slice []map[string]interface{}
	str := "[{\"age\":13,\"name\":\"111\"},{\"age\":13,\"name\":\"111\"}]\n"
	err := json.Unmarshal([]byte(str),&slice)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(slice)
}
func main()  {
	reserialize()
	remap()
	reslice()
}
