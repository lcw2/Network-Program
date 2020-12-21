package main

import (
	"Network-Program/proto"
	"fmt"
	"net"
)

func main() {
	conn ,err := net.Dial("tcp","127.0.0.1:20000")
	if err != nil{
		fmt.Println("fail to dial up tcp connection.")
	}

	for i:=0;i<20;i++{
		strByte,err := proto.Encode("hello world")
		if err != nil{
			fmt.Println("fail to encode data")
			return
		}
		conn.Write(strByte)
	}
}