package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	//return *file,error
	file,err := os.OpenFile("/Volumes/Data/coding/test.txt",os.O_RDWR,0666)
	if err != nil {
		fmt.Printf("open file err=%v\n",err)
		return
	}
	defer  file.Close()
	str :="Hello world"
	//return *writer
	writer := bufio.NewWriter(file)
	//向file添加str,此方法写入缓存
	writer.WriteString(str)
	//将缓存写入file，并清楚缓存
	writer.Flush()
}
