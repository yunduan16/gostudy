// net接口
package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		buf := make([]byte, 128) //客户端和服务端接受buf不一样长，当客户端发生长数据时会截取
		n, err := reader.Read(buf)
		//当err=EOF时不要报错，特殊处理一下
		if err == io.EOF {
			fmt.Println("client close or quit")
			break
		} else if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到客户端发来的数据", recvStr)
		conn.Write([]byte(recvStr))
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn)
	}

}
