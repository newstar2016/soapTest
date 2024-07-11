## 介绍
基于github.com/hooklift/gowsdl包，自动生成soap客户端和服务端

## 步骤

一、安装命令行工具

go install github.com/hooklift/gowsdl/cmd/gowsdl@latest

二、获取wsdl文件
例如：https://xxxxxxx/zj/RichHealthThridInterfaceForCard.svc\?wsdl


先在文件夹下建立一个目标文件, rich.wsdl

命令行执行：
curl https://xxxxxx/zj/RichHealthThridInterfaceForCard.svc\?wsdl > rich.wsdl

<span style="color: red; "> 域名自行替换</span> 

三、基于wsdl文件生成go的package

gowsdl -p rich -i rich.wsdl


## Usage

Usage: gowsdl [options] myservice.wsdl
       

​        -o string  File where the generated code will be saved (default "myservice.go")
​      

​        -p string  Package under which code will be generated (default "myservice")

​        -i    Skips TLS Verification
​       

​       -v    Shows gowsdl version
