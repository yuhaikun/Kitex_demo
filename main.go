package main

import (
	kitex_gen "NewTikTok/kitex_gen/adder"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", ":8081")
	var opts []server.Option
	opts = append(opts, server.WithServiceAddr(addr))

	svr := kitex_gen.NewServer(new(AdderImpl), opts...)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
