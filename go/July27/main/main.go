package main

import (
	"bufio"
	"container/list"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main()  {
	var a list.List
	a.PushBack("123")
	b := a.Front()
	p := &b.Value
	*p = 10
	fmt.Printf(len())
	for  {
		
	}
	var str111 string
	fmt.Println(len(str111))
	str111 += "123"
	fmt.Print(a.Front().Value)
	//os.Open return *File, error
	file, err := os.Open("/Volumes/Data/coding/test.txt")
	if err != nil {
		fmt.Println("open file err=",err)
	}
	//一次性将文件输入内存，不带缓冲	适用于小文件
	//return []byte ,error
	str, err := ioutil.ReadFile("/Volumes/Data/coding/test.txt")
	if err != nil {
		fmt.Printf("read file err=%v",err)
	}
	//无需关闭file	[]byte->string
	fmt.Printf("%v",string(str))
	//return *Reader	带缓冲的方式
	reader := bufio.NewReader(file)
	for  {
		//return string,error
		str, err := reader.ReadString('\n')
		fmt.Print(str)
		if err == io.EOF {
			break
		}

	}
	fmt.Printf("file=%v",file)
	defer file.Close()
	//file.close   return error
	//err = file.Close()
	//if err != nil {
	//	fmt.Println("close file err=",err)
	//}
}