# GoBoot
一个构建Api服务的简易高效的脚手架

* 可以认为是一个go的脚手架
* 使用简单的mvc，入门简单，上手容易
* 适用于提供Api服务使用
* 经过时间考验，稳定高效

# 功能说明
* 路由规则：通过c和a进行路由控制，例如http://127.0.0.1:9360/?c=Demo&a=GetMsg&id=123 会执行DemoController的GetMsgAction方法
* 配置读取：
    * 支持key = value配置
    * 支持配置文件include
* 错误处理：
    * panic+recover来实现类似throw和catch的异常控制流
    * errorno, errormsg来区分不同错误
    * HttpApiServer会捕获最终异常并输出错误信息
* 日志处理：
    * 支持debug, access, warn, error 4个level的log
    * 支持直接打印map等复杂数据类型
    * log文件自动按天分割;
* 数据访问：支持访问mongodb, 依赖mgo, 请参考http://labix.org/mgo
* 系统监控
    * lvs存活监控：/status.go
    * 应用监控：/?c=Monitor&a=GetHealth&format=json format=php可以输入php序列化的数据格式，以兼容某些监控系统 
    * pprof：/debug/pprof/ (返回html, 在浏览器中直接访问)

访问mongo db说明：
    如果使用mongo需要当前的go path路径中有mgo模块,  mgo请参考http://labix.org/mgo
为了防止因为依赖问题导致编译通不过，qgo中默认将mongo相关的代码和配置注释掉了， 如果确定要使用，可以打开conf/, main.go,  src/lib/Mongo.go  
并打开被注释的相关代码和配置项

# 使用方法
1. 将GoBoot放到你的$GOPATH/src目录下
2. 运行./make进行编译，编译好的程序在$GOPATH/GoBoot/bin目录
3. 运行./serverctl envinit mechine-a 初始化conf环境
4. 运行./serverctl start [-f] 启动 -f选项表示强制启动不管已经存在的pid文件, 适合机器重启的场景
5. 运行./serverctl stop 停止
6. 运行./serverctl restart 重启
7. 执行curl "http://127.0.0.1:9360/?c=Demo&a=GetMsg&id=123" 查看是否正常返回
8. 执行curl "http://127.0.0.1:9360/?c=Demo&a=GetMsg&id=0" 查看是否错误返回
9. pprof: 在浏览器中输入http://[yourhost]:9360/debug/pprof/, 可实时查看程序内部消耗

