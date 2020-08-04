package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err :=net.Dial("tcp","10.0.15.100:8888")
	if err != nil {
		fmt.Println(err)
		return
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(conn)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)
	n,err := conn.Write([]byte(str))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)

}
