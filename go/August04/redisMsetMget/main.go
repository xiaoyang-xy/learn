package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

func main() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()
	_, err = c.Do("hmset","user02","name","john","age",19)
	if err != nil {
		fmt.Println(err)
	}
	str, err := redis.Strings(c.Do("hmget","user02","name","age"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(str)
	for i, i2 := range str {
		fmt.Println(i,i2)
	}
	_, err = c.Do("Set", "name", "tom")
	time.Sleep(time.Second*11)
	c.Do("expire","name",10)
	str1, err := redis.String(c.Do("Get","name"))

	fmt.Println(str1)
}
