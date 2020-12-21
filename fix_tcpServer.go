package main

import (
	"Network-Program/proto"
	"bufio"
	"fmt"
	"io"
	"net"
)

// 处理函数
func process2(conn net.Conn) {
	defer conn.Close() // 关闭连接
	reader := bufio.NewReader(conn)
	for {
		recvStr,err := proto.Decode(reader)
		if err == io.EOF{
			return
		}
		if err != nil{
			fmt.Println("fail to decode data")
			return
		}
		fmt.Println("收到client端发来的数据：", recvStr)
		conn.Write([]byte(recvStr)) // 发送数据
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	for {
		conn, err := listen.Accept() // 建立连接
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process2(conn) // 启动一个goroutine处理连接
	}
}