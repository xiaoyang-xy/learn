package main

import (
	"flag"
	"fmt"
)

func main() {
	var user string
	var pwd string
	var host string
	var port int

	flag.StringVar(&user,"u","","")
	flag.StringVar(&pwd,"pwd","","")
	flag.StringVar(&host,"h","localhost","")
	flag.IntVar(&port,"port",3306,"")
	flag.Parse()
	fmt.Println(user,pwd,host,port)
}
