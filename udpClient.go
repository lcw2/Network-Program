package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("连接服务端失败，err:", err)
		return
	}
	defer socket.Close()

	for {
		reader := bufio.NewReader(os.Stdin)
		inputString,err := reader.ReadString('\n')
		inputString = strings.Trim(inputString,"\r\n")
		if err != nil {
			return
		}
		if strings.ToUpper(inputString) == "Q"{
			fmt.Println("quit")
			return
		}
		inputByte := bytes.NewBuffer([]byte(inputString))
		_,err = socket.Write(inputByte.Bytes())
		if err != nil{
			fmt.Println("fail to send data.")
			return
		}
		data := make([]byte,1024)
		_,err = socket.Read(data)
		if err != nil{
			fmt.Println("fail to receive data.")
		}
		fmt.Println("the server send data:",string(data))

	}
}