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