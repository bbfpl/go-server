package ws

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
	"log"
	"time"
	"web/library"
)

var m = melody.New()

func SocketChat() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("访问ws")
		m.HandleRequest(c.Writer, c.Request)
	}
}

//所有用户
var users = make(map[string]string)

// 聊天室事件定义
type Event struct {
	Type      string      `json:"type"`      // 事件类型
	User      string      `json:"user"`      // 用户名
	Timestamp int64       `json:"timestamp"` // 时间戳
	Data      interface{} `json:"data"`      // 事件内容
}

type Users struct {
	Type string      `json:"type"` // 事件类型
	Data interface{} `json:"data"` // 事件内容
}

//获取auth
func getQueryAuth(s *melody.Session) (string, string, error) {
	token := s.Request.URL.Query().Get("token")
	//获取token解析成值
	parseToken, err := library.ParseToken(token)
	if err != nil {
		return "", "", err
	}
	//fmt.Println("UID:", parseToken.UID)
	//fmt.Println("Username:", parseToken.Username)
	return parseToken.UID, parseToken.Username, nil
}

func userBroadcast(swtype string, name interface{}) {
	e := Users{
		Type: swtype,
		Data: name,
	}
	r, err := json.Marshal(e)
	if err != nil {
		log.Printf("发生错误: %v", err)
	}
	//广播向所有会话广播文本消息
	m.Broadcast(r)
}

func ChatInit() {

	// 监听连接事件
	m.HandleConnect(func(s *melody.Session) {
		// 1. 实例化连接消息
		fmt.Println("实例化连接消息")
		uid, name, err := getQueryAuth(s)
		if err != nil {
			fmt.Println(err)
		}
		//连接成功 加入当前用户
		users[uid] = name
		userBroadcast("userJoin", users)

	})
	// 监听接收事件
	m.HandleMessage(func(s *melody.Session, msg []byte) {
		fmt.Println("接收消息:", string(msg))

		//BroadcastOthers将文本消息广播到会话s以外的所有会话。
		//m.BroadcastOthers(message)

		var data map[string]interface{}
		err := json.Unmarshal(msg, &data)
		if err != nil {
			fmt.Println("发生错误", err)
		}

		mtype, ok := data["type"]
		if !ok {
			log.Printf("type 不能为空")
		}

		//auth
		uid, name, err := getQueryAuth(s)
		if err != nil {
			fmt.Println(err)
		}

		switch mtype {
		//消息
		case "msg":

			data, ok := data["data"]
			if !ok {
				data = ""
			}

			e := Event{
				Type:      "msg",
				User:      name,
				Timestamp: time.Now().UnixNano() / 1e6,
				Data:      data,
			}

			r, err := json.Marshal(e)
			if err != nil {
				log.Printf("发生错误: %v", err)
			}
			//广播向所有会话广播文本消息
			m.Broadcast(r)
		//登出
		case "logout":
			//连接断开删除用户
			delete(users, uid)
			userBroadcast("userOut", users)
		//心跳检测
		case "ping":
			fmt.Println("ping")
		default:
			log.Fatalf("type 错误")
		}

		fmt.Println(data)

	})
	// 监听连接断开事件
	m.HandleDisconnect(func(s *melody.Session) {
		// 断开连接消息
		fmt.Println("断开连接消息")
		uid, _, err := getQueryAuth(s)
		if err != nil {
			fmt.Println(err)
		}
		//连接断开删除用户
		delete(users, uid)

		fmt.Println("删除后的用户:", users)

		userBroadcast("userOut", users)
	})

	// 监听连接错误
	m.HandleError(func(s *melody.Session, e error) {
		log.Println("发生错误", e.Error())
	})
}
