## Linux 常用命令

### uname

显示信息

```bash
xiaoyang@xiaoyang ~ % uname -a      //显示全部信息
Darwin xiaoyang 20.0.0 Darwin Kernel Version 20.0.0: Sat Jun 13 17:58:16 PDT 2020; root:xnu-7090.0.0.111.5~1/RELEASE_X86_64 x86_64
xiaoyang@xiaoyang ~ % uname -m			//显示电脑类型
x86_64
xiaoyang@xiaoyang ~ % uname -r			//显示操作系统内核版本
20.0.0
xiaoyang@xiaoyang ~ % uname -s			//显示操作系统名称
Darwin
xiaoyang@xiaoyang ~ % uname -v			//显示操作系统版本
Darwin Kernel Version 20.0.0: Sat Jun 13 17:58:16 PDT 2020; root:xnu-7090.0.0.111.5~1/RELEASE_X86_64
xiaoyang@xiaoyang ~ % uname -n			//显示主机名称
xiaoyang
```

### ls

显示目录和文件

```bash
-a 显示所有文件及目录 (ls内定将文件名或目录名称开头为"."的视为隐藏档，不会列出)
-l 除文件名称外，亦将文件型态、权限、拥有者、文件大小等资讯详细列出
-r 将文件以相反次序显示(原定依英文字母次序)
-t 将文件依建立时间之先后次序列出
-A 同 -a ，但不列出 "." (目前目录) 及 ".." (父目录)
-F 在列出的文件名称后加一符号；例如可执行档则加 "*", 目录则加 "/"
-R 若目录下有文件，则以下之文件亦皆依序列出
```

### cat 

显示文件内容

```bash
cat filename 显示文件内容
cat -n filename 	加上行号显示文件内容
cat -b filename   加上行号显示文件内容。空白行不添加行号
cat -s filename   当遇到有连续两行以上的空白行，就代换为一行的空白行。
cat -e filename   在每行结束后添加$
cat -t filename 	在tab处替换成^I
cat filename1 > filename2 	将文件1的内容cp到文件2  清空文件2的内容
```

### echo

字符串输出

```bash
(base) xiaoyang@xiaoyang test % echo it is a test
it is a test
(base) xiaoyang@xiaoyang test % read name
test
(base) xiaoyang@xiaoyang test % echo name
name
(base) xiaoyang@xiaoyang test % echo \" 
"
(base) xiaoyang@xiaoyang test % echo '\'
\
(base) xiaoyang@xiaoyang test % echo `date`
2020年 7月15日 星期三 10时35分41秒 CST
(base) xiaoyang@xiaoyang test % echo "this is a test" > test.txt 
(base) xiaoyang@xiaoyang test % cat test.txt 
this is a test
```

### date

时间

```bash
(base) xiaoyang@xiaoyang test % date
2020年 7月15日 星期三 10时41分15秒 CST
```

### mkdir

创建目录

```bash
mkdir d1			//创建一个不存在目录
mkdir -p d1		//先检查在创建目录
mkdir -p d2/d3	//创建目录树
```

### rm

删除

```bash
rm filename
rm -r directory
rm -rf filename or directory  
```

### Which

查找在环境变量中的文件

```bash
(base) xiaoyang@xiaoyang test % which cat
/bin/cat
(base) xiaoyang@xiaoyang ~ % which -s brew
/usr/local/bin/brew -> /usr/local/Homebrew/bin/brew
```

### mv

移动和重命名文件

```bash
(base) xiaoyang@xiaoyang d1 % ls
aaa
(base) xiaoyang@xiaoyang d1 % mv aaa bbb
(base) xiaoyang@xiaoyang d1 % ls
bbb
(base) xiaoyang@xiaoyang d1 % mv bbb ./d2 
(base) xiaoyang@xiaoyang d1 % ls
d2
(base) xiaoyang@xiaoyang d1 % ls d2 
bbb
(base) xiaoyang@xiaoyang d1 % mv d2 d3 
(base) xiaoyang@xiaoyang d1 % ls
d3
(base) xiaoyang@xiaoyang d1 % ls d3
d2
```

### cp

拷贝文件

