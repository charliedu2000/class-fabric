# default

QNXG 使用的默认 Golang 开发模板

## 程序参数

```bash
$ go run ./cmd/web --mode=dev
```

1. mode 有两个选项：dev、prod 分别加载可执行文件目录下的 `application.yml` 与 `application_dev.yml`

## 示例配置

```yaml
# 运行模式
Mode: prod

HTTP:
  # 运行端口
  Port: :7500
  # 允许的最大内容长度
  MaxContentLength: 67108864
  # http优雅关闭等待超时时长(单位秒)
  ShutdownTimeout: 30
  # 允许输出的最大日志长度
  MaxLoggerLength: 4096

CORS:
  # 是否启用
  Enable: false
  # 允许跨域请求的域名列表(*表示全部允许)
  AllowOrigins: ["*"]
  # 允许跨域请求的请求方式列表
  AllowMethods: ["GET", "POST", "PUT", "DELETE", "PATCH"]
  # 允许客户端与跨域请求一起使用的非简单标头的列表
  AllowHeaders: []
  # 请求是否可以包含cookie，HTTP身份验证或客户端SSL证书等用户凭据
  AllowCredentials: true
  # 可以缓存预检请求结果的时间（以秒为单位）
  MaxAge: 7200


GORM:
  # 是否开启调试模式
  Debug: false
  # 设置连接可以重用的最长时间(单位：秒)
  MaxLifetime: 7200
  # 设置数据库的最大打开连接数
  MaxOpenConns: 150
  # 设置空闲连接池中的最大连接数
  MaxIdleConns: 50
  # 是否启用自动映射数据库表结构
  EnableAutoMigrate: true

DB:
  # 连接地址
  Host: url
  # 连接端口
  Port: 3306
  # 用户名
  User: root
  # 密码
  Password: password
  # 数据库
  DBName: management
  # 连接参数
  Parameters: charset=utf8mb4&parseTime=True&loc=Local

Log:
  # 日志级别 -1 Debug, -2 Info, -3 Warn, -4 Error, -5 DPanic, -6Panic, -7Fatal
  Level: -2
  # 日志输出 stdout file
  Output: file

LogFileHook:
  # 日志输出文件
  Filename: ./log/default_refactor.log
  # 日志文件最大大小
  MaxSize: 128
  # 日志备份
  MaxBackups: 3
  # 日志保存时间
  MaxAge: 7
  # 是否开启压缩
  Compress: true
```
