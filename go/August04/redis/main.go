package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()
	_, err = c.Do("Set", "key1", 998)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = c.Do("Set", "name", "tom")
	str, err := redis.String(c.Do("Get","name"))
	fmt.Println(str)
	r, err := redis.Int(c.Do("Get", "key1"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(r)

}
