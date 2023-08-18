package main

import (
	"NewTikTok/kitex_gen"
	"NewTikTok/kitex_gen/adder"
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"log"
	"time"
)

func main() {
	c, err := adder.NewClient("Adder", client.WithHostPorts("0.0.0.0:8081"))
	if err != nil {
		log.Fatal(err)
	}

	req := &kitex_gen.AddRequest{
		X: 3,
		Y: 4,
	}
	resp, err := c.Add(context.Background(), req, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}
