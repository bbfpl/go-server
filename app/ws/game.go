package ws

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
	"log"
)

var mm = melody.New()

func SocketGame() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("访问ws")
		mm.HandleRequest(c.Writer, c.Request)
	}
}

func gameBroadcast(swtype string, name interface{}) {
	e := Users{
		Type: swtype,
		Data: name,
	}
	r, err := json.Marshal(e)
	if err != nil {
		log.Printf("发生错误: %v", err)
	}
	//广播向所有会话广播文本消息
	mm.Broadcast(r)
}

// 游戏消息定义
type GameEvent struct {
	Type      string `json:"type"`      // 事件类型
	User      string `json:"user"`      // 用户名
	Uid       string `json:"uid"`       // 用户id
	PosX      string `json:"x"`         // x坐标
	PosY      string `json:"y"`         // x坐标
	Direction string `json:"direction"` // x坐标
}

type BulletEvent struct {
	Type    string      `json:"type"`    // 事件类型
	Bullets interface{} `json:"bullets"` // 数组
}

//在线用户
var OnlineUsers = make(map[string]GameEvent)

func GameInit() {

	// 监听连接事件
	mm.HandleConnect(func(s *melody.Session) {
		// 1. 实例化连接消息
		fmt.Println("ws game 连接消息")
		uid, name, err := getQueryAuth(s)
		if err != nil {
			fmt.Println(err)
		}
		//OnlineUsers["1234567"] = GameEvent{
		//	Type:      "pos",
		//	User:      "admin",
		//	Uid:       "1234567",
		//	PosX:      "250",
		//	PosY:      "250",
		//	Direction: "0",
		//}
		OnlineUsers[uid] = GameEvent{
			Type:      "pos",
			User:      name,
			Uid:       uid,
			PosX:      "150",
			PosY:      "150",
			Direction: "0",
		}
		//玩家加入
		gameBroadcast("playJoin", OnlineUsers)
	})

	// 监听接收事件
	mm.HandleMessage(func(s *melody.Session, msg []byte) {
		//fmt.Println("ws game 接收消息:", string(msg))
		var data map[string]interface{}
		err := json.Unmarshal(msg, &data)
		if err != nil {
			fmt.Println("发生错误", err)
		}
		switch data["type"] {
		//坦克位置消息
		case "pos":
			uid, name, err := getQueryAuth(s)
			if err != nil {
				fmt.Println(err)
			}
			e := GameEvent{
				Type:      "pos",
				User:      name,
				Uid:       uid,
				PosX:      data["x"].(string),
				PosY:      data["y"].(string),
				Direction: data["direction"].(string),
			}
			//更新数据
			OnlineUsers[uid] = e

			r, err := json.Marshal(e)
			if err != nil {
				log.Printf("发生错误: %v", err)
			}
			//广播向所有会话广播文本消息
			mm.Broadcast(r)
		//子弹位置消息
		case "bullets":
			_, _, err := getQueryAuth(s)
			if err != nil {
				fmt.Println(err)
			}
			e := BulletEvent{
				Type:    "bullets",
				Bullets: data["bullets"],
			}

			r, err := json.Marshal(e)
			if err != nil {
				log.Printf("发生错误: %v", err)
			}
			//广播向所有会话广播文本消息
			mm.Broadcast(r)
		default:
			log.Fatalf("type 错误")
		}

	})

	// 监听连接断开事件
	mm.HandleDisconnect(func(s *melody.Session) {
		// 断开连接消息
		fmt.Println("ws game 断开连接消息")
		uid, _, err := getQueryAuth(s)
		if err != nil {
			fmt.Println(err)
		}
		delete(OnlineUsers, uid)
		gameBroadcast("playOut", uid)
	})

	// 监听连接错误
	mm.HandleError(func(s *melody.Session, e error) {
		log.Println("ws game 发生错误", e.Error())
	})
}
