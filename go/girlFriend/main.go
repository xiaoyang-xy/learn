package main

import "fmt"

type girlFriend struct {
	Name string `json:"name"`
	Age uint8	`json:"age"`
	Sex string	`json:"sex"`
	BoyFriendName string	`json:"boy_friend_name"`
	Length float32	`json:"length"`
}

func main() {
	xiaobaoGirlFriend := girlFriend{
		Name:          "object",
		Age:           5,
		Sex:           "男",
		BoyFriendName: "郭小宝",
		Length:        70.63,
	}
	fmt.Println(xiaobaoGirlFriend)
}
