### go后台模板


1. 初始化操作
```bash
# 初始化项目
go mod init test
```

2. 安装相关依赖

```bash
# 安装路由和mysql操作的包
go get -u github.com/gin-gonic/gin
go get -u github.com/jinzhu/gorm

# 配置文件ini包
go get -u github.com/go-ini/ini

# 跨域的包
go get -u github.com/gin-contrib/cors
```


3. 手动向初始化数据库文件中导入mysql
```go
_ "github.com/jinzhu/gorm/dialects/mysql"
```

4. 一些默认信息说明

只有release模式才会生成log日志信息,默认日志位置runtime/logs  
除了启动成功,只有致命的程序本身出错才会有日志产生.

5. 提交说明

增加了redis和image上传支持