## 学习笔记-Go

### hello,world

```go
package main
import "fmt"
func main()
{
	fmt.Println("hello,world")
}
```

### 变量

```go
var size int		//声明变量
var num = 0			//声明变量并初始化
num :=0					//声明一个局部变量并初始化
```

### 指针

```bash
	var x int
	var p *int;
	p = &x
	x=1;
	var p = new(int)  //return *int
```

