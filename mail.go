package main

import (
	"gopkg.in/gomail.v2"
)

var content = `<div><h1 id="toc_0" style="-webkit-print-color-adjust: exact; margin-right: 0px; margin-bottom: 10px; margin-left: 0px; padding: 0px; -webkit-font-smoothing: antialiased; cursor: text; position: relative; font-size: 28px; font-family: Helvetica, arial, sans-serif; widows: auto; margin-top: 0px !important;">Backbone-cmd</h1><p style="-webkit-print-color-adjust: exact; margin: 15px 0px; font-family: Helvetica, arial, sans-serif; line-height: 22px; widows: auto;">Backbone-cmd是用来开发，运维，测试部署backone的命令行工具，help如下：</p><pre style="-webkit-print-color-adjust: exact; margin-top: 15px; margin-bottom: 15px; border: 1px solid rgb(204, 204, 204); font-size: 13px; line-height: 19px; overflow: auto; padding: 6px 10px; border-radius: 3px; widows: auto; background-color: rgb(248, 248, 248);"><code style="-webkit-print-color-adjust: exact; margin: 0px; padding: 0px; white-space: pre; border: none; background-color: transparent; border-top-left-radius: 3px; border-top-right-radius: 3px; border-bottom-right-radius: 3px; border-bottom-left-radius: 3px; background-position: initial initial; background-repeat: initial initial;">NAME:
   backbone-cli - The Backbone Framework Helper

USAGE:
   backbone-cli [global options] command [command options] [arguments...]

VERSION:
   0.0.0

AUTHOR(S):
   YanHuang &lt;yanhuang@deepglint.com&gt;

COMMANDS:
   component    Components actions
   vulcand  Generate Vulcand Script
   git      For github things
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h       show help
   --version, -v    print the version</code></pre><h3 id="toc_1" style="-webkit-print-color-adjust: exact; margin: 20px 0px 10px; padding: 0px; -webkit-font-smoothing: antialiased; cursor: text; position: relative; font-size: 18px; font-family: Helvetica, arial, sans-serif; widows: auto;">component</h3><p style="-webkit-print-color-adjust: exact; margin: 15px 0px; font-family: Helvetica, arial, sans-serif; line-height: 22px; widows: auto;">建立一个基于backbone下的组件单元，其中命令如下:</p><pre style="-webkit-print-color-adjust: exact; margin-top: 15px; margin-bottom: 15px; border: 1px solid rgb(204, 204, 204); font-size: 13px; line-height: 19px; overflow: auto; padding: 6px 10px; border-radius: 3px; widows: auto; background-color: rgb(248, 248, 248);"><code style="-webkit-print-color-adjust: exact; margin: 0px; padding: 0px; white-space: pre; border: none; background-color: transparent; border-top-left-radius: 3px; border-top-right-radius: 3px; border-bottom-right-radius: 3px; border-bottom-left-radius: 3px; background-position: initial initial; background-repeat: initial initial;">backbone-cmd component create hello
</code></pre><p style="-webkit-print-color-adjust: exact; margin: 15px 0px; font-family: Helvetica, arial, sans-serif; line-height: 22px; widows: auto;">产生如下的文件夹</p><pre style="-webkit-print-color-adjust: exact; margin-top: 15px; margin-bottom: 15px; border: 1px solid rgb(204, 204, 204); font-size: 13px; line-height: 19px; overflow: auto; padding: 6px 10px; border-radius: 3px; widows: auto; background-color: rgb(248, 248, 248);"><code style="-webkit-print-color-adjust: exact; margin: 0px; padding: 0px; white-space: pre; border: none; background-color: transparent; border-top-left-radius: 3px; border-top-right-radius: 3px; border-bottom-right-radius: 3px; border-bottom-left-radius: 3px; background-position: initial initial; background-repeat: initial initial;">$ ls
hello.go        hello_test.vue
hello.html      package.json
hello.js        webpack.config.js
hello.vue
</code></pre><p style="-webkit-print-color-adjust: exact; margin: 15px 0px; font-family: Helvetica, arial, sans-serif; line-height: 22px; widows: auto;">之后运行</p><pre style="-webkit-print-color-adjust: exact; margin-top: 15px; margin-bottom: 15px; border: 1px solid rgb(204, 204, 204); font-size: 13px; line-height: 19px; overflow: auto; padding: 6px 10px; border-radius: 3px; widows: auto; background-color: rgb(248, 248, 248);"><code style="-webkit-print-color-adjust: exact; margin: 0px; padding: 0px; white-space: pre; border: none; background-color: transparent; border-top-left-radius: 3px; border-top-right-radius: 3px; border-bottom-right-radius: 3px; border-bottom-left-radius: 3px; background-position: initial initial; background-repeat: initial initial;">npm install

go run hello.go
</code></pre><p style="-webkit-print-color-adjust: exact; margin: 15px 0px; font-family: Helvetica, arial, sans-serif; line-height: 22px; widows: auto;">即可</p><h3 id="toc_2" style="-webkit-print-color-adjust: exact; margin: 20px 0px 10px; padding: 0px; -webkit-font-smoothing: antialiased; cursor: text; position: relative; font-size: 18px; font-family: Helvetica, arial, sans-serif; widows: auto;">vulcand</h3><p style="-webkit-print-color-adjust: exact; margin: 15px 0px; font-family: Helvetica, arial, sans-serif; line-height: 22px; widows: auto;">产生vulcand配置脚本在指定的位置，使用方法如下：</p><pre style="-webkit-print-color-adjust: exact; margin-top: 15px; margin-bottom: 15px; border: 1px solid rgb(204, 204, 204); font-size: 13px; line-height: 19px; overflow: auto; padding: 6px 10px; border-radius: 3px; widows: auto; background-color: rgb(248, 248, 248);"><code style="-webkit-print-color-adjust: exact; margin: 0px; padding: 0px; white-space: pre; border: none; background-color: transparent; border-top-left-radius: 3px; border-top-right-radius: 3px; border-bottom-right-radius: 3px; border-bottom-left-radius: 3px; background-position: initial initial; background-repeat: initial initial;">backbone-cmd vulcand hello world [vulcand-admin-server] [service-server]
</code></pre><p style="-webkit-print-color-adjust: exact; margin: 15px 0px; font-family: Helvetica, arial, sans-serif; line-height: 22px; widows: auto;">执行完会生成<code style="-webkit-print-color-adjust: exact; margin: 0px 2px; padding: 0px 5px; white-space: nowrap; border: 1px solid rgb(234, 234, 234); background-color: rgb(248, 248, 248); border-top-left-radius: 3px; border-top-right-radius: 3px; border-bottom-right-radius: 3px; border-bottom-left-radius: 3px;">gen_to_vulcand.sh</code>的脚本</p><p style="-webkit-print-color-adjust: exact; margin: 15px 0px; font-family: Helvetica, arial, sans-serif; line-height: 22px; widows: auto;">运行之后就会在指定的vulcand上部署服务,通过访问</p><pre style="-webkit-print-color-adjust: exact; margin-top: 15px; margin-bottom: 15px; border: 1px solid rgb(204, 204, 204); font-size: 13px; line-height: 19px; overflow: auto; padding: 6px 10px; border-radius: 3px; widows: auto; background-color: rgb(248, 248, 248);"><code style="-webkit-print-color-adjust: exact; margin: 0px; padding: 0px; white-space: pre; border: none; background-color: transparent; border-top-left-radius: 3px; border-top-right-radius: 3px; border-bottom-right-radius: 3px; border-bottom-left-radius: 3px; background-position: initial initial; background-repeat: initial initial;">[vulcand-admin-server]/hello/world/</code></pre><p style="-webkit-print-color-adjust: exact; margin: 15px 0px; font-family: Helvetica, arial, sans-serif; line-height: 22px; widows: auto;">来访问服务</p><h3 id="toc_3" style="-webkit-print-color-adjust: exact; margin: 20px 0px 10px; padding: 0px; -webkit-font-smoothing: antialiased; cursor: text; position: relative; font-size: 18px; font-family: Helvetica, arial, sans-serif; widows: auto;">git</h3><p style="-webkit-print-color-adjust: exact; margin: 15px 0px; font-family: Helvetica, arial, sans-serif; line-height: 22px; widows: auto;">关于自动化git的操作</p><pre style="-webkit-print-color-adjust: exact; margin-top: 15px; margin-bottom: 15px; border: 1px solid rgb(204, 204, 204); font-size: 13px; line-height: 19px; overflow: auto; padding: 6px 10px; border-radius: 3px; widows: auto; background-color: rgb(248, 248, 248);"><code style="-webkit-print-color-adjust: exact; margin: 0px; padding: 0px; white-space: pre; border: none; background-color: transparent; border-top-left-radius: 3px; border-top-right-radius: 3px; border-bottom-right-radius: 3px; border-bottom-left-radius: 3px; background-position: initial initial; background-repeat: initial initial;">NAME:
   backbone-cli git - For github things

USAGE:
   backbone-cli git command [command options] [arguments...]

COMMANDS:
   sync     make a client to update
   watch    always update the github repo
   help, h  Shows a list of commands or help for one command

OPTIONS:
   --help, -h   show help
</code></pre><p style="-webkit-print-color-adjust: exact; margin: 15px 0px; font-family: Helvetica, arial, sans-serif; line-height: 22px; widows: auto;">watch功能指的是监控当前的文件夹，每10s进行<code style="-webkit-print-color-adjust: exact; margin: 0px 2px; padding: 0px 5px; white-space: nowrap; border: 1px solid rgb(234, 234, 234); background-color: rgb(248, 248, 248); border-top-left-radius: 3px; border-top-right-radius: 3px; border-bottom-right-radius: 3px; border-bottom-left-radius: 3px;">git pull</code>命令，同时提供restful服务对服务进行更新</p><pre style="-webkit-print-color-adjust: exact; margin-top: 15px; margin-bottom: 15px; border: 1px solid rgb(204, 204, 204); font-size: 13px; line-height: 19px; overflow: auto; padding: 6px 10px; border-radius: 3px; widows: auto; background-color: rgb(248, 248, 248);"><code style="-webkit-print-color-adjust: exact; margin: 0px; padding: 0px; white-space: pre; border: none; background-color: transparent; border-top-left-radius: 3px; border-top-right-radius: 3px; border-bottom-right-radius: 3px; border-bottom-left-radius: 3px; background-position: initial initial; background-repeat: initial initial;">backbone-cmd git watch -p 8031
</code></pre><p style="-webkit-print-color-adjust: exact; margin: 15px 0px; font-family: Helvetica, arial, sans-serif; line-height: 22px; widows: auto;">sync指的是同步指定的远程repo</p><pre style="-webkit-print-color-adjust: exact; margin-top: 15px; border: 1px solid rgb(204, 204, 204); font-size: 13px; line-height: 19px; overflow: auto; padding: 6px 10px; border-radius: 3px; widows: auto; margin-bottom: 0px !important; background-color: rgb(248, 248, 248);"><code style="-webkit-print-color-adjust: exact; margin: 0px; padding: 0px; white-space: pre; border: none; background-color: transparent; border-top-left-radius: 3px; border-top-right-radius: 3px; border-bottom-right-radius: 3px; border-bottom-left-radius: 3px; background-position: initial initial; background-repeat: initial initial;">backbone-md git sync http://192.168.5.46:8031</code></pre></div><div><includetail><!--<![endif]--></includetail></div>`

func main2() {
	m := gomail.NewMessage()
	m.SetHeader("From", "yanhuang@deepglint.com")
	m.SetHeader("To", "yanhuang@deepglint.com")
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", content)
	m.Attach("attach.png")

	d := gomail.NewPlainDialer("smtp.exmail.qq.com", 465, "aaa", "bbb")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
