package main

import (
	"demorpc"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main()  {
	err := rpc.Register(demorpc.DemoService{})
	if err != nil {
		panic(err)
	}
	listener, err :=net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error:%s", err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}
