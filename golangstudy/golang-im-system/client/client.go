package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
	flag       int
}

func NewClient(serverIp string, serverPort int) *Client {
	// 创建客户端对象
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		flag:       999,
	}
	// 链接server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("net dial error:", err)
		return nil
	}
	client.conn = conn
	// 返回客户端对象
	return client
}

func (client *Client) menu() bool {
	var f int
	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.更新用户名")
	fmt.Println("0.退出")
	fmt.Scanln(&f)
	if f >= 0 && f <= 3 {
		client.flag = f
		return true
	} else {
		fmt.Println(">>>>请输入合法范围内的数字<<<<")
		return false
	}
}

func (client *Client) SelectUsers() {
	sendMsg := "who\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("client.selectUsers conn write error: ", err)
		return
	}
}

// 私聊模式，需要先查询在线用户
func (client *Client) PrivateChat() {
	var remoteName string
	var chatMsg string
	client.SelectUsers()
	fmt.Println(">>>>>请输入聊天对象【用户名】,exit退出")
	fmt.Scanln(&remoteName)
	for remoteName != "exit" && len(remoteName) != 0 {
		fmt.Println(">>>>>>>>>>>>请输入聊天内容，exit")
		fmt.Scanln(&chatMsg)
		for len(chatMsg) != 0 && chatMsg != "exit" {
			sendMsg := "to|" + remoteName + "|" + chatMsg + "\n"
			_, err := client.conn.Write([]byte(sendMsg))
			if err != nil {
				fmt.Println("client.PrivateChat conn write error: ", err)
				return
			}
			chatMsg = ""
			fmt.Println(">>>>>>>>>>>>请输入聊天内容，exit退出")
			fmt.Scanln(&chatMsg)
		}
		client.SelectUsers()
		fmt.Println(">>>>>请输入聊天对象【用户名】,exit退出")
		fmt.Scanln(&remoteName)
	}
}

// 公聊模式
func (client *Client) PublicChat() {
	// 提示用户输入消息
	var chatMsg string
	fmt.Println(">>>>>>>>请输入聊天内容，exit退出<<<<<<<<")
	fmt.Scanln(&chatMsg)
	for chatMsg != "exit" {
		if len(chatMsg) != 0 {
			sendMsg := chatMsg + "\n"
			_, err := client.conn.Write([]byte(sendMsg))
			if err != nil {
				fmt.Println("client conn write error: ", err)
				break
			}
		}
		chatMsg = ""
		fmt.Println(">>>>>>>>请输入聊天内容，exit退出<<<<<<<<")
		fmt.Scanln(&chatMsg)
	}
}

func (client *Client) UpdateName() bool {
	fmt.Println(">>>>请输入用户名:")
	fmt.Scanln(&client.Name)
	sendMsg := "rename|" + client.Name + "\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn.Write err:", err)
		return false
	}
	return true
}

// 处理server回应的消息， 直接显示到标准输出即可
func (client *Client) DealResponse() {
	// 一旦client.conn有数据，就直接copy到stdout标准输出上, 永久阻塞监听
	io.Copy(os.Stdout, client.conn)
	// for {
	// 	buf := make([]byte, 1024)
	// 	client.conn.Read(buf)
	// 	fmt.Println(buf)
	// }
}

func (client *Client) Run() {
	for client.flag != 0 {
		for client.menu() != true {
		}

		// 根据不同的模式处理不同的业务
		switch client.flag {
		case 1:
			// 公聊模式
			client.PublicChat()
			break
		case 2:
			// 私聊模式
			client.PrivateChat()
			break
		case 3:
			// 更新用户名
			client.UpdateName()
			break
		}
	}
}

var serverIp string
var serverPort int

func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器默认地址(默认是127.0.0.1)")
	flag.IntVar(&serverPort, "port", 8888, "设置服务器默认端口(默认是8888)")
}

func main() {
	// 命令行解析
	flag.Parse()
	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>>>> 链接服务器失败...")
		return
	}
	// 单独开启一个goroutine去处理server的回执消息
	go client.DealResponse()
	fmt.Println(">>>>>链接服务器成功...")
	// 启动客户端的业务
	client.Run()
}
