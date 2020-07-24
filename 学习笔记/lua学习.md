## lua学习

### hello world

```lua
print("Hello World!")
```

### 未初始化和初始化变量

变量未初始化为nil，如果你想删除一个全局变量，b = nil即可

```lua
a 
print(a)
b = 10
print(b)
```

```bash
nil
10
```

### 注释

```lua
-- 单行注释
--[[
	多行注释
]]--
```

### lua数据类型

| 数据类型 | 描述                                                         |
| :------- | ------------------------------------------------------------ |
| nil      | 无效值，未初始化或者未定义变量type(a) == nil                 |
| boolean  | false和true                                                  |
| number   | 数字，lua将整型和浮点型放在了一次                            |
| string   | 字符串，‘123’ or "123" or [[123]]                            |
| function | 由C或者lua编写的函数                                         |
| userdata | 数据结构                                                     |
| thread   | 协程 类似于线程                                              |
| table    | 表，相当于关联数组，[key,value]  key默认从1开始如果不初始化。 |

```lua
print(type("111"),type(123),type(print),type(true),type(c),type(type(123)))
string	number	function	boolean	nil	string
```

### Lua变量

变量类型： 全局变量、局部变量、表中的域

```lua
a = 5	--全局变量  作用域程序开始至结束
local = 6 --局部变量  作用域随着函数或者循环的结束而消失
```

### Lua循环

#### while循环

```lua
a = 10
while(a < 20)
do
   print(a)
   a = a + 1
end
```

