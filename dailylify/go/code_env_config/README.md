# intellj IDEA 配置go 开发环境

1. 下载安装版本go 1.8
2. 按照提示下一步下一步安装好 go 后，应该是在C:\Go 目录
3. 可以在cmd 命令行$go 试试看能不能答应出usag
4. mac 可以使用homebrew 安装，安装完后应该是在homebrew 下的Celler 下面
5. 配置gopath 目录指向一个你用于存放go 源码、包、可执行文件的目录，
   并且go 强制要求在此目录下简历src bin pkg目录存储相关的go 源码文件和包文件
   当你使用go get github.com/urlshor.....之类的包的时候就会存储在这些src 目录下了
6. 下载安装intellj IDEA , 地址：http://www.jetbrains.com/idea/
	激活可以使用licenseServer 激活，server可以使用：http://idea.iteblog.com/key.php

7. 启动idea ,并在config中找到plugins ->browse repositories中输入go 安装go language
8. 新建项目 配置go sdk 选择你安装go的目录，也就是那个C:\Go 目录
9. 编写go 代码
#10. 关于包的管理,IDEA 强制包名同目录名，如果你想建立一个package在本地，然后main 包引用
	那么你需要建立相应的目录在src 目录中，然后通过import "./packagename" 引入
