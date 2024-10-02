package main

import (
	"context"
	"dailyPractice/grpc/server/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := proto.NewTestClient(conn)
	res, err := client.Ask(context.Background(), &proto.AskReq{Question: "hihihi"})
	if err != nil {
		panic(err)
	}
	log.Printf("res = %s", res.Answer)
}