```bash
(base) xiaoyang@xiaoyang d1 % cp /Users/xiaoyang/test.py ./                  
(base) xiaoyang@xiaoyang d1 % ls
d3	test.py
(base) xiaoyang@xiaoyang d1 % cp -r  /Users/xiaoyang/learngit d3 
(base) xiaoyang@xiaoyang d1 % ls d3 
d2		learngit	test
```

![img](https://img-blog.csdn.net/20180823095620695?watermark/2/text/aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L1plcm9fY2hlbmRh/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70)

### ps

显示进程

```bash
ps -a 显示当前用户所有进程
ps -A 显示所有进程
ps -u username 显示用户进程
ps -ef 显示所有命令连带命令行
```

### find

find path -xxx xx

```bash
-mount, -xdev : 只检查和指定目录在同一个文件系统下的文件，避免列出其它文件系统中的文件
-amin n : 在过去 n 分钟内被读取过
-anewer file : 比文件 file 更晚被读取过的文件
-atime n : 在过去n天内被读取过的文件
-cmin n : 在过去 n 分钟内被修改过
-cnewer file :比文件 file 更新的文件
-ctime n : 在过去n天内被修改过的文件
-empty : 空的文件-gid n or -group name : gid 是 n 或是 group 名称是 name
-ipath p, -path p : 路径名称符合 p 的文件，ipath 会忽略大小写
-name name, -iname name : 文件名称符合 name 的文件。iname 会忽略大小写
-size n : 文件大小 是 n 单位，b 代表 512 位元组的区块，c 表示字元数，k 表示 kilo bytes，w 是二个位元组。
-type c : 文件类型是 c 的文件。
```

```bash
(base) xiaoyang@xiaoyang d1 % find . -name "*.py"
./test.py
```

### grep





### whereis

```bash
(base) xiaoyang@xiaoyangdeMBP ~ % whereis ls  
/bin/ls
```

### locate

sudo /usr/libexec/locate.updatedb 进行库更新

```bash
locate filename
```

### more

### 语法

```
more [-dlfpcsu] [-num] [+/pattern] [+linenum] [fileNames..]
```

**参数**：

- -num 一次显示的行数
- -d 提示使用者，在画面下方显示 [Press space to continue, 'q' to quit.] ，如果使用者按错键，则会显示 [Press 'h' for instructions.] 而不是 '哔' 声
- -l 取消遇见特殊字元 ^L（送纸字元）时会暂停的功能
- -f 计算行数时，以实际上的行数，而非自动换行过后的行数（有些单行字数太长的会被扩展为两行或两行以上）
- -p 不以卷动的方式显示每一页，而是先清除萤幕后再显示内容
- -c 跟 -p 相似，不同的是先显示内容再清除其他旧资料
- -s 当遇到有连续两行以上的空白行，就代换为一行的空白行
- -u 不显示下引号 （根据环境变数 TERM 指定的 terminal 而有所不同）
- +/pattern 在每个文档显示前搜寻该字串（pattern），然后从该字串之后开始显示
- +num 从第 num 行开始显示
- fileNames 欲显示内容的文档，可为复数个数

### less

```
less [参数] 文件 
```

**参数说明**：

- -b <缓冲区大小> 设置缓冲区的大小
- -e 当文件显示结束后，自动离开
- -f 强迫打开特殊文件，例如外围设备代号、目录和二进制文件
- -g 只标志最后搜索的关键词
- -i 忽略搜索时的大小写
- -m 显示类似more命令的百分比
- -N 显示每行的行号
- -o <文件名> 将less 输出的内容在指定文件中保存起来
- -Q 不使用警告音
- -s 显示连续空行为一行
- -S 行过长时间将超出部分舍弃
- -x <数字> 将"tab"键显示为规定的数字空格
- /字符串：向下搜索"字符串"的功能
- ?字符串：向上搜索"字符串"的功能
- n：重复前一个搜索（与 / 或 ? 有关）
- N：反向重复前一个搜索（与 / 或 ? 有关）
- b 向上翻一页
- d 向后翻半页
- h 显示帮助界面
- Q 退出less 命令
- u 向前滚动半页
- y 向前滚动一行
- 空格键 滚动一页
- 回车键 滚动一行
- [pagedown]： 向下翻动一页
- [pageup]： 向上翻动一页

### 附加备注

1.全屏导航

- ctrl + F - 向前移动一屏
- ctrl + B - 向后移动一屏
- ctrl + D - 向前移动半屏
- ctrl + U - 向后移动半屏

2.单行导航

- j - 向前移动一行
- k - 向后移动一行

3.其它导航

- G - 移动到最后一行
- g - 移动到第一行
- q / ZZ - 退出 less 命令

4.其它有用的命令

- v - 使用配置的编辑器编辑当前文件
- h - 显示 less 的帮助文档
- &pattern - 仅显示匹配模式的行，而不是整个文件

5.标记导航

当使用 less 查看大文件时，可以在任何一个位置作标记，可以通过命令导航到标有特定标记的文本位置：

- ma - 使用 a 标记文本的当前位置
- 'a - 导航到标记 a 处

### top

实时现实进程的动态

### iostat

输入/输出统计,iostat工具将对系统的磁盘操作活动进行监视

汇报磁盘活动统计情况，同时也会汇报出CPU使用情况

### scp

语法

```
scp [-1246BCpqrv] [-c cipher] [-F ssh_config] [-i identity_file]
[-l limit] [-o ssh_option] [-P port] [-S program]
[[user@]host1:]file1 [...] [[user@]host2:]file2
```

简易写法:

```
scp [可选参数] file_source file_target 
```

**参数说明：**

- -1： 强制scp命令使用协议ssh1
- -2： 强制scp命令使用协议ssh2
- -4： 强制scp命令只使用IPv4寻址
- -6： 强制scp命令只使用IPv6寻址
- -B： 使用批处理模式（传输过程中不询问传输口令或短语）
- -C： 允许压缩。（将-C标志传递给ssh，从而打开压缩功能）
- -p：保留原文件的修改时间，访问时间和访问权限。
- -q： 不显示传输进度条。
- -r： 递归复制整个目录。
- -v：详细方式显示输出。scp和ssh(1)会显示出整个过程的调试信息。这些信息用于调试连接，验证和配置问题。
- -c cipher： 以cipher将数据传输进行加密，这个选项将直接传递给ssh。
- -F ssh_config： 指定一个替代的ssh配置文件，此参数直接传递给ssh。
- -i identity_file： 从指定文件中读取传输时使用的密钥文件，此参数直接传递给ssh。
- -l limit： 限定用户所能使用的带宽，以Kbit/s为单位。
- -o ssh_option： 如果习惯于使用ssh_config(5)中的参数传递方式，
- -P port：注意是大写的P, port是指定数据传输用到的端口号
- -S program： 指定加密传输时所使用的程序。此程序必须能够理解ssh(1)的选项。

### 实例

#### 1、从本地复制到远程 

命令格式：

```
scp local_file remote_username@remote_ip:remote_folder 
或者 
scp local_file remote_username@remote_ip:remote_file 
或者 
scp local_file remote_ip:remote_folder 
或者 
scp local_file remote_ip:remote_file 
```



- 第1,2个指定了用户名，命令执行后需要再输入密码，第1个仅指定了远程的目录，文件名字不变，第2个指定了文件名； 
- 第3,4个没有指定用户名，命令执行后需要输入用户名和密码，第3个仅指定了远程的目录，文件名字不变，第4个指定了文件名；

应用实例：

```
scp /home/space/music/1.mp3 root@www.runoob.com:/home/root/others/music 
scp /home/space/music/1.mp3 root@www.runoob.com:/home/root/others/music/001.mp3 
scp /home/space/music/1.mp3 www.runoob.com:/home/root/others/music 
scp /home/space/music/1.mp3 www.runoob.com:/home/root/others/music/001.mp3 
```

复制目录命令格式： 

```
scp -r local_folder remote_username@remote_ip:remote_folder 
或者 
scp -r local_folder remote_ip:remote_folder 
```

- 第1个指定了用户名，命令执行后需要再输入密码；
- 第2个没有指定用户名，命令执行后需要输入用户名和密码；

应用实例：

```
scp -r /home/space/music/ root@www.runoob.com:/home/root/others/ 
scp -r /home/space/music/ www.runoob.com:/home/root/others/ 
```

上面命令将本地 music 目录复制到远程 others 目录下。

#### 2、从远程复制到本地 

从远程复制到本地，只要将从本地复制到远程的命令的后2个参数调换顺序即可，如下实例 

应用实例：

```
scp root@www.runoob.com:/home/root/others/music /home/space/music/1.mp3 
scp -r www.runoob.com:/home/root/others/ /home/space/music/
```

### 说明

1.如果远程服务器防火墙有为scp命令设置了指定的端口，我们需要使用 -P 参数来设置命令的端口号，命令格式如下：

```
#scp 命令使用端口号 4588
scp -P 4588 remote@www.runoob.com:/usr/local/sin.sh /home/administrator
```

2.使用scp命令要确保使用的用户具有可读取远程服务器相应文件的权限，否则scp命令是无法起作用的。

### sftp

数据传输工具

```bash
sftp> put /Volumes/Data/coding/go/hello.go /
Uploading /Volumes/Data/coding/go/hello.go to /hello.go
/Volumes/Data/coding/go/hello.go              100%   61     0.2KB/s   00:00    
sftp> ls
anaconda-ks.cfg     
sftp> get anaconda-ks.cfg /Volumes/Data
Fetching /root/anaconda-ks.cfg to /Volumes/Data/anaconda-ks.cfg
/root/anaconda-ks.cfg                         100% 1322   643.3KB/s   00:00    
```

### curl

curl发出GET请求不带有任何参数时，curl 就是发出 GET 请求。

> ```bash
> $ curl https://www.example.com
> ```

上面命令向`www.example.com`发出 GET 请求，服务器返回的内容会在命令行输出。

## **-A**

`-A`参数指定客户端的用户代理标头，即`User-Agent`。curl 的默认用户代理字符串是`curl/[version]`。

> ```bash
> $ curl -A 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.100 Safari/537.36' https://google.com
> ```

上面命令将`User-Agent`改成 Chrome 浏览器。

> ```bash
> $ curl -A '' https://google.com
> ```

上面命令会移除`User-Agent`标头。

也可以通过`-H`参数直接指定标头，更改`User-Agent`。

> ```bash
> $ curl -H 'User-Agent: php/1.0' https://google.com
> ```

## **-b*   *

`-b`参数用来向服务器发送 Cookie。

> ```bash
> $ curl -b 'foo=bar' https://google.com
> ```

上面命令会生成一个标头`Cookie: foo=bar`，向服务器发送一个名为`foo`、值为`bar`的 Cookie。

> ```bash
> $ curl -b 'foo1=bar;foo2=bar2' https://google.com
> ```

上面命令发送两个 Cookie。

> ```bash
> $ curl -b cookies.txt https://www.google.com
> ```

上面命令读取本地文件`cookies.txt`，里面是服务器设置的 Cookie（参见`-c`参数），将其发送到服务器。

## **-c**

`-c`参数将服务器设置的 Cookie 写入一个文件。

> ```bash
> $ curl -c cookies.txt https://www.google.com
> ```

上面命令将服务器的 HTTP 回应所设置 Cookie 写入文本文件`cookies.txt`。

## **-d**

`-d`参数用于发送 POST 请求的数据体。

> ```bash
> $ curl -d'login=emma＆password=123'-X POST https://google.com/login
> # 或者
> $ curl -d 'login=emma' -d 'password=123' -X POST  https://google.com/login
> ```

使用`-d`参数以后，HTTP 请求会自动加上标头`Content-Type : application/x-www-form-urlencoded`。并且会自动将请求转为 POST 方法，因此可以省略`-X POST`。

`-d`参数可以读取本地文本文件的数据，向服务器发送。

> ```bash
> $ curl -d '@data.txt' https://google.com/login
> ```

上面命令读取`data.txt`文件的内容，作为数据体向服务器发送。

## **--data-urlencode**

`--data-urlencode`参数等同于`-d`，发送 POST 请求的数据体，区别在于会自动将发送的数据进行 URL 编码。

> ```bash
> $ curl --data-urlencode 'comment=hello world' https://google.com/login
> ```

上面代码中，发送的数据`hello world`之间有一个空格，需要进行 URL 编码。

## **-e**

`-e`参数用来设置 HTTP 的标头`Referer`，表示请求的来源。

> ```bash
> curl -e 'https://google.com?q=example' https://www.example.com
> ```

上面命令将`Referer`标头设为`https://google.com?q=example`。

`-H`参数可以通过直接添加标头`Referer`，达到同样效果。

> ```bash
> curl -H 'Referer: https://google.com?q=example' https://www.example.com
> ```

## **-F**

`-F`参数用来向服务器上传二进制文件。

> ```bash
> $ curl -F 'file=@photo.png' https://google.com/profile
> ```

上面命令会给 HTTP 请求加上标头`Content-Type: multipart/form-data`，然后将文件`photo.png`作为`file`字段上传。

`-F`参数可以指定 MIME 类型。

> ```bash
> $ curl -F 'file=@photo.png;type=image/png' https://google.com/profile
> ```

上面命令指定 MIME 类型为`image/png`，否则 curl 会把 MIME 类型设为`application/octet-stream`。

`-F`参数也可以指定文件名。

> ```bash
> $ curl -F 'file=@photo.png;filename=me.png' https://google.com/profile
> ```

上面命令中，原始文件名为`photo.png`，但是服务器接收到的文件名为`me.png`。

## **-G**

`-G`参数用来构造 URL 的查询字符串。

> ```bash
> $ curl -G -d 'q=kitties' -d 'count=20' https://google.com/search
> ```

上面命令会发出一个 GET 请求，实际请求的 URL 为`https://google.com/search?q=kitties&count=20`。如果省略`--G`，会发出一个 POST 请求。

如果数据需要 URL 编码，可以结合`--data--urlencode`参数。

> ```bash
> $ curl -G --data-urlencode 'comment=hello world' https://www.example.com
> ```

## **-H**

`-H`参数添加 HTTP 请求的标头。

> ```bash
> $ curl -H 'Accept-Language: en-US' https://google.com
> ```

上面命令添加 HTTP 标头`Accept-Language: en-US`。

> ```bash
> $ curl -H 'Accept-Language: en-US' -H 'Secret-Message: xyzzy' https://google.com
> ```

上面命令添加两个 HTTP 标头。

> ```bash
> $ curl -d '{"login": "emma", "pass": "123"}' -H 'Content-Type: application/json' https://google.com/login
> ```

上面命令添加 HTTP 请求的标头是`Content-Type: application/json`，然后用`-d`参数发送 JSON 数据。

## **-i**

`-i`参数打印出服务器回应的 HTTP 标头。

> ```bash
> $ curl -i https://www.example.com
> ```

上面命令收到服务器回应后，先输出服务器回应的标头，然后空一行，再输出网页的源码。

## **-I**

`-I`参数向服务器发出 HEAD 请求，然会将服务器返回的 HTTP 标头打印出来。

> ```bash
> $ curl -I https://www.example.com
> ```

上面命令输出服务器对 HEAD 请求的回应。

`--head`参数等同于`-I`。

> ```bash
> $ curl --head https://www.example.com
> ```

## **-k**

`-k`参数指定跳过 SSL 检测。

> ```bash
> $ curl -k https://www.example.com
> ```

上面命令不会检查服务器的 SSL 证书是否正确。

## **-L**

`-L`参数会让 HTTP 请求跟随服务器的重定向。curl 默认不跟随重定向。

> ```bash
> $ curl -L -d 'tweet=hi' https://api.twitter.com/tweet
> ```

## **--limit-rate**

`--limit-rate`用来限制 HTTP 请求和回应的带宽，模拟慢网速的环境。

> ```bash
> $ curl --limit-rate 200k https://google.com
> ```

上面命令将带宽限制在每秒 200K 字节。

## **-o**

`-o`参数将服务器的回应保存成文件，等同于`wget`命令。

> ```bash
> $ curl -o example.html https://www.example.com
> ```

上面命令将`www.example.com`保存成`example.html`。

## **-O**

`-O`参数将服务器回应保存成文件，并将 URL 的最后部分当作文件名。

> ```bash
> $ curl -O https://www.example.com/foo/bar.html
> ```

上面命令将服务器回应保存成文件，文件名为`bar.html`。

## **-s**

`-s`参数将不输出错误和进度信息。

> ```bash
> $ curl -s https://www.example.com
> ```

上面命令一旦发生错误，不会显示错误信息。不发生错误的话，会正常显示运行结果。

如果想让 curl 不产生任何输出，可以使用下面的命令。

> ```bash
> $ curl -s -o /dev/null https://google.com
> ```

## **-S**

`-S`参数指定只输出错误信息，通常与`-s`一起使用。

> ```bash
> $ curl -s -o /dev/null https://google.com
> ```

上面命令没有任何输出，除非发生错误。

## **-u**

`-u`参数用来设置服务器认证的用户名和密码。

> ```bash
> $ curl -u 'bob:12345' https://google.com/login
> ```

上面命令设置用户名为`bob`，密码为`12345`，然后将其转为 HTTP 标头`Authorization: Basic Ym9iOjEyMzQ1`。

curl 能够识别 URL 里面的用户名和密码。

> ```bash
> $ curl https://bob:12345@google.com/login
> ```

上面命令能够识别 URL 里面的用户名和密码，将其转为上个例子里面的 HTTP 标头。

> ```bash
> $ curl -u 'bob' https://google.com/login
> ```

上面命令只设置了用户名，执行后，curl 会提示用户输入密码。

## **-v**

`-v`参数输出通信的整个过程，用于调试。

> ```bash
> $ curl -v https://www.example.com
> ```

`--trace`参数也可以用于调试，还会输出原始的二进制数据。

> ```bash
> $ curl --trace - https://www.example.com
> ```

## **-x**

`-x`参数指定 HTTP 请求的代理。

> ```bash
> $ curl -x socks5://james:cats@myproxy.com:8080 https://www.example.com
> ```

上面命令指定 HTTP 请求通过`myproxy.com:8080`的 socks5 代理发出。

如果没有指定代理协议，默认为 HTTP。

> ```bash
> $ curl -x james:cats@myproxy.com:8080 https://www.example.com
> ```

上面命令中，请求的代理使用 HTTP 协议。

## **-X**

`-X`参数指定 HTTP 请求的方法。

> ```bash
> $ curl -X POST https://www.example.com
> ```

上面命令对`https://www.example.com`发出 POST 请求。

### lsof

```bash
lsof (选项)
```

选项	描述
-a							列出打开文件存在的进程；
-c <进程名>.  		列出指定进程所打开的文件；
-g							列出GID号进程详情；
-d  <文件号>	 	列出占用该文件号的进程；
+d <目录>			 列出目录下被打开的文件；
+D <目录>			递归列出目录下被打开的文件；
-n <目录>		 	列出使用NFS的文件；
-i <条件>			  列出符合条件的进程。（4、6、协议、:端口、 @ip ）
-p<进程号>		  列出指定进程号所打开的文件；
-u						  列出UID号进程详情；
-h						  显示帮助信息；
-v						   显示版本信息。
lsof输出各列信息的意义如下

COMMAND：进程的名称
PID：进程标识符
PPID：父进程标识符（需要指定-R参数）
USER：进程所有者
PGID：进程所属组
FD：文件描述符，应用程序通过文件描述符识别该文件。
DEVICE：指定磁盘的名称
SIZE：文件的大小
NODE：索引节点（文件在磁盘上的标识）
NAME：打开文件的确切名称
FD文件描述符列表

cwd：表示current work dirctory，即：应用程序的当前工作目录，这是该应用程序启动的目录，除非它本身对这个目录进行更改
txt：该类型的文件是程序代码，如应用程序二进制文件本身或共享库，如上列表中显示的 /sbin/init 程序
lnn：library references (AIX)（库引用）;
er：FD information error (see NAME column)（fd信息错误）;
jld：jail directory (FreeBSD)（监控目录）;
ltx：shared library text (code and data)（共享库文本）;
mxx ：hex memory-mapped type number xx（十六进制内存映射类型号xx）；
m86：DOS Merge mapped file(DOS合并映射文件);
mem：memory-mapped file(内存映射文件);
mmap：memory-mapped device（内存映射设备）;
pd：parent directory（父目录）;
rtd：root directory（跟目录）;
tr：kernel trace file (OpenBSD)（内核跟踪文件）;
v86 VP/ix mapped file（VP/IX映射文件）;
0：表示标准输出
1：表示标准输入
2：表示标准错误
一般在标准输出、标准错误、标准输入后还跟着文件状态模式：

u：表示该文件被打开并处于读取/写入模式。
r：表示该文件被打开并处于只读模式。
w：表示该文件被打开并处于。
空格：表示该文件的状态模式为unknow，且没有锁定。
-：表示该文件的状态模式为unknow，且被锁定。
同时在文件状态模式后面，还跟着相关的锁：

N：for a Solaris NFS lock of unknown type(对于未知类型的Solaris NFS锁);
r：for read lock on part of the file(用于对文件的一部分进行读取锁定);
R：for a read lock on the entire file(整个文件的读取锁定);
w：for a write lock on part of the file;（文件的部分写锁）
W：for a write lock on the entire file;（整个文件的写锁）
u：for a read and write lock of any length(对于任意长度的读写锁);
U：for a lock of unknown type(对于未知类型的锁);
x：for an SCO OpenServer Xenix lock on part of the file(对于文件的sco openserver xenix锁);
X：for an SCO OpenServer Xenix lock on the entire file(对于整个文件的sco openserver xenix锁);
space：if there is no lock(如果没有锁).
2.1.2 文件类型：

DIR：表示目录。
CHR：表示字符类型。
BLK：块设备类型。
UNIX： UNIX 域套接字。
FIFO：先进先出 (FIFO) 队列。
IPv4：网际协议 (IP) 套接字。
DEVICE：指定磁盘的名称
SIZE：文件的大小
NODE：索引节点（文件在磁盘上的标识）
NAME：打开文件的确切名称
2.2 常用方式

```bash
# 查看谁正在使用某个文件

lsof   /filepath/file

#递归查看某个目录的文件信息
lsof +D /filepath/filepath2/
备注: 使用了+D，对应目录下的所有子目录和文件都会被列出

# 比使用+D选项，遍历查看某个目录的所有文件信息 的方法

lsof | grep ‘/filepath/filepath2/’

# 列出某个用户打开的文件信息

lsof  -u username
备注: -u 选项，u其实是user的缩写

# 列出某个程序所打开的文件信息

lsof -c mysql
备注: -c 选项将会列出所有以mysql开头的程序的文件，其实你也可以写成lsof | grep mysql,但是第一种方法明显比第二种方法要少打几个字符了

# 列出多个程序多打开的文件信息

lsof -c mysql -c apache

# 列出某个用户以及某个程序所打开的文件信息

lsof -u test -c mysql

# 列出除了某个用户外的被打开的文件信息

lsof   -u ^root
备注：^这个符号在用户名之前，将会把是root用户打开的进程不让显示

# 通过某个进程号显示该进行打开的文件

lsof -p 1

# 列出多个进程号对应的文件信息

lsof -p 123,456,789

# 列出除了某个进程号，其他进程号所打开的文件信息

lsof -p ^1

# 列出所有的网络连接

lsof -i

# 列出所有tcp 网络连接信息

lsof  -i tcp

# 列出所有udp网络连接信息

lsof  -i udp

# 列出谁在使用某个端口

lsof -i :3306

# 列出谁在使用某个特定的udp端口

lsof -i udp:55

# 特定的tcp端口

lsof -i tcp:80

# 列出某个用户的所有活跃的网络端口

lsof  -a -u test -i

# 列出所有网络文件系统

lsof -N

#域名socket文件
lsof -u

#某个用户组所打开的文件信息
lsof -g 5555

# 根据文件描述列出对应的文件信息

lsof -d description(like 2)

# 根据文件描述范围列出文件信息

lsof -d 2-3

```

### file

查看是什么文件

```bash
(base) xiaoyang@xiaoyangdeMBP go % file hello.go 
hello.go: c program text, ASCII text
(base) xiaoyang@xiaoyangdeMBP go % file -b hello.go 
c program text, ASCII text
(base) xiaoyang@xiaoyangdeMBP go % file -i hello.go
hello.go: regular file
(base) xiaoyang@xiaoyangdeMBP go % file -bi hello.go
regular file
```

### wget


