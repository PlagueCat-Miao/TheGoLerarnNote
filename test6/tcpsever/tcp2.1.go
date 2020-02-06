package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

/**
  需求:
      socket编程实现 客户端 与 服务端进行通讯
      通讯测试场景：
          1.client 发送 ping, server 返回 pong
          2.client 发送 hello, server 返回 world
          3.其余client发送内容, server 回显即可

  抽象&解决方案:
      1. socket 编程是对 tcp通讯 过程的封装，unix server端网络编程过程为 Server->Bind->Listen－>Accept
         go 中直接使用 Listen + Accept
      2. client 与客户端建立好的请求 可以被新建的 goroutine(go func) 处理 named connHandler
      3. goroutine 的处理过程其实是 输入流/输出流 的应用场景

  积累:
      1.基础语法
      2.基本数据结构 slice 使用
      3.goroutine 使用
      4.switch 使用
      4.socket 编程核心流程
      5.net 网络包使用
*/

func connHandler(c net.Conn) {
	//1.conn是否有效
	if c == nil {
		log.Panic("无效的 socket 连接")
	}

	//2.新建网络数据流存储结构
	buf := make([]byte, 4096)
	//3.循环读取网络数据流
	for {
		//3.1 网络数据流读入 buffer
		cnt, err := c.Read(buf)
		//3.2 数据读尽、读取错误 对方关闭连接  此时 关闭 socket 连接
		if cnt == 0 || err != nil {
			c.Close()
			fmt.Printf("来自 %v 的连接关闭\n", c.RemoteAddr())
			break
		}

		//3.3 根据输入流进行逻辑处理
		//buf数据 -> 去两端空格的string
		inStr := strings.TrimSpace(string(buf[0:cnt]))
		//去除 string 内部空格
		cInputs := strings.Split(inStr, " ")
		//获取 客户端输入第一条命令
		fCommand := cInputs[0]
		fmt.Printf("来自 %v 的", c.RemoteAddr())
		fmt.Println("客户端传输->" + fCommand)

		switch fCommand {
		case "ping":
			c.Write([]byte("服务器端回复-> pong\n"))
		case "hello":
			c.Write([]byte("服务器端回复-> world\n"))
		default:
			c.Write([]byte("服务器端回复" + fCommand + "\n"))
		}

		//c.Close() //关闭client端的连接，telnet 被强制关闭

		//此时还在连接，只是进行了一次recv
	}
}


//开启serverSocket
func ServerSocket() {
	//1.监听端口
	server, err := net.Listen("tcp", ":8087")

	if err != nil {
		fmt.Println("开启socket服务失败")
	}

	fmt.Println("Server 开始监听...")

	for {
		//2.接收来自 client 的连接,会阻塞
		conn, err := server.Accept()

		if err != nil {
			fmt.Println("连接出错")
		}

		//并发模式 接收来自客户端的连接请求，一个连接 建立一个 conn，服务器资源有可能耗尽 BIO模式
		fmt.Println("new socket ",conn.RemoteAddr(), "Departing now")
		go connHandler(conn)
	}
	//fmt.Println("Server 意外结束")
}
//client
func main() {

	ServerSocket();

}
