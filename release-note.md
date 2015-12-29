#Backbone-cmd

Backbone-cmd是用来开发，运维，测试部署backone的命令行工具，help如下：

```
NAME:
   backbone-cli - The Backbone Framework Helper

USAGE:
   backbone-cli [global options] command [command options] [arguments...]

VERSION:
   0.0.0

AUTHOR(S):
   YanHuang <yanhuang@deepglint.com>

COMMANDS:
   component	Components actions
   vulcand	Generate Vulcand Script
   git		For github things
   help, h	Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h		show help
   --version, -v	print the version
```

###component

建立一个基于backbone下的组件单元，其中命令如下:

```
backbone-cmd component create hello

```

产生如下的文件夹

```
$ ls
hello.go		hello_test.vue
hello.html		package.json
hello.js		webpack.config.js
hello.vue

```

之后运行

```
npm install

go run hello.go

```
即可

###vulcand

产生vulcand配置脚本在指定的位置，使用方法如下：

```
backbone-cmd vulcand hello world [vulcand-admin-server] [service-server]

```
执行完会生成```gen_to_vulcand.sh```的脚本

运行之后就会在指定的vulcand上部署服务,通过访问

```
[vulcand-admin-server]/hello/world/
```
来访问服务

###git

关于自动化git的操作

```
NAME:
   backbone-cli git - For github things

USAGE:
   backbone-cli git command [command options] [arguments...]

COMMANDS:
   sync		make a client to update
   watch	always update the github repo
   help, h	Shows a list of commands or help for one command

OPTIONS:
   --help, -h	show help

```

watch功能指的是监控当前的文件夹，每10s进行```git pull```命令，同时提供restful服务对服务进行更新

```
backbone-cmd git watch -p 8031

```

sync指的是同步指定的远程repo

```
backbone-md git sync http://192.168.5.46:8031

```

