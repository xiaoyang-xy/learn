package main
//文件从一个文件传到另一个文件
import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func PathExists(path string)(bool,error){
	_,err := os.Stat(path)
	//文件夹或者文件存在
	if err == nil {
		return true,nil
	}
	//文件夹或者文件不存在
	if os.IsExist(err){
		return false,nil
	}
	//别的错误，原路返回
	return false,err
}
func main() {
	file1path := "/Volumes/Data/coding/test.txt"
	file2path := "/Volumes/Data/coding/read.txt"
	content, err := ioutil.ReadFile(file1path)
	file1, err := os.Open(file1path)
	file2, err := os.OpenFile(file2path,os.O_WRONLY | os.O_CREATE,0666)
	reader := bufio.NewReader(file1)
	writer := bufio.NewWriter(file2)
	defer file1.Close()
	_, err = io.Copy(writer, reader)

	if err != nil {
		fmt.Printf("open file1 err=%v\n",err)
		return
	}
	err = ioutil.WriteFile(file2path,content,0777)
	if err != nil {
		fmt.Printf("write file2 error=%v",err)
	}
}
