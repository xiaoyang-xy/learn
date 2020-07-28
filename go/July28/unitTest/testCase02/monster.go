package testCase02

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

)

type Monster struct {
	Name string
	Age int
	Skill string
}

func (this *Monster) Store() bool{

	data,err := json.Marshal(this)
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println(string(data))

	filepath := "/Volumes/Data/coding/test.json"
	err = ioutil.WriteFile(filepath,data,0777)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func (this *Monster)ReStore()  bool{
	filepath := "/Volumes/Data/coding/test.json"
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
		return false
	}
	err = json.Unmarshal([]byte(data),this)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true


	}

