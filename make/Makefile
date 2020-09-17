a.txt:b.txt c.txt
	cat b.txt c.txt > a.txt
test:
	echo "test"
build:
	echo "build"
#当目录下有clean文件时，需要加入.PHONY来表示clean这个命令强制执行
.PHONY:clean
clean:
	rm *.o *.s
var-kept:
	@export foo=bar;\
	echo "foo=[$$foo]"
