package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip   string
	Port int
	// 在线用户列表
	OnlineMap map[string]*User
	mapLock   sync.RWMutex
	// 广播消息的channel
	Message chan string
}

// NewServer 创建服务,SERVER,
func NewServer(ip string, port int) *Server {
	return &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
}

// 监听Message广播消息channel的goroutine，一旦有消息就发送给全部的在线的user
func (this *Server) ListenMessager() {
	for {
		msg := <-this.Message
		// 将message发送给所有的user
		this.mapLock.Lock()
		for _, cli := range this.OnlineMap {
			// 发送消息给用户的C，用于发送给客户端C
			cli.C <- msg
		}
		this.mapLock.Unlock()
	}
}

// 广播消息
func (this *Server) Broadcast(user *User, msg string) {
	// 遍历OnlineMap，发送消息
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	// 用户上线消息发送到Channel
	this.Message <- sendMsg
}

func (this *Server) Handler(conn net.Conn) {
	// 当前连接的业务
	fmt.Println("连接创建成功........")
	user := NewUser(conn, this)
	// 用户上线了，用户加入到OlineMap中
	user.Online()
	isAlive := make(chan bool)
	// 接收客户端发送的消息
	go func() {
		for {
			// 读取数据
			buf := make([]byte, 4096)
			n, err := conn.Read(buf)
			if n == 0 {
				user.OffLine()
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println("conn read error: ", err)
				return
			}
			// 提取用户的消息
			msg := string(buf[:n-1])
			// 广播消息
			user.DoMessage(msg)
		}
		// 用户任意消息，代表用户活跃
		isAlive <- true
	}()
	for {
		// 当前handler阻塞
		select {
		case <-isAlive:
		// 当前用户是活跃的，重置定时器，不做任何事情，为了激活select，重置定时器
		case <-time.After(300 * time.Second):
			user.SendMsg("您已断开连接")
			// 关闭资源
			close(user.C)
			// 管理链接
			conn.Close()
			// 退出当前handler,runtime.Goexit()
			return
		}
	}
}

// Start 启动服务
func (this *Server) Start() {
	// 创建socket  listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("listen error: ", err)
		return
	}
	// 关闭listen socket listen
	defer listener.Close()

	// 启动一个goroutine，监听Message channel
	go this.ListenMessager()
	// 循环监听accept
	for {
		// accept，阻塞等待，获取连接，之后用户上线了
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept error: ", err)
			continue
		}
		// 创建goroutine处理,do handler
		go this.Handler(conn)
	}
}
