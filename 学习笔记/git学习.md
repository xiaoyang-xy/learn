# git学习

## 配置用户信息

GitHub上的username和email

```bash
$ git config --global user.name "xiaoyang-xy"
$ git config --global user.email test@qq.com
```

## 配置文本编辑器

配置git会调用的文本编辑器，不设置会使用操作系统默认文本编辑器

```bash
$ git config --global core.editor "'d:/program/Notepad++/notepad++.exe' -multiInst -nottabbar -nosesson -noPlugin"
$ git config --list 
```

## 初始化git仓库 git init

创建本地仓库，生成 .git文件夹

```bash
$ cd /d/coding/git-test
$ git init
```

## 追踪文件 git add

```bash
$ touch test.c
$ git add *.c
$ git add LICENSE
```

## 初次提交 git commit 

```bash
$ git commit -m 'test1'
```

## 克隆现有仓库

```bash
$ git clone url     //默认名字
$ git clone url dict-name 	//目录 new name
```

## 版本回退

```bash
$ git reset --hard HEAD^      //代表返回修改前的状态。 HEAD^^代表修改两次前的状态 HEAD-100
$ git reset --hard 版本号
```

## 查看版本号

```bash
$ git log 
$ git log --pretty=oneline
$ git reflog
```

## 查看状态

```bash
$ git status
```

## 查看修改了那些

```bash
$ git diff 
```

## 撤销所有修改的东西

```bash
$ git checkout -- readme.txt
```

## 把暂存区的修改撤销掉 回到工作区

```bash
$ git reset HEAD readme.txt
```

## 删除文件

```bash
$ git rm readme.txt
$ git commit -m "remove file"
```

## 恢复删除的文件

```bash
$ git checkout -- readme.txt
```

## 连接远程库

```bash
$ git remote add origin git@github.com:xiaoyang-xy/learngit.git
$ git push -u origin master
$ git push origin master
要关联一个远程库，使用命令git remote add origin git@server-name:path/repo-name.git；
关联后，使用命令git push -u origin master第一次推送master分支的所有内容；
此后，每次本地提交后，只要有必要，就可以使用命令git push origin master推送最新修改；
```

## 创建分支

```bash
$ git checkout -b dev
```

`git checkout`命令加上`-b`参数表示创建并切换，相当于以下两条命令：

```bash
$ git branch dev
$ git checkout dev
```

## 将dev分支的工作成果合并到master分支上

```bash
$ git merge dev				//合并分支
$ git branch -d dev   //删除dev分支
```

创建并切换到新的`dev`分支，可以使用：

```
$ git switch -c dev
```

直接切换到已有的`master`分支，可以使用：

```
$ git switch master
```

使用新的`git switch`命令，比`git checkout`要更容易理解。

### 小结

查看分支：`git branch`

创建分支：`git branch <name>`

切换分支：`git checkout <name>`或者`git switch <name>`

创建+切换分支：`git checkout -b <name>`或者`git switch -c <name>`

合并某分支到当前分支：`git merge <name>`

删除分支：`git branch -d <name>`

## 解决冲突



