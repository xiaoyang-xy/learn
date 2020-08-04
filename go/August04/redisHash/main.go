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
	_, err = c.Do("hset","user01", "name", "tom")
	if err != nil {
		fmt.Println(err)
		return
	}
	str, err := redis.String(c.Do("hget","user01","name"))
	fmt.Println(str)
	if err != nil {
		fmt.Println(err)
		return
	}

}
