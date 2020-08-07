### 1.redis配置

config get * 				获取所有配置信息

config set name new value   修改名字为name的配置信息

### 2.redis键(keys)命令

del key    				删除键

dump key 				序列化给定key，并返回被序列化的值

exists key				检查key是否存在

expire key time			设置key的过期时间

expireat key timestamp		设置key的过期时间(使用unix时间戳）

pexpire key millseconds	设置key的过期时间已毫秒计

pexpireat key mill-timestamp 设置key的过期时间以毫秒计(使用unix时间戳)

keys pattern				查找给定模式pattern的key    keys runno* 。  keys  *	

move key db    move key 1 转移到数据库1。默认0

persist key		移除过期时间

pttl key				显示key过期时间。-1为0	以毫秒为单位

Til key 				显示key过期时间 以秒为单位

randomkey 		从数据库中随机返回一个key

rename key newkey		重命名

renamenx key newkey		当newkey不存在时改名

Type key 返回key的类型

Scan cursor. [pattern] [count] 游标。匹配模式		返回多少元素

### 3.redis字符串(String)

set key  key_value			设置指定key的值

get key					获取指定key的值

getrange key start end  	返回key字符串值的子字符串

getset key value			将给定key的值设置为value 并获取旧值

getbit key offset			对 key 所储存的字符串值，获取指定偏移量上的位(bit)。

mget key1[key2..]			   获取一个或多个key值

setbit key offset 			  getbit  bit 10086 1  对key所存储的字符串值，设置或者清除指定偏移量的位

setex key seconds value 		将值value关联到key，并将key的过期时间设置为seconds

setnx key value			只有在key不存在时设置key的值

setrange key offset value 	用value参数覆写给定key所存储的字符串值，从偏移量offset开始

strlen	key				返回字符串长度

mset	 key value[key value]	同时设置多个key-value对

msetnx key value[key value]       同时设置多个 。 在他们不存在时

psetex key milliseconds value	类似于setex 。  但是时间为毫秒

incr key						将key中存储的数字值加一

incrby key increment			将key所存储的值加上给定增量

incrbyfloat	increment		将key中储存的值加上给定浮点增量值

decr	key					减一

decrby	key increment		j减多少

append key value 				字符串添加 。   key +value字符串连接

### 4.hash

hmset key field value	创建一个hash

hdel key field1 field2		删除hash一个子值

hgetall key 		获取key的所有值

hexists	key	field	查看key的指定字段是否存在

hget	key field 		获取key中指定字段的值

Hincrby key field n	指定字段增量n

hincrbyfloat key field n	指定字段增量float n

hkeys key		key中的key

Hlen key		key的字段数量

hset key field value

hsetnx key field value 	只在字段不存在时 设置哈希表字段的值

hvals key 	获取所有的value

hscan key cursor [pattern] [count]. 迭代哈希表中的key-value对

### 5.list

lpush	key	field		双向队列。  默认先进先出

Bloop key timeout 	移除 并获取列表的第一个元素 如果列表没有元素会阻塞列表知道有或者超时

Brpop key timeout 	移除并获取列表的最后一个元素	后等上

lindex key index		通过索引获取元素

llen key 	获取长度

lpop	key		移除并获取列表的第一个元素

lpush  key	在列表头部添加元素

lpushx key	在列表头部添加元素。列表里必须要有元素

linsert key before|after "元素1" “元素2” 	在元素1前后添加元素2

lrange key start stop 	获取 [strat:stop]元素

lrem key count value 	遍历全部。删除元素为value的count次数

lset key index value	通过索引设置列表元素的值

ltrim key start stop 	保留[start:stop]中的元素

rpop key 移除列表最后一个元素返回该元素

rpush key 往右边添加元素

rpushx key	往最右边添加元素，列表要存在

rpoplpush key1 key2	移除最后一个元素添加到另一个列表

### 6.set

set。无序。唯一

sadd key mem1 mem2 mem3. 向集合里添加数据

scard key 返回集合成员数

sdiff key1 key2	返回key1 和key2的差异

sdiffstore key key1 key2 将key1 和ke y2的差异放在key中

sinter key1 key2 返回key1和key2的交集

sinterstore key key1 key2 将交集返回key

sismember key member 判断member是否是key中的成员

smembers key 返回集合中所有成员

smove key1 key2 member 将member从key1移动到key2

spop key 移除并返回集合中的一个随机元素

srandmember key [count]	返回集合中一个或多个随机数

srem key member1 member2 删除集合中一个或者多个成员

sunion key1 key2 返回给定集合并集

sunionstore key key1 key2 将并集放入key

sscan key cursor [pattern] [count] 迭代集合元素

### 7.有序集合

zadd key num member. 添加

相对集合多了个参数来表示位置

### 8.基数统计

pfadd key field	添加

pfcount key 返回key的数量

Pfmerge key key1 key2 将多个key1和key2合并为key'

### 9.发布订阅

Subscribe channel 订阅给定频道信息

Unsubscribe channel 退订频道

psubscribe pattern 	订阅给定匹配模式的频道

unpsbscribe pattern 退订匹配的频道

pubsub subcommand	查看订阅和发布系统状态

publish channel message 将信息发送到指定频道

### 10.事务

multi 标记事务块的开始

exec 执行事务块

discard 取消事务

watch key1 key2 监视key1活着

unwatch 取消监视

### 11.脚本

Eval script numkeys key [key ...] arg [arg ...]

运行lua脚本。

```lua
eval "return {KEYS[1],KEYS[2],ARGV[1],ARGV[2]}" 2 key3 key2 first second
1) "key3"
2) "key2"
3) "first"
4) "second"
```

evalsha sha1 numkeys key [key...] arg [arg ...]

sha1 通过script load生成的sha1校验码

Numkeys 用于指定键名参数的个数

```lua
SCRIPT LOAD "return 'hello moto'"
"232fd51614574cf0867b83d384a5e898cfd24e5a"

redis 127.0.0.1:6379> EVALSHA "232fd51614574cf0867b83d384a5e898cfd24e5a" 0
```

script flush

删除脚本缓存

script kill 杀死当前运行的lua脚本

script load 将脚本添加到缓存

script exists script 查看置顶脚本是否被保存

```lua
script exists sha1
```

### 12.Redis连接

Auth password 

检测给定的密码活着配置文件的密码是否相符

```
127.0.0.1:6379> CONFIG SET requirepass "mypass"
OK
127.0.0.1:6379> auth mypass
OK
```

Echo message 

打印字符串

Ping -> pong 查看服务是否运行

quit 退出当前连接

select index 	切换指定数据库

13.redis服务器命令

info 输出服务器信息

Bgrewriteaof   	手动触发重写操作 异步执行一个aof文件重写操作

bgsave 将后台异步保存到当前数据库的数据到磁盘

BGSAVE 命令执行之后立即返回 OK ，然后 Redis fork 出一个新子进程，原来的 Redis 进程(父进程)继续处理客户端请求，而子进程则负责将数据保存到磁盘，然后退出。

Client  kill ip:port		关闭一个redis-cli连接

client list 返回所有连接到服务器的客户端信息和统计数据

以下是域的含义：

- addr ： 客户端的地址和端口
- fd ： 套接字所使用的文件描述符
- age ： 以秒计算的已连接时长
- idle ： 以秒计算的空闲时长
- flags ： 客户端 flag
- db ： 该客户端正在使用的数据库 ID
- sub ： 已订阅频道的数量
- psub ： 已订阅模式的数量
- multi ： 在事务中被执行的命令数量
- qbuf ： 查询缓冲区的长度（字节为单位， 0 表示没有分配查询缓冲区）
- qbuf-free ： 查询缓冲区剩余空间的长度（字节为单位， 0 表示没有剩余空间）
- obl ： 输出缓冲区的长度（字节为单位， 0 表示没有分配输出缓冲区）
- oll ： 输出列表包含的对象数量（当输出缓冲区没有剩余空间时，命令回复会以字符串对象的形式被入队到这个队列里）
- omem ： 输出缓冲区和输出列表占用的内存总量
- events ： 文件描述符事件
- cmd ： 最近一次执行的命令

客户端 flag 可以由以下部分组成：

- O ： 客户端是 MONITOR 模式下的附属节点（slave）

- S ： 客户端是一般模式下（normal）的附属节点

- M ： 客户端是主节点（master）

- x ： 客户端正在执行事务

- b ： 客户端正在等待阻塞事件

- i ： 客户端正在等待 VM I/O 操作（已废弃）

- d ： 一个受监视（watched）的键已被修改， EXEC 命令将失败

- c : 在将回复完整地写出之后，关闭链接

- u : 客户端未被阻塞（unblocked）

- A : 尽可能快地关闭连接

- N : 未设置任何 flag

  client getname

  client setname  “name”

  client pause timeout 	阻塞客户端 timeout时间 毫秒

  Cluster slots	查看当前集群的状态 以数组形式展示

  command 返回redis命令的详细信息

  Command count 返回redis命令数量

  command getkeys mset a b c d e f	返回所有key

  command info command-name 查看命令描述信息

  config get parameter 用户获取服务的配置参数

  config rewrite parameer 对配置文件进行改写

  config set parameter value 对配置文件动态调整

  config resetstat 重置info命令的某些统计数据

  dbsize。返回当前数据库的key的数量

  Debug object key. 调试

  Debug segfault  执行一个非法的内存访问让redis崩溃 仅在开发时用于bug调试

  Flushall 删除服务器的数据

  flushdb	清空当前数据库的所有key

  lastsave 	返回上一次将数据保存在磁盘上的时间

  实时打印redis服务器接收到的命令

  role 查看主从实例所属的角色 角色有master slave sentinel

  将redis实例数据快照以rdb文件的形式保存在硬盘

  Redis Shutdown 命令执行以下操作：

  - 停止所有客户端

  - 如果有至少一个保存点在等待，执行 SAVE 命令

  - 如果 AOF 选项被打开，更新 AOF 文件

  - 关闭 redis 服务器(server)

    slaveof将当前服务器转变为指定服务器的丛书服务器

    showlog command记录查询执行时间的日志系统

    Sync  同步主从服务器

    ### Redis 数据备份和恢复

    save。bgsave

    创建当前数据库的备份

    如果需要恢复数据，只需将备份文件 (dump.rdb) 移动到 redis 安装目录并启动服务即可。获取 redis 目录可以使用 **CONFIG** 命令

    config get dir

    time 返回当前服务器时间

### Redis 默认无密码

config get requirepass		获得密码

config set requirepass "password" 	设置密码

auth password 	登陆

### Redis 性能测试

Redis-benchmark [option] [option value]

redis 性能测试工具可选参数如下所示：

| 序号 | 选项      | 描述                                       | 默认值    |
| :--- | :-------- | :----------------------------------------- | :-------- |
| 1    | **-h**    | 指定服务器主机名                           | 127.0.0.1 |
| 2    | **-p**    | 指定服务器端口                             | 6379      |
| 3    | **-s**    | 指定服务器 socket                          |           |
| 4    | **-c**    | 指定并发连接数                             | 50        |
| 5    | **-n**    | 指定请求数                                 | 10000     |
| 6    | **-d**    | 以字节的形式指定 SET/GET 值的数据大小      | 2         |
| 7    | **-k**    | 1=keep alive 0=reconnect                   | 1         |
| 8    | **-r**    | SET/GET/INCR 使用随机 key, SADD 使用随机值 |           |
| 9    | **-P**    | 通过管道传输 <numreq> 请求                 | 1         |
| 10   | **-q**    | 强制退出 redis。仅显示 query/sec 值        |           |
| 11   | **--csv** | 以 CSV 格式输出                            |           |
| 12   | **-l**    | 生成循环，永久执行测试                     |           |
| 13   | **-t**    | 仅运行以逗号分隔的测试命令列表。           |           |
| 14   | **-I**    | Idle 模式。仅打开 N 个 idle 连接并等待。   |           |

### redis 客户端连接

client list	返回连接到redis服务的客户端列表

client setname	name	设置当前连接name

client pause timeout	阻塞当前连接

client kill  ip:port关闭客户端连接