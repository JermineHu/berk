# 个人博客的api服务

### 主要用于以下模块的管理

    1、用户管理；
    2、文章管理； 
    3、文章类型管理； 
    4、菜单管理；
    5、菜单类型管理
    6、标签管理
    7、图片上传,文件上传，文件管理
    8、每個模塊都有回收站的功能
    9、用户头像设置，1对多的关系
    10、用户联系方式设置，1对多的关系，类似一对多的关系都可以放到mongodb
    11、系统设置，一个json格式的配置，对于后台来说就是一个map结构
    12、文件管理，可以一键清除没用的图片，和资源
    13、
    
    
    
    
### 注意事项：
    
    1、使用快捷方式进行代码的映射：
    
     ln -s /home/coding/workspace/* /home/coding/go/src/git.vdo.space/berk
     
    2、使用go的最小化构建生成代码：
         
        1、使用 CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o berk 进行 go的编译；

        2、使用空镜像 scratch 或者带有 sh 等基础命令的 busybox 、alpine；

        3、打印日志要放到 go里面的话，只需要使用 scratch 就可以了，obyte的大小，也就说完全将自己的程序跑在镜像里面不需要额外支持；

        4、windows下需要 如下设置：

        set GOPATH=E:\GO_WORKSPACE
        set CGO_ENABLED=0
        set GOARCH=amd64
        set GOOS=linux
        go build -a -installsuffix cgo -o berk



    
