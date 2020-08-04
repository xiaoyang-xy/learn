package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func process(conn net.Conn)  {
	defer conn.Close()
	for true {
		buf := make([]byte,1024)
		n, _ := conn.Read(buf)
		if n == 0 {
			return
		}
		fmt.Print(string(buf[:n]))
	}
}

func main() {
	fmt.Println("listen")
	listen, err := net.Listen("tcp","0.0.0.0:8888")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listen.Close()
	for true {
		con, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(con,con.RemoteAddr(),con.LocalAddr())
		go process(con)
		reader := bufio.NewReader(os.Stdin)
		str, err := reader.ReadString('\n')
		str = strings.Trim(str," \r\n")
		if str =="exit"{
			return
		}
	}
}
