package main

import (
	"../common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

func readPkg(conn net.Conn) (mes message.Message,err error) {
		buf := make([]byte, 8096)
		n, err := conn.Read(buf[:4])
		if n == 0 {
			fmt.Println("read error = ", err)
			return
		}
		fmt.Println(buf[:4])
		var pkgLen uint32
		pkgLen = binary.BigEndian.Uint32(buf[0:4])

		n, err = conn.Read(buf[:pkgLen])
		if n == 0 || err != nil {
			fmt.Println(err)
			return
		}
	err = json.Unmarshal(buf[:pkgLen],&mes)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
func process(conn net.Conn)  {
	//读客户端发送的信息
	defer conn.Close()



}
func main() {
	fmt.Println("服务器在8889端口监听....")
	listen, err := net.Listen("tcp","0.0.0.0:8889")
	if err != nil {
		fmt.Println("listen err=",err)
		return
	}
	for true {
		fmt.Println("wait client connet.....")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("connet error=",err)
		}
		go process(conn)
	}
}