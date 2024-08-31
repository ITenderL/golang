package main

import (
	"net"
	"strings"
)

type User struct {
	// 用户名
	Name string
	// 客户端地址
	Addr string
	// 每个用户有一个channel, 跟用户绑定
	C chan string
	// 当前用户，用于跟客户端通信的
	conn net.Conn

	server *Server
}

// NewUser 创建用户
func NewUser(conn net.Conn, server *Server) *User {
	addr := conn.RemoteAddr().String()
	user := &User{
		Name:   addr,
		Addr:   addr,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}
	// 启动一个goroutine，监听当前用户channel消息
	go user.ListenMessage()
	return user
}

// ListenMessage 监听当前user的channel，一旦有消息，就发送给客户端，需要跟server一起启动
func (user *User) ListenMessage() {
	for {
		// 阻塞等待,监听C channel中的消息，如有如消息，就发送给客户端
		msg := <-user.C
		user.conn.Write([]byte(msg + "\n"))
	}
}

// 上线
func (user *User) Online() {
	// 用户上线了，用户加入到OlineMap中
	user.server.mapLock.Lock()
	user.server.OnlineMap[user.Name] = user
	user.server.mapLock.Unlock()
	// 广播当前用户上线的消息
	user.server.Broadcast(user, "已上线")
}

// 下线
func (user *User) OffLine() {
	// 用户下线了，用户OlineMap中删除
	user.server.mapLock.Lock()
	delete(user.server.OnlineMap, user.Name)
	user.server.mapLock.Unlock()
	// 广播当前用户下线的消息
	user.server.Broadcast(user, "断开连接")
}

func (user *User) SendMsg(msg string) {
	user.conn.Write([]byte(msg))
}

// 处理消息
func (user *User) DoMessage(msg string) {
	if msg == "who" {
		user.server.mapLock.Lock()
		for _, u := range user.server.OnlineMap {
			onlineMsg := "[" + u.Addr + "]" + u.Name + ": 在线\n"
			user.SendMsg(onlineMsg)
		}
		user.server.mapLock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		// 修改用户名
		newName := strings.Split(msg, "|")[1]
		// 判断name是否存在
		_, ok := user.server.OnlineMap[newName]
		if ok {
			user.SendMsg("当前用户名已经存在")
			return
		}
		user.server.mapLock.Lock()
		delete(user.server.OnlineMap, user.Name)
		user.server.OnlineMap[newName] = user
		user.server.mapLock.Unlock()
		user.Name = newName
		user.SendMsg("您已经更新用户名：" + newName + "\n")
	} else if len(msg) > 4 && msg[:3] == "to|" {
		// 消息格式 to|userName|msg
		// 1. 获取消息接收人的名称
		receiverName := strings.Split(msg, "|")[1]
		if receiverName == "" {
			user.SendMsg("消息格式不正确，请使用\"to|张三|你好\"的格式\n")
			return
		}
		// 获取接收人
		receiver, ok := user.server.OnlineMap[receiverName]
		if !ok {
			user.SendMsg("该用户不存在")
			return
		}
		message := strings.Split(msg, "|")[2]
		if message == "" {
			user.SendMsg("消息内容为空，请重发消息\n")
			return
		}
		receiver.SendMsg(user.Name + "对您说：" + message)
	} else {
		user.server.Broadcast(user, msg)
	}
}
