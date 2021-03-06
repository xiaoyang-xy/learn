package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type CharCount struct {
	chcount int
	numcount int
	spacecount int
	othercount int
}
func main() {
	file,err := os.Open("/Volumes/Data/coding/test.txt")
	if err != nil {
		fmt.Println("open file error=",err)
		return
	}
	var s = "123"
	fmt.Println(s[1])
	var count CharCount
	defer file.Close()
	reader := bufio.NewReader(file)
	for i:=0; i< 10 ;i++ {
		str,err := reader.ReadString('\n')
		if  != nil {

		}
		if err == io.EOF {
			break
		}
		for _, v := range str {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				count.chcount++
			case v >= '0' && v <= '9':
				count.numcount++
			case v == ' ' || v == '\t':
				count.spacecount++
			default:
				count.othercount++
			}
		}
	}
	fmt.Println(count)
}