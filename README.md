# go-server
go+gin+mysql+redis+websocket

给[chat-client](https://github.com/bbfpl/chat-client "chat-client")提供api和ws服务

目录结构
```go
.
├─app
│  ├─controller
│  │      user.go
│  │
│  ├─models
│  │      user.go
│  │
│  ├─router
│  │      router.go
│  │
│  └─ws
│          chat.go
│          game.go
│
├─conf
│      config.yaml
│
├─library
│  │  config.go
│  │  db.go
│  │  jwt.go
│  │  util.go
│  │
│  └─redis
│          pool.go
│          util.go
│
└─public
    ├─static
    └─web
```

第三方模块列表
```go
json web token 模块
	github.com/dgrijalva/jwt-go
redis
	github.com/garyburd/redigo
web 框架
	github.com/gin-gonic/gin
orm 框架
	github.com/jinzhu/gorm
	github.com/go-sql-driver/mysql
监听config.yaml是否改变,用于热更新
	github.com/fsnotify/fsnotify
	github.com/spf13/viper
melody封装了websocket
	gopkg.in/olahol/melody.v1
```

预览
go run main.go 运行服务
![1](http://demo.uihtml.com/gitimg/goserver/1.png "1")

聊天时的数据
![2](http://demo.uihtml.com/gitimg/goserver/2.png "2")

游戏时的数据
![3](http://demo.uihtml.com/gitimg/goserver/3.png "3")





